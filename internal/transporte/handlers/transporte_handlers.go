package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/transporte/models"
	"for-neos-api/internal/transporte/storage"
)

type Server struct {
	transporteStorage *storage.MemoriaTransporte
}

func NewServer(s *storage.MemoriaTransporte) *Server {
	return &Server{transporteStorage: s}
}

// Funcion para listar todas las rutas
// Recibe la peticion GET /rutas y devuelve un JSON con todas las rutas disponibles
// Atiende GET /api/v1/transporte
func (s *Server) ListarRutas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.transporteStorage.ListarRutas()

	ResponderJSON(w, http.StatusOK, rutas)
}

// Funcion para obtener una ruta por ID
// Recibe la peticion GET /rutas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
// ObtenerRutaPorID atiende GET /api/v1/transporte/{id}.
func (s *Server) ObtenerRutaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	ruta, encontrado := s.transporteStorage.ObtenerRutaPorID(uint(id))
	if !encontrado {
		RespondError(w, http.StatusNotFound, "Ruta no encontrada")
		return
	}

	ResponderJSON(w, http.StatusOK, ruta)
		w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ruta); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// Funcion para crear una nueva ruta
// Recibe la peticion POST /rutas con un JSON en el cuerpo de la solicitud que contiene los datos de la
// nueva ruta, y devuelve un JSON con la ruta creada
// Atiende POST /api/v1/transporte
func (s *Server) AgregarRuta(w http.ResponseWriter, r *http.Request) {
	var nuevaRuta models.RutaTransporte
	if err := json.NewDecoder(r.Body).Decode(&nuevaRuta); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos de ruta inválidos"+err.Error())
		return
	}
	if strings.TrimSpace(nuevaRuta.NombreLinea) == "" {
		RespondError(w, http.StatusBadRequest, "El nombre de la línea es obligatorio")
		return
	}
	if strings.TrimSpace(nuevaRuta.FrecuenciaAprox) == "" {
		RespondError(w, http.StatusBadRequest, "La frecuencia aproximada es obligatoria")
		return
	}
	if nuevaRuta.Precio <= 0 {
		RespondError(w, http.StatusBadRequest, "El precio debe ser mayor a cero")
		return
	}
	if strings.TrimSpace(nuevaRuta.DescripcionRuta) == "" {
		RespondError(w, http.StatusBadRequest, "La descripción de la ruta es obligatoria")
		return
	}
	if nuevaRuta.CooperativaID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID de la cooperativa es obligatorio")
		return
	}
	if nuevaRuta.SectorOrigenID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del sector de origen es obligatorio")
		return
	}
	if nuevaRuta.SectorDestinoID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del sector de destino es obligatorio")
		return
	}
	if nuevaRuta.ProviderID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del proveedor es obligatorio")
		return
	}
	// Llamar al método AgregarRuta del almacenamiento
	rutaCreada := s.transporteStorage.AgregarRuta(nuevaRuta)
	ResponderJSON(w, http.StatusCreated, rutaCreada)
}

// Funcion para actualizar una ruta
// Recibe la peticion PUT /rutas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la ruta, y devuelve un JSON con la ruta actualizada
// Atiende PUT /api/v1/transporte/{id}.
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var rutaActualizada models.RutaTransporte
	if err := json.NewDecoder(r.Body).Decode(&rutaActualizada); err != nil {
		RespondError(w, http.StatusBadRequest, "Datos de ruta inválidos: "+err.Error())
		return
	}
	if strings.TrimSpace(rutaActualizada.NombreLinea) == "" {
		RespondError(w, http.StatusBadRequest, "El nombre de la línea es obligatorio")
		return
	}
	if strings.TrimSpace(rutaActualizada.FrecuenciaAprox) == "" {
		RespondError(w, http.StatusBadRequest, "La frecuencia aproximada es obligatoria")
		return
	}
	if rutaActualizada.Precio <= 0 {
		RespondError(w, http.StatusBadRequest, "El precio debe ser mayor a cero")
		return
	}
	if strings.TrimSpace(rutaActualizada.DescripcionRuta) == "" {
		RespondError(w, http.StatusBadRequest, "La descripción de la ruta es obligatoria")
		return
	}
	if rutaActualizada.CooperativaID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID de la cooperativa es obligatorio")
		return
	}
	if rutaActualizada.SectorOrigenID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del sector de origen es obligatorio")
		return
	}
	if rutaActualizada.SectorDestinoID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del sector de destino es obligatorio")
		return
	}
	if rutaActualizada.ProviderID == 0 {
		RespondError(w, http.StatusBadRequest, "El ID del proveedor es obligatorio")
		return
	}

	// Llamar al método ActualizarRuta del almacenamiento
	ruta, actualizado := s.transporteStorage.ActualizarRuta(uint(id), rutaActualizada)
	if !actualizado {
		RespondError(w, http.StatusNotFound, "Ruta no encontrada")
		return
	}
	ResponderJSON(w, http.StatusOK, ruta)
}

// Funcion para eliminar una ruta
// Recibe la peticion DELETE /rutas/:id y elimina la ruta correspondiente al ID proporcionado,
// devolviendo un mensaje de éxito o error
// Atiende DELETE /api/v1/transporte/{id}.
func (s *Server) EliminarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if !s.transporteStorage.EliminarRuta(uint(id)) {
		RespondError(w, http.StatusNotFound, "Ruta no encontrada")
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

