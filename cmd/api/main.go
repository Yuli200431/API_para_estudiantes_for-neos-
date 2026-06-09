package main
import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"for-neos-api/internal/alimentacion/handlers"
	"for-neos-api/internal/alimentacion/storage"
)

func main() {
	storage := storage.NewMemoria()
	storage.SeedAlimentaciones()
	server := handlers.NewServer(storage)

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
	r.Get("/alimentaciones", server.ListarAlimentaciones)
	r.Get("/alimentaciones/{id}", server.BuscarAlimentacionesPorID)
	r.Post("/alimentaciones", server.CrearAlimentacion)
	r.Put("/alimentaciones/{id}", server.ActualizarAlimentacion)
	r.Delete("/alimentaciones/{id}", server.BorrarAlimentacion)
})	
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}