package handlers

import (
	"encoding/json"
	"net/http"

	"for-neos-api/internal/alimentacion/models"
)

func (s *Server) ListarMenuDiarios(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.MenuDiario.Listar())
}

func (s *Server) BuscarMenuDiarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	menuDiario, err := s.MenuDiario.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, menuDiario)
}

func (s *Server) CrearMenuDiario(w http.ResponseWriter, r *http.Request) {
	var nuevo models.MenuDiario
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos inválidos: "+err.Error())
		return
	}
	creado, err := s.MenuDiario.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarMenuDiario(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var datos models.MenuDiario
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos inválidos: "+err.Error())
		return
	}
	actualizado, err := s.MenuDiario.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarMenuDiario(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := s.MenuDiario.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
