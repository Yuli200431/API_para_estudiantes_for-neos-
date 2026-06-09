package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"log"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/storage"
	"github.com/go-chi/chi/v5"
	"encoding/json"
)

// ListarPlatos maneja la solicitud GET /platos
func (s *Server) ListarPlatos(w http.ResponseWriter, _ *http.Request) {
	platos := s.storage.ListarPlatos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(platos); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

// BuscarPlatosPorID maneja la solicitud GET /platos/{id}
func (s *Server) BuscarPlatosPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	plato, encontrado := s.storage.BuscarPlatoPorID(id)
	if !encontrado {
		http.Error(w, "Plato no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(plato); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

//CrearPlato maneja la solicitud POST/api/c /platos	

func (s *Server) CrearPlato(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Plato
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)	
		return
	}	
	if strings.TrimSpace(nuevo.NombrePlato) == "" {
		http.Error(w, "El nombre del plato es obligatorio", http.StatusBadRequest)	
		return
	}
	if strings.TrimSpace(nuevo.Descripcion) == "" {
		http.Error(w, "La descripción es obligatoria", http.StatusBadRequest)	
		return
	}
	if strings.TrimSpace(nuevo.Categoria) == "" {
		http.Error(w, "La categoria es obligatoria", http.StatusBadRequest)	
		return
	}
	creado := s.storage.CrearPlato(nuevo)
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusCreated)	
	if err := json.NewEncoder(w).Encode(creado); err != nil {	
		log.Printf("Error al codificar la respuesta: %v", err)	

	}	
}	
// ACTUALIZAR plato maneja la solicitud PUT/api/v /platos/{id}	
func (s *Server) ActualizarPlato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))	
	if err != nil {	
		http.Error(w, "ID inválido", http.StatusBadRequest)	
		return	
	}	
	var datos models.Plato	
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {	
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)	
		return	
	}
	if strings.TrimSpace(datos.NombrePlato) == "" {	
		http.Error(w, "El nombre del plato es obligatorio", http.StatusBadRequest)	
		return	
		}	
	if strings.TrimSpace(datos.Descripcion) == "" {	
		http.Error(w, "La descripción es obligatoria", http.StatusBadRequest)	
		return	
		}	
	if strings.TrimSpace(datos.Categoria) == "" {	
		http.Error(w, "La categoria es obligatoria", http.StatusBadRequest)	
		return	
		}	
	actualizado, encontrado := s.storage.ActualizarPlato(id, datos)	
	if !encontrado {	
		http.Error(w, "Plato no encontrada", http.StatusNotFound)	
		return	
		}	
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusOK)	
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {	
		log.Printf("Error al codificar la respuesta: %v", err)	
	}	
}

// BorrarPlato maneja la solicitud DELETE /platos/{id}
func (s *Server) BorrarPlato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))	
	if err != nil {	
		http.Error(w, "ID inválido", http.StatusBadRequest)	
		return	
		}	
	if !s.storage.BorrarPlato(id) {	
		http.Error(w, "Plato no encontrada", http.StatusNotFound)	
		return	
		}	
	w.WriteHeader(http.StatusNoContent)	
}