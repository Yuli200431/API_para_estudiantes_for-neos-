// Package handlers contiene los handlers HTTP de la API.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-chi/chi/v5"

	"for-neos-api/internal/vivienda/models"
	"for-neos-api/internal/vivienda/storage"
)

// Server agrupa las dependencias compartidas por los handlers.
// Recibe el storage por inyección de dependencias desde main.
type Server struct {
	Storage *storage.Memoria
}

// NewServer construye un Server listo para usar.
func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

// ListarViviendas atiende GET /api/v1/viviendas.
func (s *Server) ListarViviendas(w http.ResponseWriter, _ *http.Request) {
	viviendas := s.Storage.ListarViviendas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(viviendas); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerVivienda atiende GET /api/v1/viviendas/{id}.
func (s *Server) ObtenerVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	vivienda, encontrado := s.Storage.BuscarViviendaPorID(id)
	if !encontrado {
		http.Error(w, "Vivienda no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vivienda); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearVivienda atiende POST /api/v1/viviendas.
func (s *Server) CrearVivienda(w http.ResponseWriter, r *http.Request) {
	var nueva models.Vivienda
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Titulo) == "" {
		http.Error(w, "El campo Titulo es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.TipoVivienda) == "" {
		http.Error(w, "El campo Tipo de Vivienda es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.TipoVivienda != "Casa" && nueva.TipoVivienda != "Departamento" && nueva.TipoVivienda != "Cuarto" && nueva.TipoVivienda != "Departamento Compartido" {
		http.Error(w, "El campo Tipo de Vivienda debe ser 'Casa' o 'Departamento' o 'Departamento Compartido' o 'Cuarto'", http.StatusBadRequest)
		return
	}
	if nueva.Precio < 0 {
		http.Error(w, "El campo Precio no puede ser negativo", http.StatusBadRequest)
		return
	}
	if nueva.Garantia == nil {
		http.Error(w, "El campo Garantia es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.PrecioGarantia < 0 {
		http.Error(w, "El campo Precio de Garantia no puede ser negativo", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Direccion) == "" {
		http.Error(w, "El campo Direccion es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Luz == nil {
		http.Error(w, "El campo Luz es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Agua == nil {
		http.Error(w, "El campo Agua es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Amueblado == nil {
		http.Error(w, "El campo Amueblado es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Internet == nil {
		http.Error(w, "El campo Internet es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.BañoPrivado == nil {
		http.Error(w, "El campo Baño Privado es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.NumeroHabitaciones < 0 {
		http.Error(w, "El campo Número de Habitaciones no puede ser negativo", http.StatusBadRequest)
		return
	}
	if nueva.Mascotas == nil {
		http.Error(w, "El campo Mascotas es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.GeneroPreferido) == "" {
		http.Error(w, "El campo Genero Preferido es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.GeneroPreferido != "Mixto" && nueva.GeneroPreferido != "Masculino" && nueva.GeneroPreferido != "Femenino" {
		http.Error(w, "El campo Genero Preferido debe ser una de las siguientes opciones:'Mixto', 'Masculino' o 'Femenino'", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.ReglasConvivencia) == "" {
		http.Error(w, "El campo Reglas de Convivencia es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Telefono) == "" {
		http.Error(w, "El campo Teléfono es obligatorio", http.StatusBadRequest)
		return
	}
	for _, r := range nueva.Telefono {
		if !unicode.IsDigit(r) && r != '+' && r != '-' && r != ' ' {
			http.Error(w, "El campo Teléfono no tiene un formato válido", http.StatusBadRequest)
			return
		}
	}
	if strings.TrimSpace(nueva.Email) == "" {
		http.Error(w, "El campo Email es obligatorio", http.StatusBadRequest)
		return
	}
	if _, err := mail.ParseAddress(nueva.Email); err != nil {
		http.Error(w, "El campo Email no tiene un formato válido", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Estado) == "" {
		http.Error(w, "El campo Estado es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Estado != "Disponible" && nueva.Estado != "Ocupado" && nueva.Estado != "Reservado" {
		http.Error(w, "El campo Estado debe ser 'Disponible', 'Ocupado' o 'Reservado'", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Comentario) == "" {
		http.Error(w, "El campo Comentario es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.SectorID <= 0 {
		http.Error(w, "El campo Sector_Id debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if len(nueva.Fotos) == 0 {
		http.Error(w, "El campo Fotos no puede estar vacío", http.StatusBadRequest)
		return
	}
	if nueva.ProveedorID <= 0 {
		http.Error(w, "El campo Proveedor_Id debe ser un número entero positivo", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearVivienda(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// ActualizarVivienda atiende PUT /api/v1/viviendas/{id}.
func (s *Server) ActualizarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Vivienda
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Titulo) == "" {
		http.Error(w, "El campo Titulo es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.TipoVivienda) == "" {
		http.Error(w, "El campo Tipo de Vivienda es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.TipoVivienda != "Casa" && datos.TipoVivienda != "Departamento" && datos.TipoVivienda != "Cuarto" && datos.TipoVivienda != "Departamento Compartido" {
		http.Error(w, "El campo Tipo de Vivienda debe ser 'Casa' o 'Departamento' o 'Departamento Compartido' o 'Cuarto'", http.StatusBadRequest)
		return
	}
	if datos.Precio < 0 {
		http.Error(w, "El Precio no puede ser negativo", http.StatusBadRequest)
		return
	}
	if datos.Garantia == nil {
		http.Error(w, "El campo Garantia es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.PrecioGarantia < 0 {
		http.Error(w, "El campo Precio de Garantia no puede ser negativo", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Direccion) == "" {
		http.Error(w, "El campo Direccion es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Luz == nil {
		http.Error(w, "El campo Luz es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Agua == nil {
		http.Error(w, "El campo Agua es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Amueblado == nil {
		http.Error(w, "El campo Amueblado es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Internet == nil {
		http.Error(w, "El campo Internet es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.BañoPrivado == nil {
		http.Error(w, "El campo Baño Privado es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.NumeroHabitaciones < 0 {
		http.Error(w, "El campo NÚmero de Habitaciones no puede ser negativo", http.StatusBadRequest)
		return
	}
	if datos.Mascotas == nil {
		http.Error(w, "El campo Mascotas es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.GeneroPreferido) == "" {
		http.Error(w, "El campo Género Preferido es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.GeneroPreferido != "Mixto" && datos.GeneroPreferido != "Masculino" && datos.GeneroPreferido != "Femenino" {
		http.Error(w, "El campo Género Preferido debe ser una de las siguientes opciones:'Mixto', 'Masculino' o 'Femenino'", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.ReglasConvivencia) == "" {
		http.Error(w, "El campo Reglas de Convivencia es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Telefono) == "" {
		http.Error(w, "El campo Teléfono es obligatorio", http.StatusBadRequest)
		return
	}
	for _, r := range datos.Telefono {
		if !unicode.IsDigit(r) && r != '+' && r != '-' && r != ' ' {
			http.Error(w, "El campo Télefono no tiene un formato válido", http.StatusBadRequest)
			return
		}
	}
	if strings.TrimSpace(datos.Email) == "" {
		http.Error(w, "el campo email es obligatorio", http.StatusBadRequest)
		return
	}
	if _, err := mail.ParseAddress(datos.Email); err != nil {
		http.Error(w, "El campo Email no tiene un formato válido", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Estado) == "" {
		http.Error(w, "El campo Estado es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Estado != "Disponible" && datos.Estado != "Ocupado" && datos.Estado != "Reservado" {
		http.Error(w, "El campo Estado debe ser 'Disponible', 'Ocupado' o 'Reservado'", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Comentario) == "" {
		http.Error(w, "El campo Comentario es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.SectorID <= 0 {
		http.Error(w, "El campo Sector Id debe ser un número entero positivo", http.StatusBadRequest)
		return
	}
	if len(datos.Fotos) == 0 {
		http.Error(w, "El campo Fotos no puede estar vacío", http.StatusBadRequest)
		return
	}
	if datos.ProveedorID <= 0 {
		http.Error(w, "El campo Proveedor Id debe ser un número entero positivo", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.ActualizarVivienda(id, datos)
	if !encontrado {
		http.Error(w, "La Vivienda no fue encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("Error codificando JSON: %v", err)
	}
}

// BorrarVivienda atiende DELETE /api/v1/viviendas/{id}.
func (s *Server) BorrarVivienda(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "El Id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarVivienda(id) {
		http.Error(w, "La Vivienda no fue encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
