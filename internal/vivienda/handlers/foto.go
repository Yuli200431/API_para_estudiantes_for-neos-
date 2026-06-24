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
// ListarFotos atiende GET /api/v1/Fotos.
func (s *Server) ListarFotos(w http.ResponseWriter, _ *http.Request) {
	fotos := s.Fotos.Listar()
	RespondJSON(w, http.StatusOK, fotos)
}

// ObtenerFoto atiende GET /api/v1/Fotos/{id}.
func (s *Server) ObtenerFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	foto, err := s.Fotos.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, foto)
}

// CrearFoto atiende POST /api/v1/Fotos.
func (s *Server) CrearFoto(w http.ResponseWriter, r *http.Request) {
	var nueva models.Foto
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Fotos.Crear(nueva)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarFoto atiende PUT /api/v1/Fotos/{id}.
func (s *Server) ActualizarFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.Foto
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Fotos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarFoto atiende DELETE /api/v1/Fotos/{id}.
func (s *Server) BorrarFoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	if err := s.Fotos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
