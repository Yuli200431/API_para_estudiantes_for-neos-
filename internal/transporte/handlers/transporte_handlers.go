package handlers

import (
	//"encoding/json"
	//"log"
	"net/http"
	//"strconv"
	//"strings"

	//"github.com/go-chi/chi/v5"

	//"for-neos-api/internal/transporte/models"
	"for-neos-api/internal/transporte/storage"
)

type Server struct {
	transporteStorage *storage.MemoriaTransporte
}

func NewServer(s *storage.MemoriaTransporte) *Server {
	return &Server{transporteStorage: s}
}

// Funcion para listar todas las rutas
// Recibe la peticion GET /rutas y devuelve un JSON con todas las rutas disponibles
// Atiende GET /api/v1/rutas
func (s *Server) ListarRutas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.transporteStorage.ListarRutas()

	ResponderJSON(w, http.StatusOK, rutas)
}

// Funcion para obtener una ruta por ID
// Recibe la peticion GET /rutas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
func (s *Server) ObtenerRutaPorID(w http.ResponseWriter, r *http.Request) {

}

// Funcion para crear una nueva ruta
// Recibe la peticion POST /rutas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva ruta, y devuelve un JSON con la ruta creada
func (s *Server) CrearRuta(w http.ResponseWriter, r *http.Request) {

}

// Funcion para actualizar una ruta
// Recibe la peticion PUT /rutas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la ruta, y devuelve un JSON con la ruta actualizada
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {

}

// Funcion para eliminar una ruta
// Recibe la peticion DELETE /rutas/:id y elimina la ruta correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
func (s *Server) EliminarRuta(w http.ResponseWriter, r *http.Request) {

}

// Funcion para buscar rutas por sector de origen y destino
// Recibe la peticion GET /rutas/buscar?origen=sectorOrigen&destino=sectorDestino
// y devuelve un JSON con las rutas que coinciden con los criterios de búsqueda
func (s *Server) BuscarRutasOrigenDestino(w http.ResponseWriter, r *http.Request) {

}
