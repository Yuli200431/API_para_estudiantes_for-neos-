package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/transporte/models"
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
// ObtenerRutaPorID atiende GET /api/v1/rutas/{id}.
func (s *Server) ObtenerRutaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	ruta, encontrado := s.transporteStorage.ObtenerRutaPorID(uint(id))
	if !encontrado {
		RespondError(w, http.StatusNotFound, "Ruta no encontrada")
		return
	}

	ResponderJSON(w, http.StatusOK, ruta)
		w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ruta); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// Funcion para crear una nueva ruta
// Recibe la peticion POST /rutas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva ruta, y devuelve un JSON con la ruta creada
func (s *Server) AgregarRuta(w http.ResponseWriter, r *http.Request) {
	var nuevaRuta models.Transporte
	if err := json.NewDecoder(r.Body).Decode(&nuevaRuta); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos de ruta inválidos"+err.Error())
		return
	}
	if strings.TrimSpace(nuevaRuta.NombreLinea) == "" {
		RespondError(w, http.StatusBadRequest, "El nombre de la línea es obligatorio")
		return
	}
	if strings.TrimSpace(nuevaRuta.Cooperativa) == "" {
		RespondError(w, http.StatusBadRequest, "La cooperativa es obligatoria")
		return
	}
	if strings.TrimSpace(nuevaRuta.SectorOrigen) == "" {
		RespondError(w, http.StatusBadRequest, "El sector de origen es obligatorio")
		return
	}
	if strings.TrimSpace(nuevaRuta.SectorDestino) == "" {
		RespondError(w, http.StatusBadRequest, "El sector de destino es obligatorio")
		return
	}
	if strings.TrimSpace(nuevaRuta.SectoresRecorridos) == "" {
		RespondError(w, http.StatusBadRequest, "Los sectores recorridos son obligatorios")
		return
	}
	if strings.TrimSpace(nuevaRuta.FrecuenciaAprox) == "" {
		RespondError(w, http.StatusBadRequest, "La frecuencia aproximada es obligatoria")
		return
	}
	if nuevaRuta.Precio <= 0 {
		RespondError(w, http.StatusBadRequest, "El precio debe ser mayor a cero")
		return
	}
	if strings.TrimSpace(nuevaRuta.DescripcionRuta) == "" {
		RespondError(w, http.StatusBadRequest, "La descripción de la ruta es obligatoria")
		return
	}
	if nuevaRuta.ProviderID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del proveedor es obligatorio")
		return
	}
	// Llamar al método AgregarRuta del almacenamiento
	rutaCreada := s.transporteStorage.AgregarRuta(nuevaRuta)
	ResponderJSON(w, http.StatusCreated, rutaCreada)
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
