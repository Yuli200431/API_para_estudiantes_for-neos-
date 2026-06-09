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

// ListarFotos atiende GET /api/v1/fotos.
func (s *Server) ListarFotos(w http.ResponseWriter, _ *http.Request) {
	fotos := s.Storage.ListarFotos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(fotos); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// ObtenerFoto atiende GET /api/v1/fotos/{id}.
func (s *Server) ObtenerFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	foto, encontrado := s.Storage.BuscarFotoPorID(id)
	if !encontrado {
		http.Error(w, "Foto no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(foto); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// CrearFoto atiende POST /api/v1/fotos.
func (s *Server) CrearFoto(w http.ResponseWriter, r *http.Request) {
	var nueva models.Foto
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.URL) == "" {
		http.Error(w, "El campo URL es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.ViviendaID < 0 {
		http.Error(w, "El Id de la Vivienda debe ser un número entero positivo", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearFoto(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// ActualizarFoto atiende PUT /api/v1/fotos/{id}.
func (s *Server) ActualizarFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Foto
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.URL) == "" {
		http.Error(w, "El campo URL es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.ViviendaID < 0 {
		http.Error(w, "El Id de la Vivienda debe ser un número entero positivo", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.ActualizarFoto(id, datos)
	if !encontrado {
		http.Error(w, "Foto no fueencontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// BorrarFoto atiende DELETE /api/v1/fotos/{id}.
func (s *Server) BorrarFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarFoto(id) {
		http.Error(w, "Foto no fue encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
