package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"strings"
	"strconv"
     
	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/storage"
)

type Server struct {
	storage *storage.Memoria
}

func NewServer(storage *storage.Memoria) *Server {
	return &Server{
		storage: storage,
	}
}
// ListarAlimentaciones maneja la solicitud GET /alimentaciones
func (s *Server) ListarAlimentaciones(w http.ResponseWriter, _ *http.Request) {
alimentaciones := s.storage.ListarAlimentaciones()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(alimentaciones); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

// BuscarAlimentacionesPorID maneja la solicitud GET /alimentaciones/{id}
func (s *Server) BuscarAlimentacionesPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	alimentacion, encontrado := s.storage.BuscarAlimentacionPorID(id)
	if !encontrado {
		http.Error(w, "Alimentación no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(alimentacion); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

//CrearAlimentacion maneja la solicitud POST/api/c /alimentaciones	

func (s *Server) CrearAlimentacion(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Alimentacion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos:"+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevo.NombreLocal) == "" {
		http.Error(w, "El nombre del local es obligatorio", http.StatusBadRequest)
		return
	}
	creado := s.storage.CrearAlimentacion(nuevo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

// ACTUALIZAR ALIMENTACION maneja la solicitud PUT/api/v /alimentaciones/{id}	
func (s *Server) ActualizarAlimentacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var datos models.Alimentacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "Datos inválidos:"+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.NombreLocal) == "" {
		http.Error(w, "El nombre del local es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.ID != 0 && datos.ID != id {
		http.Error(w, "El ID en el cuerpo no coincide con el ID de la URL", http.StatusBadRequest)
		return
	}
	actualizado, encontrado := s.storage.ActualizarAlimentacion(id, datos)
	if !encontrado {
		http.Error(w, "Alimentación no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}	
}

// BorrarAlimentacion maneja la solicitud DELETE /alimentaciones/{id}
func (s *Server) BorrarAlimentacion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if !s.storage.BorrarAlimentacion(id) {
		http.Error(w, "Alimentación no encontrada", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
