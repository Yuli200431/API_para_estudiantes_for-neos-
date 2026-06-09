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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rutas); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para obtener una ruta por ID
// Recibe la peticion GET /rutas/:id y devuelve un JSON con la ruta correspondiente al ID proporcionado
// ObtenerRutaPorID atiende GET /api/v1/transporte/{id}.
func (s *Server) ObtenerRutaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ruta, encontrado := s.transporteStorage.ObtenerRutaPorID(uint(id))
	if !encontrado {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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
		http.Error(w, "Datos de ruta inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaRuta.NombreLinea) == "" {
		http.Error(w, "El nombre de la línea es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaRuta.FrecuenciaAprox) == "" {
		http.Error(w, "La frecuencia aproximada es obligatoria", http.StatusBadRequest)
		return
	}
	if nuevaRuta.Precio <= 0 {
		http.Error(w, "El precio debe ser mayor a cero", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevaRuta.DescripcionRuta) == "" {
		http.Error(w, "La descripción de la ruta es obligatoria", http.StatusBadRequest)
		return
	}
	if nuevaRuta.CooperativaID == 0 {
		http.Error(w, "El ID de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if nuevaRuta.SectorOrigenID == 0 {
		http.Error(w, "El ID del sector de origen es obligatorio", http.StatusBadRequest)
		return
	}
	if nuevaRuta.SectorDestinoID == 0 {
		http.Error(w, "El ID del sector de destino es obligatorio", http.StatusBadRequest)
		return
	}
	if nuevaRuta.ProviderID == 0 {
		http.Error(w, "El ID del proveedor es obligatorio", http.StatusBadRequest)
		return
	}
	// Llamar al método AgregarRuta del almacenamiento
	rutaCreada := s.transporteStorage.AgregarRuta(nuevaRuta)

		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(rutaCreada); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// Funcion para actualizar una ruta
// Recibe la peticion PUT /rutas/:id con un JSON en el cuerpo de la solicitud que contiene
// los datos actualizados de la ruta, y devuelve un JSON con la ruta actualizada
// Atiende PUT /api/v1/transporte/{id}.
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var rutaActualizada models.RutaTransporte
	if err := json.NewDecoder(r.Body).Decode(&rutaActualizada); err != nil {
		http.Error(w, "Datos de ruta inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(rutaActualizada.NombreLinea) == "" {
		http.Error(w, "El nombre de la línea es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(rutaActualizada.FrecuenciaAprox) == "" {
		http.Error(w, "La frecuencia aproximada es obligatoria", http.StatusBadRequest)
		return
	}
	if rutaActualizada.Precio <= 0 {
		http.Error(w, "El precio debe ser mayor a cero", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(rutaActualizada.DescripcionRuta) == "" {
		http.Error(w, "La descripción de la ruta es obligatoria", http.StatusBadRequest)
		return
	}
	if rutaActualizada.CooperativaID == 0 {
		http.Error(w, "El ID de la cooperativa es obligatorio", http.StatusBadRequest)
		return
	}
	if rutaActualizada.SectorOrigenID == 0 {
		http.Error(w, "El ID del sector de origen es obligatorio", http.StatusBadRequest)
		return
	}
	if rutaActualizada.SectorDestinoID == 0 {
		http.Error(w, "El ID del sector de destino es obligatorio", http.StatusBadRequest)
		return
	}
	if rutaActualizada.ProviderID == 0 {
		http.Error(w, "El ID del proveedor es obligatorio", http.StatusBadRequest)
		return
	}

	// Llamar al método ActualizarRuta del almacenamiento
	actualizado, encontrado := s.transporteStorage.ActualizarRuta(uint(id), rutaActualizada)
	if !encontrado {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
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

	if !s.transporteStorage.EliminarRuta(uint(id)) {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

