package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/vivienda/handlers"
	"for-neos-api/internal/vivienda/storage"

	transporteHandlers "for-neos-api/internal/transporte/handlers"
	transporteStorage "for-neos-api/internal/transporte/storage"
)

func main() {
	memoria := storage.NuevaMemoria()

	memoria.SeedViviendas()
	memoria.SeedFotos()
	memoria.SeedSectores()
	memoria.SeedAplicarViviendas()

	//Transporte
	memoriaTransporte := transporteStorage.NuevaMemoriaTransporte()

	memoriaTransporte.SeedTransportes()
	memoriaTransporte.SeedCooperativas()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(memoria)
	transporteServidor := transporteHandlers.NewServer(memoriaTransporte)

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
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
