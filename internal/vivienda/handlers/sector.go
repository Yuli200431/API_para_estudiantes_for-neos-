// Package handlers contiene los handlers HTTP de la API.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/vivienda/models"
)

// ListarSectores atiende GET /api/v1/sectores.
func (s *Server) ListarSectores(w http.ResponseWriter, _ *http.Request) {
	sectores := s.Sectores.Listar()
	RespondJSON(w, http.StatusOK, sectores)
}

// ObtenerSector atiende GET /api/v1/sectores/{id}.
func (s *Server) ObtenerSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	sector, err := s.Sectores.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, sector)
}

// CrearSector atiende POST /api/v1/sectores.
func (s *Server) CrearSector(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Sector
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "El campo Nombre no puede estar vacío")
		return
	}

	creado, err := s.Sectores.Crear(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarSector atiende PUT /api/v1/sectores/{id}.
func (s *Server) ActualizarSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.Sector
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "El campo Nombre no puede estar vacío")
		return
	}

	actualizado, err := s.Sectores.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarSector atiende DELETE /api/v1/sectores/{id}.
func (s *Server) BorrarSector(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	if err := s.Sectores.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
