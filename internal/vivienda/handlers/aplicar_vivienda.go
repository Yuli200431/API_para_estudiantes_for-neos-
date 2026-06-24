// Package handlers contiene los handlers HTTP de la API.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/vivienda/models"
)

// Server agrupa las dependencias compartidas por los handlers.
// Recibe el storage por inyección de dependencias desde main.
// ListarAplicarViviendas atiende GET /api/v1/aplicarviviendas.
func (s *Server) ListarAplicarViviendas(w http.ResponseWriter, _ *http.Request) {
	aplicarviviendas := s.Aplicar.Listar()
	RespondJSON(w, http.StatusOK, aplicarviviendas)
}

// ObtenerAplicarVivienda atiende GET /api/v1/AplicarViviendas/{id}.
func (s *Server) ObtenerAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	vivienda, err := s.Aplicar.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, vivienda)
}

// CrearAplicarVivienda atiende POST /api/v1/aplicarviviendas.
func (s *Server) CrearAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	var nueva models.AplicarVivienda
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Aplicar.Crear(nueva)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarAplicarVivienda atiende PUT /api/v1/AplicarViviendas/{id}.
func (s *Server) ActualizarAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.AplicarVivienda
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Aplicar.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarAplicarVivienda atiende DELETE /api/v1/AplicarViviendas/{id}.
func (s *Server) BorrarAplicarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	if err := s.Aplicar.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, nil)
}
