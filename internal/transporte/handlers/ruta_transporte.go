package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"


	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/transporte/models"
)


// Funcion para listar todas las rutas
// Recibe la peticion GET /rutas y devuelve un JSON con todas las rutas disponibles
// Atiende GET /api/v1/transporte
func (s *Server) ListarRutas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.Rutas.ListarRuta()
	RespondJSON(w, http.StatusOK, rutas)
}

// Funcion para obtener una ruta por ID
// Recibe la peticion GET /rutas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
// ObtenerRutaPorID atiende GET /api/v1/transporte/{id}.
func (s *Server) ObtenerRutaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	rutas, err := s.Rutas.ObtenerRutas(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, rutas)
}

// Funcion para crear una nueva ruta
// Recibe la peticion POST /rutas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva ruta, y devuelve un JSON con la ruta creada
// Atiende POST /api/v1/transporte
func (s *Server) AgregarRuta(w http.ResponseWriter, r *http.Request) {
	var nuevaRuta models.RutaTransporte
	if err := json.NewDecoder(r.Body).Decode(&nuevaRuta); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creada, err := s.Rutas.CrearRuta(nuevaRuta)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// Funcion para actualizar una ruta
// Recibe la peticion PUT /rutas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la ruta, y devuelve un JSON con la ruta actualizada
// Atiende PUT /api/v1/transporte/{id}.
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.RutaTransporte

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Rutas.ActualizarRuta(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// Funcion para eliminar una ruta
// Recibe la peticion DELETE /rutas/:id y elimina la ruta correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/transporte/{id}.
func (s *Server) EliminarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Rutas.BorrarRuta(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, nil)
}
