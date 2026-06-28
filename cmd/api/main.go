package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver GORM (pure-Go)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"for-neos-api/internal/middleware"
	viviendaHandlers "for-neos-api/internal/vivienda/handlers"
	viviendaModels "for-neos-api/internal/vivienda/models"
	viviendaService "for-neos-api/internal/vivienda/service"

	transporteHandlers "for-neos-api/internal/transporte/handlers"
	transporteModels "for-neos-api/internal/transporte/models"
	transporteService "for-neos-api/internal/transporte/service"

	alimentacionHandlers "for-neos-api/internal/alimentacion/handlers"
	alimentacionModels "for-neos-api/internal/alimentacion/models"
	alimentacionService "for-neos-api/internal/alimentacion/service"

	usuarioHandlers "for-neos-api/internal/usuario/handlers"
	usuarioModels "for-neos-api/internal/usuario/models"
	usuarioService "for-neos-api/internal/usuario/service"

	"for-neos-api/internal/storage"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB, migra y siembra.
	//Esto corre siempre, sin importar qué backend sirva después. Esta es la única decisión que cambia entre GORM y sqlc.
	gdb, err := gorm.Open(sqlite.Open("for-neos-api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(&viviendaModels.Vivienda{}, &viviendaModels.Foto{}, &viviendaModels.AplicarVivienda{}, &viviendaModels.Sector{}, &alimentacionModels.Alimentacion{}, &alimentacionModels.MenuDiario{}, &alimentacionModels.Plato{}, &alimentacionModels.Resena{}, &transporteModels.Cooperativa{}, &transporteModels.RutaTransporte{}, &transporteModels.ParadaBus{}, &usuarioModels.Usuario{}); err != nil {
		log.Fatal("Falló AutoMigrate: ", err)
	}
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend que SIRVE las peticiones según la variable STORAGE.
	//    >>> Esta es la ÚNICA decisión que cambia entre GORM y sqlc. <<<
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo cafeteria.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "for-neos-api.db")
		if err != nil {
			log.Fatal("No se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	// 3. Server con inyección de dependencias. No sabe qué backend recibió.

	usuarioRepo := storage.NewUsuarioGORM(gdb)
	auth := usuarioService.NuevoAuthService(usuarioRepo)
	authServer := usuarioHandlers.NewServer(auth)

	//Vivienda

	viviendaSrv := viviendaService.NuevaViviendaService(almacen)
	fotoSrv := viviendaService.NuevaFotoService(almacen)
	aplicarSrv := viviendaService.NuevaAplicarViviendaService(almacen)
	sectorSrv := viviendaService.NuevaSectorService(almacen)
	servidorVivienda := viviendaHandlers.NewServer(viviendaSrv, fotoSrv, aplicarSrv, sectorSrv)

	//Transporte
	cooperativaSrv := transporteService.NuevaCooperaticaService(almacen)
	paradaBusSrv := transporteService.NuevaParadaBusService(almacen)
	rutaTransporteSrv := transporteService.NuevaRutaService(almacen)
	transporteServidor := transporteHandlers.NewServer(cooperativaSrv, paradaBusSrv, rutaTransporteSrv)

	//Alimentacion
	alimentacionSrv := alimentacionService.NuevaAlimentacionService(almacen)
	menuDiarioSrv := alimentacionService.NuevoMenuDiarioService(almacen)
	platoSrv := alimentacionService.NuevoPlatoService(almacen)
	resenaSrv := alimentacionService.NuevaResenaService(almacen)
	servidorAlimentacion := alimentacionHandlers.NewServer(alimentacionSrv, menuDiarioSrv, platoSrv, resenaSrv)

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)
	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/auth/registrar", authServer.Registrar)
		r.Post("/auth/login", authServer.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(auth))

			// Viviendas: CRUD completo.
			r.Get("/viviendas", servidorVivienda.ListarViviendas)
			r.Post("/viviendas", servidorVivienda.CrearVivienda)
			r.Get("/viviendas/{id}", servidorVivienda.ObtenerVivienda)
			r.Put("/viviendas/{id}", servidorVivienda.ActualizarVivienda)
			r.Delete("/viviendas/{id}", servidorVivienda.BorrarVivienda)

			// Fotos: CRUD completo.
			r.Get("/fotos", servidorVivienda.ListarFotos)
			r.Post("/fotos", servidorVivienda.CrearFoto)
			r.Get("/fotos/{id}", servidorVivienda.ObtenerFoto)
			r.Put("/fotos/{id}", servidorVivienda.ActualizarFoto)
			r.Delete("/fotos/{id}", servidorVivienda.BorrarFoto)

			//Sectores: CRUD completo.
			r.Get("/sectores", servidorVivienda.ListarSectores)
			r.Post("/sectores", servidorVivienda.CrearSector)
			r.Get("/sectores/{id}", servidorVivienda.ObtenerSector)
			r.Put("/sectores/{id}", servidorVivienda.ActualizarSector)
			r.Delete("/sectores/{id}", servidorVivienda.BorrarSector)

			// AplicarViviendas: CRUD completo.
			r.Get("/aplicarviviendas", servidorVivienda.ListarAplicarViviendas)
			r.Post("/aplicarviviendas", servidorVivienda.CrearAplicarVivienda)
			r.Get("/aplicarviviendas/{id}", servidorVivienda.ObtenerAplicarVivienda)
			r.Put("/aplicarviviendas/{id}", servidorVivienda.ActualizarAplicarVivienda)
			r.Delete("/aplicarviviendas/{id}", servidorVivienda.BorrarAplicarVivienda)

			// Alimentacion: CRUD completo.
			r.Get("/alimentaciones", servidorAlimentacion.ListarAlimentaciones)
			r.Post("/alimentaciones", servidorAlimentacion.CrearAlimentacion)
			r.Get("/alimentaciones/{id}", servidorAlimentacion.BuscarAlimentacionesPorID)
			r.Put("/alimentaciones/{id}", servidorAlimentacion.ActualizarAlimentacion)
			r.Delete("/alimentaciones/{id}", servidorAlimentacion.BorrarAlimentacion)

			// MenuDiarios: CRUD completo.
			r.Get("/menudiarios", servidorAlimentacion.ListarMenuDiarios)
			r.Post("/menudiarios", servidorAlimentacion.CrearMenuDiario)
			r.Get("/menudiarios/{id}", servidorAlimentacion.BuscarMenuDiarioPorID)
			r.Put("/menudiarios/{id}", servidorAlimentacion.ActualizarMenuDiario)
			r.Delete("/menudiarios/{id}", servidorAlimentacion.BorrarMenuDiario)

			// Platos: CRUD completo.
			r.Get("/platos", servidorAlimentacion.ListarPlatos)
			r.Post("/platos", servidorAlimentacion.CrearPlato)
			r.Get("/platos/{id}", servidorAlimentacion.BuscarPlatosPorID)
			r.Put("/platos/{id}", servidorAlimentacion.ActualizarPlato)
			r.Delete("/platos/{id}", servidorAlimentacion.BorrarPlato)

			// Resenas: CRUD completo.
			r.Get("/resenas", servidorAlimentacion.ListarResenas)
			r.Post("/resenas", servidorAlimentacion.CrearResena)
			r.Get("/resenas/{id}", servidorAlimentacion.BuscarResenasPorID)
			r.Put("/resenas/{id}", servidorAlimentacion.ActualizarResena)
			r.Delete("/resenas/{id}", servidorAlimentacion.BorrarResena)

			// Rutas de Transporte: CRUD completo.
			r.Get("/transporte", transporteServidor.ListarRutas)
			r.Post("/transporte", transporteServidor.AgregarRuta)
			r.Get("/transporte/{id}", transporteServidor.ObtenerRutaPorID)
			r.Put("/transporte/{id}", transporteServidor.ActualizarRuta)
			r.Delete("/transporte/{id}", transporteServidor.EliminarRuta)

			// Cooperativas de Transporte: CRUD completo.
			r.Get("/cooperativa", transporteServidor.ListarCooperativas)
			r.Post("/cooperativa", transporteServidor.AgregarCooperativa)
			r.Get("/cooperativa/{id}", transporteServidor.ObtenerCooperativaPorID)
			r.Put("/cooperativa/{id}", transporteServidor.ActualizarCooperativa)
			r.Delete("/cooperativa/{id}", transporteServidor.EliminarCooperativa)

			//Paradas de Bus: CRUD completo
			r.Get("/paradas", transporteServidor.ListarParadas)
			r.Post("/paradas", transporteServidor.AgregarParada)
			r.Get("/paradas/{id}", transporteServidor.ObtenerParadaPorID)
			r.Put("/paradas/{id}", transporteServidor.ActualizarParada)
			r.Delete("/paradas/{id}", transporteServidor.EliminarParada)

		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
