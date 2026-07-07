package handlers

import (
	"encoding/json"
	"net/http"

	"for-neos-api/internal/alimentacion/models"
)

func (s *Server) ListarAlimentaciones(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Alimentacion.Listar())
}

func (s *Server) BuscarAlimentacionesPorID(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	alimentacion, err := s.Alimentacion.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, alimentacion)
}

func (s *Server) CrearAlimentacion(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Alimentacion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos inválidos: "+err.Error())
		return
	}
	creado, err := s.Alimentacion.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarAlimentacion(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var datos models.Alimentacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos inválidos: "+err.Error())
		return
	}
	actualizado, err := s.Alimentacion.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarAlimentacion(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := s.Alimentacion.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
