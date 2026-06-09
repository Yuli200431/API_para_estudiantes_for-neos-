// Package handlers contiene los handlers HTTP de la API.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/vivienda/models"
)

// ListarSectores atiende GET /api/v1/sectores.
func (s *Server) ListarSectores(w http.ResponseWriter, _ *http.Request) {
	sectores := s.Storage.ListarSectores()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sectores); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerSector atiende GET /api/v1/sectores/{id}.
func (s *Server) ObtenerSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	sector, encontrado := s.Storage.BuscarSectorPorID(id)
	if !encontrado {
		http.Error(w, "Sector no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sector); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearSector atiende POST /api/v1/sectores.
func (s *Server) CrearSector(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Sector
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		http.Error(w, "el campo nombre es obligatorio", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearSector(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarSector atiende PUT /api/v1/sectores/{id}.
func (s *Server) ActualizarSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Sector
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		http.Error(w, "el campo nombre es obligatorio", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.ActualizarSector(id, datos)
	if !encontrado {
		http.Error(w, "Sector no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// BorrarSector atiende DELETE /api/v1/sectores/{id}.
func (s *Server) BorrarSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarSector(id) {
		http.Error(w, "Sector no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
