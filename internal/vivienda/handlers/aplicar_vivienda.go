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

// ListarAplicarViviendas atiende GET /api/v1/aplicarviviendas.
func (s *Server) ListarAplicarViviendas(w http.ResponseWriter, _ *http.Request) {
	aplicarviviendas := s.Storage.ListarAplicarViviendas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aplicarviviendas); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// ObtenerAplicarVivienda atiende GET /api/v1/AplicarViviendas/{id}.
func (s *Server) ObtenerAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	aplicarvivienda, encontrado := s.Storage.BuscarAplicarViviendasPorID(id)
	if !encontrado {
		http.Error(w, "AplicarVivienda no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aplicarvivienda); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// CrearAplicarVivienda atiende POST /api/v1/aplicarviviendas.
func (s *Server) CrearAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	var nuevo models.AplicarVivienda
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if nuevo.EstudianteID < 0 {
		http.Error(w, "El Id del Estudiante debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if nuevo.ViviendaID < 0 {
		http.Error(w, "El Id de la Vivienda debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevo.Estado) == "" {
		http.Error(w, "El campo Estado es obligatorio", http.StatusBadRequest)
		return
	}
	if nuevo.Estado != "Pendiente" && nuevo.Estado != "Aceptada" && nuevo.Estado != "Rechazada" {
		http.Error(w, "El campo Estado debe ser 'Pendiente', 'Aceptada' o 'Rechazada'", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearAplicarVivienda(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// ActualizarAplicarVivienda atiende PUT /api/v1/AplicarViviendas/{id}.
func (s *Server) ActualizarAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.AplicarVivienda
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if datos.EstudianteID < 0 {
		http.Error(w, "El Id del Estudiante debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if datos.ViviendaID < 0 {
		http.Error(w, "El Id de la Vivienda debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Estado) == "" {
		http.Error(w, "El campo Estado es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Estado != "Pendiente" && datos.Estado != "Aceptada" && datos.Estado != "Rechazada" {
		http.Error(w, "El campo Estado debe ser 'Pendiente', 'Aceptada' o 'Rechazada'", http.StatusBadRequest)
		return
	}
	actualizado, encontrado := s.Storage.ActualizarAplicarVivienda(id, datos)
	if !encontrado {
		http.Error(w, "AplicarVivienda no fue encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// BorrarAplicarVivienda atiende DELETE /api/v1/AplicarViviendas/{id}.
func (s *Server) BorrarAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarAplicarVivienda(id) {
		http.Error(w, "AplicarVivienda no fue encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
