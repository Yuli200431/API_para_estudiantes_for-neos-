package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"log"

	"for-neos-api/internal/alimentacion/models"
	"github.com/go-chi/chi/v5"
	"encoding/json"

)

// ListarMenuDiarios maneja la solicitud GET /menudiarios
func (s *Server) ListarMenuDiarios(w http.ResponseWriter, _ *http.Request) {
	
menuDiarios := s.storage.ListarMenuDiarios()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(menuDiarios); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

// BuscarMenuDiarioPorID maneja la solicitud GET /menudiarios/{id}
func (s *Server) BuscarMenuDiarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return 
	}

	menuDiario, encontrado := s.storage.BuscarMenuDiarioPorID(id)
	if !encontrado {
		http.Error(w, "Alimentación no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(menuDiario); err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

//CrearMenuDiario maneja la solicitud POST/api/c /menudiarios	

func (s *Server) CrearMenuDiario(w http.ResponseWriter, r *http.Request) {
	var nuevo models.MenuDiario
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)	
		return
	}	
	if strings.TrimSpace(nuevo.Fecha) == "" {
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)	
		return
	}
	creado := s.storage.CrearMenuDiario(nuevo)
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusCreated)	
	if err := json.NewEncoder(w).Encode(creado); err != nil {	
		log.Printf("Error al codificar la respuesta: %v", err)	

	}	
}	
// ACTUALIZAR menuDiario maneja la solicitud PUT/api/v /menudiarios/{id}	
func (s *Server) ActualizarMenuDiario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))	
	if err != nil {	
		http.Error(w, "ID inválido", http.StatusBadRequest)	
		return	
	}	
	var datos models.MenuDiario	
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {	
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)	
		return	
	}
	if strings.TrimSpace(datos.Fecha) == "" {	
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)	
		return	
		}	
	actualizado, encontrado := s.storage.ActualizarMenuDiario(id, datos)	
	if !encontrado {	
		http.Error(w, "MenuDiario no encontrada", http.StatusNotFound)	
		return	
		}	
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusOK)	
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {	
		log.Printf("Error al codificar la respuesta: %v", err)	
	}	
}

// BorrarMenuDiario maneja la solicitud DELETE /menudiarios/{id}
func (s *Server) BorrarMenuDiario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))	
	if err != nil {	
		http.Error(w, "ID inválido", http.StatusBadRequest)	
		return	
		}	
	if !s.storage.BorrarMenuDiario(id) {	
		http.Error(w, "MenuDiario no encontrada", http.StatusNotFound)	
		return	
		}	
	w.WriteHeader(http.StatusNoContent)	
}
