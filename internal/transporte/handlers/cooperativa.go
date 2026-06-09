package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/transporte/models"

)


// Funcion para listar todas las cooperativas
// Recibe la peticion GET /cooperativas y devuelve un JSON con todas las cooperativas disponibles
// Atiende GET /api/v1/transporte/cooperativas

func (s *Server) ListarCooperativas(w http.ResponseWriter, _ *http.Request) {
	cooperativas := s.transporteStorage.ListarCooperativas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cooperativas); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para obtener una cooperativa por ID
// Recibe la peticion GET /cooperativas/:id y devuelve un JSON con la cooperativa correspondiente al ID proporcionado
// ObtenerCooperativaPorID atiende GET /api/v1/transporte/{id}.
func (s *Server) ObtenerCooperativaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cooperativa, encontrado := s.transporteStorage.ObtenerCooperativaPorID(uint(id))
	if !encontrado {
		http.Error(w, "Cooperativa no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cooperativa); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// Funcion para crear una nueva cooperativa
// Recibe la peticion POST /cooperativas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva cooperativa, y devuelve un JSON con la cooperativa creada
// Atiende POST /api/v1/cooperativa
func (s *Server) AgregarCooperativa(w http.ResponseWriter, r *http.Request) {
	var nuevaCooperativa models.Cooperativa
	if err := json.NewDecoder(r.Body).Decode(&nuevaCooperativa); err != nil {
		http.Error(w, "Datos de cooperativa inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaCooperativa.Nombre) == "" {
		http.Error(w, "El nombre de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaCooperativa.Telefono) == "" {
		http.Error(w, "El teléfono de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaCooperativa.Descripcion) == "" {
		http.Error(w, "La descripción de la cooperativa es obligatoria", http.StatusBadRequest)
		return
	}
	
	// Llamar al método AgregarCooperativa del almacenamiento
	cooperativaCreada := s.transporteStorage.AgregarCooperativa(nuevaCooperativa)

		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cooperativaCreada); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para actualizar una cooperativa
// Recibe la peticion PUT /cooperativas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la cooperativa, y devuelve un JSON con la cooperativa actualizada
// Atiende PUT /api/v1/cooperativa/{id}.
func (s *Server) ActualizarCooperativa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var cooperativaActualizada models.Cooperativa
	if err := json.NewDecoder(r.Body).Decode(&cooperativaActualizada); err != nil {
		http.Error(w, "Datos de cooperativa inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(cooperativaActualizada.Nombre) == "" {
		http.Error(w, "El nombre de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(cooperativaActualizada.Telefono) == "" {
		http.Error(w, "El teléfono de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(cooperativaActualizada.Descripcion) == "" {
		http.Error(w, "La descripción de la cooperativa es obligatoria", http.StatusBadRequest)
		return
	}
	
	// Llamar al método ActualizarCooperativa del almacenamiento
	actualizado, encontrado := s.transporteStorage.ActualizarCooperativa(uint(id), cooperativaActualizada)
	if !encontrado {
		http.Error(w, "Cooperativa no encontrada", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para eliminar una ruta
// Recibe la peticion DELETE /rutas/:id y elimina la ruta correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/cooperativa/{id}.
func (s *Server) EliminarCooperativa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !s.transporteStorage.EliminarCooperativa(uint(id)) {
		http.Error(w, "Cooperativa no encontrada", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}