package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"for-neos-api/internal/config"
	"for-neos-api/internal/httpserver"
	"for-neos-api/internal/middleware"
	"for-neos-api/internal/storage"

	alimentacionHandlers "for-neos-api/internal/alimentacion/handlers"
	alimentacionService "for-neos-api/internal/alimentacion/service"

	transporteHandlers "for-neos-api/internal/transporte/handlers"
	transporteService "for-neos-api/internal/transporte/service"

	usuarioHandlers "for-neos-api/internal/usuario/handlers"
	usuarioService "for-neos-api/internal/usuario/service"

	viviendaHandlers "for-neos-api/internal/vivienda/handlers"
	viviendaService "for-neos-api/internal/vivienda/service"
)

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Factory: abre DB, migra, siembra y elige backend.
	recursos, err := storage.Inicializar(cfg.RutaDB, cfg.Backend, cfg.DBDriver, cfg.DBDsn)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Backend de almacenamiento: %s", recursos.BackendUsado)

	// 2. Servicios.
	auth := usuarioService.NuevoAuthService(
		recursos.Usuarios,
		usuarioService.WithSecreto(cfg.JWTSecreto),
		usuarioService.WithDuracionToken(cfg.JWTDuracion),
	)
	authServer := usuarioHandlers.NewServer(
		usuarioHandlers.Deps{
			Auth: auth,
		},
	)

	// Vivienda
	viviendaSrv := viviendaService.NuevaViviendaService(recursos.Almacen)
	fotoSrv := viviendaService.NuevaFotoService(recursos.Almacen)
	aplicarSrv := viviendaService.NuevaAplicarViviendaService(recursos.Almacen)
	sectorSrv := viviendaService.NuevaSectorService(recursos.Almacen)
	servidorVivienda := viviendaHandlers.NewServer(
		viviendaHandlers.Deps{
			Viviendas: viviendaSrv,
			Fotos:     fotoSrv,
			Aplicar:   aplicarSrv,
			Sectores:  sectorSrv,
		},
	)
	// Transporte
	cooperativaSrv := transporteService.NuevaCooperaticaService(recursos.Almacen)
	paradaBusSrv := transporteService.NuevaParadaBusService(recursos.Almacen)
	rutaTransporteSrv := transporteService.NuevaRutaService(recursos.Almacen)
	transporteServidor := transporteHandlers.NewServer(
		transporteHandlers.Deps{
			Cooperativas: cooperativaSrv,
			Paradas:      paradaBusSrv,
			Rutas:        rutaTransporteSrv,
		},
	)
	// Alimentacion
	alimentacionSrv := alimentacionService.NuevaAlimentacionService(recursos.Almacen)
	menuDiarioSrv := alimentacionService.NuevoMenuDiarioService(recursos.Almacen)
	platoSrv := alimentacionService.NuevoPlatoService(recursos.Almacen)
	resenaSrv := alimentacionService.NuevaResenaService(recursos.Almacen)
	servidorAlimentacion := alimentacionHandlers.NewServer(
		alimentacionHandlers.Deps{
			Alimentacion: alimentacionSrv,
			MenuDiario:   menuDiarioSrv,
			Plato:        platoSrv,
			Resena:       resenaSrv,
		},
	)
	// 3. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", authServer.Registrar)
		r.Post("/auth/login", authServer.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(auth))

			// Viviendas
			r.Get("/viviendas", servidorVivienda.ListarViviendas)
			r.Post("/viviendas", servidorVivienda.CrearVivienda)
			r.Get("/viviendas/{id}", servidorVivienda.ObtenerVivienda)
			r.Put("/viviendas/{id}", servidorVivienda.ActualizarVivienda)
			r.Delete("/viviendas/{id}", servidorVivienda.BorrarVivienda)

			// Fotos
			r.Get("/fotos", servidorVivienda.ListarFotos)
			r.Post("/fotos", servidorVivienda.CrearFoto)
			r.Get("/fotos/{id}", servidorVivienda.ObtenerFoto)
			r.Put("/fotos/{id}", servidorVivienda.ActualizarFoto)
			r.Delete("/fotos/{id}", servidorVivienda.BorrarFoto)

			// Sectores
			r.Get("/sectores", servidorVivienda.ListarSectores)
			r.Post("/sectores", servidorVivienda.CrearSector)
			r.Get("/sectores/{id}", servidorVivienda.ObtenerSector)
			r.Put("/sectores/{id}", servidorVivienda.ActualizarSector)
			r.Delete("/sectores/{id}", servidorVivienda.BorrarSector)

			// AplicarViviendas
			r.Get("/aplicarviviendas", servidorVivienda.ListarAplicarViviendas)
			r.Post("/aplicarviviendas", servidorVivienda.CrearAplicarVivienda)
			r.Get("/aplicarviviendas/{id}", servidorVivienda.ObtenerAplicarVivienda)
			r.Put("/aplicarviviendas/{id}", servidorVivienda.ActualizarAplicarVivienda)
			r.Delete("/aplicarviviendas/{id}", servidorVivienda.BorrarAplicarVivienda)

			// Alimentacion
			r.Get("/alimentaciones", servidorAlimentacion.ListarAlimentaciones)
			r.Post("/alimentaciones", servidorAlimentacion.CrearAlimentacion)
			r.Get("/alimentaciones/{id}", servidorAlimentacion.BuscarAlimentacionesPorID)
			r.Put("/alimentaciones/{id}", servidorAlimentacion.ActualizarAlimentacion)
			r.Delete("/alimentaciones/{id}", servidorAlimentacion.BorrarAlimentacion)

			// MenuDiarios
			r.Get("/menudiarios", servidorAlimentacion.ListarMenuDiarios)
			r.Post("/menudiarios", servidorAlimentacion.CrearMenuDiario)
			r.Get("/menudiarios/{id}", servidorAlimentacion.BuscarMenuDiarioPorID)
			r.Put("/menudiarios/{id}", servidorAlimentacion.ActualizarMenuDiario)
			r.Delete("/menudiarios/{id}", servidorAlimentacion.BorrarMenuDiario)

			// Platos
			r.Get("/platos", servidorAlimentacion.ListarPlatos)
			r.Post("/platos", servidorAlimentacion.CrearPlato)
			r.Get("/platos/{id}", servidorAlimentacion.BuscarPlatosPorID)
			r.Put("/platos/{id}", servidorAlimentacion.ActualizarPlato)
			r.Delete("/platos/{id}", servidorAlimentacion.BorrarPlato)

			// Resenas
			r.Get("/resenas", servidorAlimentacion.ListarResenas)
			r.Post("/resenas", servidorAlimentacion.CrearResena)
			r.Get("/resenas/{id}", servidorAlimentacion.BuscarResenasPorID)
			r.Put("/resenas/{id}", servidorAlimentacion.ActualizarResena)
			r.Delete("/resenas/{id}", servidorAlimentacion.BorrarResena)

			// Transporte
			r.Get("/transporte", transporteServidor.ListarRutas)
			r.Post("/transporte", transporteServidor.AgregarRuta)
			r.Get("/transporte/{id}", transporteServidor.ObtenerRutaPorID)
			r.Put("/transporte/{id}", transporteServidor.ActualizarRuta)
			r.Delete("/transporte/{id}", transporteServidor.EliminarRuta)

			// Cooperativas
			r.Get("/cooperativa", transporteServidor.ListarCooperativas)
			r.Post("/cooperativa", transporteServidor.AgregarCooperativa)
			r.Get("/cooperativa/{id}", transporteServidor.ObtenerCooperativaPorID)
			r.Put("/cooperativa/{id}", transporteServidor.ActualizarCooperativa)
			r.Delete("/cooperativa/{id}", transporteServidor.EliminarCooperativa)

			// Paradas
			r.Get("/paradas", transporteServidor.ListarParadas)
			r.Post("/paradas", transporteServidor.AgregarParada)
			r.Get("/paradas/{id}", transporteServidor.ObtenerParadaPorID)
			r.Put("/paradas/{id}", transporteServidor.ActualizarParada)
			r.Delete("/paradas/{id}", transporteServidor.EliminarParada)
		})
	})

	// 4. Servidor HTTP con timeouts.
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// 5. Graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}
