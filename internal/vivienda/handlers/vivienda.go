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
// ListarViviendas atiende GET /api/v1/viviendas.
func (s *Server) ListarViviendas(w http.ResponseWriter, _ *http.Request) {
	viviendas := s.Viviendas.Listar()
	RespondJSON(w, http.StatusOK, viviendas)
}

// ObtenerVivienda atiende GET /api/v1/viviendas/{id}.
func (s *Server) ObtenerVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	vivienda, err := s.Viviendas.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, vivienda)
}

// CrearVivienda atiende POST /api/v1/categorias.
func (s *Server) CrearVivienda(w http.ResponseWriter, r *http.Request) {
	var nueva models.Vivienda
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Viviendas.Crear(nueva)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarVivienda atiende PUT /api/v1/viviendas/{id}.
func (s *Server) ActualizarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.Vivienda
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Viviendas.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarVivienda atiende DELETE /api/v1/viviendas/{id}.
func (s *Server) BorrarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	if err := s.Viviendas.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
