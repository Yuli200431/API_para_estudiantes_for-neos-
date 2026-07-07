package handlers

import (
	"encoding/json"
	"net/http"

	"for-neos-api/internal/transporte/models"
)

// Funcion para listar todas las paradas
// Recibe la peticion GET /paradas y devuelve un JSON con todas las paradas disponibles
// Atiende GET /api/v1/paradas
func (s *Server) ListarParadas(w http.ResponseWriter, _ *http.Request) {
	paradas := s.Paradas.ListarParadas()
	RespondJSON(w, http.StatusOK, paradas)
}

// Funcion para obtener una paradas por ID
// Recibe la peticion GET /paradas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
// ObtenerParadaPorID atiende GET /api/v1/parada/{id}.
func (s *Server) ObtenerParadaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	paradas, err := s.Paradas.ObtenerParadas(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, paradas)
}

// Funcion para crear una nueva parada
// Recibe la peticion POST /parada con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva parada, y devuelve un JSON con la parada creada
// Atiende POST /api/v1/parada
func (s *Server) AgregarParada(w http.ResponseWriter, r *http.Request) {
	var nuevaParada models.ParadaBus
	if err := json.NewDecoder(r.Body).Decode(&nuevaParada); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	paradaCreada, err := s.Paradas.CrearParada(nuevaParada)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, paradaCreada)
}

// Funcion para actualizar una parada
// Recibe la peticion PUT /parada/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la parada, y devuelve un JSON con la parada actualizada
// Atiende PUT /api/v1/parada/{id}.
func (s *Server) ActualizarParada(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "El Id debe ser un número entero")
		return
	}

	var datos models.ParadaBus
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Paradas.ActualizarParadas(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// Funcion para eliminar una parada
// Recibe la peticion DELETE /parada/:id y elimina la parada correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/parada/{id}.
func (s *Server) EliminarParada(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Paradas.BorrarParadas(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, nil)
}
