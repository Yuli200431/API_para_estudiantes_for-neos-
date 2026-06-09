package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	// Usamos alias para que no choquen los nombres de los paquetes
	handlersAlimentacion "for-neos-api/internal/alimentacion/handlers"
	storageAlimentacion "for-neos-api/internal/alimentacion/storage"
	//handlersVivienda "for-neos-api/internal/vivienda/handlers"
	//storageVivienda "for-neos-api/internal/vivienda/storage"
)

func main() {
	// --- CONFIGURACIÓN DE VIVIENDA ---
	//memoriaVivienda := storageVivienda.NuevaMemoria()
	//memoriaVivienda.SeedViviendas()
	//memoriaVivienda.SeedFotos()
	//memoriaVivienda.SeedSectores()
	//memoriaVivienda.SeedAplicarViviendas()

	//servidorVivienda := handlersVivienda.NewServer(memoriaVivienda)

	// --- CONFIGURACIÓN DE ALIMENTACIÓN ---
	memoriaAlimentacion := storageAlimentacion.NewMemoria()
	memoriaAlimentacion.SeedAlimentaciones()
	memoriaAlimentacion.SeedMenuDiarios()
	memoriaAlimentacion.SeedPlatos()
	memoriaAlimentacion.SeedResenas()

	servidorAlimentacion := handlersAlimentacion.NewServer(memoriaAlimentacion)

	// --- ROUTER ---
	r := chi.NewRouter()

	// Decidan con su grupo si prefieren usar "/api" o "/api/v1"
	r.Route("/api/v1", func(r chi.Router) {
		
		// Viviendas: CRUD completo.
		//r.Get("/viviendas", servidorVivienda.ListarViviendas)
		//r.Post("/viviendas", servidorVivienda.CrearVivienda)
		//r.Get("/viviendas/{id}", servidorVivienda.ObtenerVivienda)
		//r.Put("/viviendas/{id}", servidorVivienda.ActualizarVivienda)
		//r.Delete("/viviendas/{id}", servidorVivienda.BorrarVivienda)

		// Fotos: CRUD completo.
		//r.Get("/fotos", servidorVivienda.ListarFotos)
		//r.Post("/fotos", servidorVivienda.CrearFoto)
		//r.Get("/fotos/{id}", servidorVivienda.ObtenerFoto)
		//r.Put("/fotos/{id}", servidorVivienda.ActualizarFoto)
		//r.Delete("/fotos/{id}", servidorVivienda.BorrarFoto)

		// Sectores: CRUD completo.
		//r.Get("/sectores", servidorVivienda.ListarSectores)
		//r.Post("/sectores", servidorVivienda.CrearSector)
		//r.Get("/sectores/{id}", servidorVivienda.ObtenerSector)
		//r.Put("/sectores/{id}", servidorVivienda.ActualizarSector)
		//r.Delete("/sectores/{id}", servidorVivienda.BorrarSector)

		// AplicarViviendas: CRUD completo.
		//r.Get("/aplicarviviendas", servidorVivienda.ListarAplicarViviendas)
		//r.Post("/aplicarviviendas", servidorVivienda.CrearAplicarVivienda)
		//r.Get("/aplicarviviendas/{id}", servidorVivienda.ObtenerAplicarVivienda)
		//r.Put("/aplicarviviendas/{id}", servidorVivienda.ActualizarAplicarVivienda)
		//r.Delete("/aplicarviviendas/{id}", servidorVivienda.BorrarAplicarVivienda)

		// Alimentación: CRUD completo.
		r.Get("/alimentaciones", servidorAlimentacion.ListarAlimentaciones)
		r.Get("/alimentaciones/{id}", servidorAlimentacion.BuscarAlimentacionesPorID)
		r.Post("/alimentaciones", servidorAlimentacion.CrearAlimentacion)
		r.Put("/alimentaciones/{id}", servidorAlimentacion.ActualizarAlimentacion)
		r.Delete("/alimentaciones/{id}", servidorAlimentacion.BorrarAlimentacion)

		// MenuDiario: CRUD completo.
		r.Get("/menudiarios", servidorAlimentacion.ListarMenuDiarios)
		r.Get("/menudiarios/{id}", servidorAlimentacion.BuscarMenuDiarioPorID)
		r.Post("/menudiarios", servidorAlimentacion.CrearMenuDiario)
		r.Put("/menudiarios/{id}", servidorAlimentacion.ActualizarMenuDiario)
		r.Delete("/menudiarios/{id}", servidorAlimentacion.BorrarMenuDiario)	
		
		// Platos: CRUD completo.
		r.Get("/platos", servidorAlimentacion.ListarPlatos)
		r.Get("/platos/{id}", servidorAlimentacion.BuscarPlatosPorID)
		r.Post("/platos", servidorAlimentacion.CrearPlato)
		r.Put("/platos/{id}", servidorAlimentacion.ActualizarPlato)	
		r.Delete("/platos/{id}", servidorAlimentacion.BorrarPlato)		

		// Resenas: CRUD completo.
		r.Get("/resenas", servidorAlimentacion.ListarResenas)
		r.Get("/resenas/{id}", servidorAlimentacion.BuscarResenasPorID)
		r.Post("/resenas", servidorAlimentacion.CrearResena)	
		r.Put("/resenas/{id}", servidorAlimentacion.ActualizarResena)	
		r.Delete("/resenas/{id}", servidorAlimentacion.BorrarResena)		
	
	})

	log.Println("Servidor escuchando en http://localhost:8080")	
	log.Fatal(http.ListenAndServe(":8080", r))	
}
