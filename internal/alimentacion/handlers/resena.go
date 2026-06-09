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

// ListarResenas maneja la solicitud GET /resenas
func (s *Server) ListarResenas(w http.ResponseWriter, _ *http.Request) {
	resenas := s.storage.ListarResenas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resenas); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

// BuscarResenasPorID maneja la solicitud GET /resenas/{id}
func (s *Server) BuscarResenasPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	resena, encontrado := s.storage.BuscarResenaPorID(id)
	if !encontrado {
		http.Error(w, "Resena no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resena); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

//CrearResena maneja la solicitud POST/api/c /resenas	

func (s *Server) CrearResena(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Resena
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)	
		return
	}	
	if strings.TrimSpace(nuevo.Comentario) == "" {
		http.Error(w, "El comentario es obligatorio", http.StatusBadRequest)	
		return
	}
	creado := s.storage.CrearResena(nuevo)
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusCreated)	
	if err := json.NewEncoder(w).Encode(creado); err != nil {	
		log.Printf("Error al codificar la respuesta: %v", err)	

	}	
}		
// BorrarResena maneja la solicitud DELETE /resenas/{id}
func (s *Server) BorrarResena(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))	
	if err != nil {	
		http.Error(w, "ID inválido", http.StatusBadRequest)	
		return	
		}	
	if !s.storage.BorrarResena(id) {	
		http.Error(w, "Resena no encontrada", http.StatusNotFound)	
		return	
		}	
	w.WriteHeader(http.StatusNoContent)	
}		