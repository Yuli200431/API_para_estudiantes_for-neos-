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


// Funcion para listar todas las paradas
// Recibe la peticion GET /paradas y devuelve un JSON con todas las paradas disponibles
// Atiende GET /api/v1/paradas
func (s *Server) ListarParadas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.transporteStorage.ListarParadas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rutas); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para obtener una paradas por ID
// Recibe la peticion GET /paradas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
// ObtenerParadaPorID atiende GET /api/v1/parada/{id}.
func (s *Server) ObtenerParadaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	parada, encontrado := s.transporteStorage.ObtenerParadaPorID(uint(id))
	if !encontrado {
		http.Error(w, "Parada no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(parada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// Funcion para crear una nueva parada
// Recibe la peticion POST /parada con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva parada, y devuelve un JSON con la parada creada
// Atiende POST /api/v1/parada
func (s *Server) AgregarParada(w http.ResponseWriter, r *http.Request) {
	var nuevaParada models.ParadaBus
	if err := json.NewDecoder(r.Body).Decode(&nuevaParada); err != nil {
		http.Error(w, "Datos de parada inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaParada.NombreParada) == "" {
		http.Error(w, "El nombre de la parada de bus es obligatorio", http.StatusBadRequest)
		return
	}
		if strings.TrimSpace(nuevaParada.Direccion) == "" {
		http.Error(w, "La direccion de la parada es obligatoria", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaParada.Descripcion) == "" {
		http.Error(w, "La descripción de la parada es obligatoria", http.StatusBadRequest)
		return
	}
	// Llamar al método AgregarParada del almacenamiento
	paradaCreada := s.transporteStorage.AgregarParada(nuevaParada)

		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(paradaCreada); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para actualizar una parada
// Recibe la peticion PUT /parada/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la parada, y devuelve un JSON con la parada actualizada
// Atiende PUT /api/v1/parada/{id}.
func (s *Server) ActualizarParada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var paradaActualizada models.ParadaBus
	if err := json.NewDecoder(r.Body).Decode(&paradaActualizada); err != nil {
		http.Error(w, "Datos de parada inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(paradaActualizada.NombreParada) == "" {
		http.Error(w, "El nombre de la parada es obligatorio", http.StatusBadRequest)
		return
	}
		if strings.TrimSpace(paradaActualizada.Direccion) == "" {
		http.Error(w, "La direccion de la parada es obligatoria", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(paradaActualizada.Descripcion) == "" {
		http.Error(w, "La descripción de la parada es obligatoria", http.StatusBadRequest)
		return
	}

	// Llamar al método ActualizarParada del almacenamiento
	actualizado, encontrado := s.transporteStorage.ActualizarParada(uint(id), paradaActualizada)
	if !encontrado {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para eliminar una parada
// Recibe la peticion DELETE /parada/:id y elimina la parada correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/parada/{id}.
func (s *Server) EliminarParada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !s.transporteStorage.EliminarParada(uint(id)) {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}