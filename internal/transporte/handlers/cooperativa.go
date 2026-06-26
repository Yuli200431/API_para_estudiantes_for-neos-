package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/transporte/models"
)

// Funcion para listar todas las cooperativas
// Recibe la peticion GET /cooperativas y devuelve un JSON con todas las cooperativas disponibles
// Atiende GET /api/v1/transporte/cooperativas

func (s *Server) ListarCooperativas(w http.ResponseWriter, _ *http.Request) {
	cooperativas := s.Cooperativas.Listar()
	RespondJSON(w, http.StatusOK, cooperativas)
}

// Funcion para obtener una cooperativa por ID
// Recibe la peticion GET /cooperativas/:id y devuelve un JSON con la cooperativa correspondiente al ID proporcionado
// ObtenerCooperativaPorID atiende GET /api/v1/transporte/{id}.
func (s *Server) ObtenerCooperativaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	cooperativas, err := s.Cooperativas.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cooperativas)
}

// Funcion para crear una nueva cooperativa
// Recibe la peticion POST /cooperativas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva cooperativa, y devuelve un JSON con la cooperativa creada
// Atiende POST /api/v1/cooperativa
func (s *Server) AgregarCooperativa(w http.ResponseWriter, r *http.Request) {
	var nuevaCooperativa models.Cooperativa
	if err := json.NewDecoder(r.Body).Decode(&nuevaCooperativa); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Cooperativas.Crear(nuevaCooperativa)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// Funcion para actualizar una cooperativa
// Recibe la peticion PUT /cooperativas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la cooperativa, y devuelve un JSON con la cooperativa actualizada
// Atiende PUT /api/v1/cooperativa/{id}.
func (s *Server) ActualizarCooperativa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.Cooperativa

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Cooperativas.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// Funcion para eliminar una ruta
// Recibe la peticion DELETE /rutas/:id y elimina la ruta correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/cooperativa/{id}.
func (s *Server) EliminarCooperativa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Cooperativas.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, nil)
}
