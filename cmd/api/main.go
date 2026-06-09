package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	
	 "for-neos-api/internal/vivienda/handlers"
	"for-neos-api/internal/vivienda/storage"

	handlersAlimentacion "for-neos-api/internal/alimentacion/handlers"
	storageAlimentacion "for-neos-api/internal/alimentacion/storage"
)
func main() {
	memoria := storage.NuevaMemoria()

	memoria.SeedViviendas()
	memoria.SeedFotos()
	memoria.SeedSectores()
	memoria.SeedAplicarViviendas()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(memoria)

    //Alimentacion
	memoriaAlimentacion := storageAlimentacion.NewMemoria()

	memoriaAlimentacion.SeedAlimentaciones()
	memoriaAlimentacion.SeedMenuDiarios()
	memoriaAlimentacion.SeedPlatos()
	memoriaAlimentacion.SeedResenas()

	servidorAlimentacion := handlersAlimentacion.NewServer(memoriaAlimentacion)


	// 3. Configurar el router con versionado /api/v1/.
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		// Viviendas: CRUD completo.
		r.Get("/viviendas", servidor.ListarViviendas)
		r.Post("/viviendas", servidor.CrearVivienda)
		r.Get("/viviendas/{id}", servidor.ObtenerVivienda)
		r.Put("/viviendas/{id}", servidor.ActualizarVivienda)
		r.Delete("/viviendas/{id}", servidor.BorrarVivienda)

		// Fotos: CRUD completo.
		r.Get("/fotos", servidor.ListarFotos)
		r.Post("/fotos", servidor.CrearFoto)
		r.Get("/fotos/{id}", servidor.ObtenerFoto)
		r.Put("/fotos/{id}", servidor.ActualizarFoto)
		r.Delete("/fotos/{id}", servidor.BorrarFoto)

		//Sectores: CRUD completo.
		r.Get("/sectores", servidor.ListarSectores)
		r.Post("/sectores", servidor.CrearSector)
		r.Get("/sectores/{id}", servidor.ObtenerSector)
		r.Put("/sectores/{id}", servidor.ActualizarSector)
		r.Delete("/sectores/{id}", servidor.BorrarSector)

		// AplicarViviendas: CRUD completo.
		r.Get("/aplicarviviendas", servidor.ListarAplicarViviendas)
		r.Post("/aplicarviviendas", servidor.CrearAplicarVivienda)
		r.Get("/aplicarviviendas/{id}", servidor.ObtenerAplicarVivienda)
		r.Put("/aplicarviviendas/{id}", servidor.ActualizarAplicarVivienda)
		r.Delete("/aplicarviviendas/{id}", servidor.BorrarAplicarVivienda)

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
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
