package handlers

import (
	"for-neos-api/internal/vivienda/service"
)

type Server struct {
	Viviendas *service.ViviendaService
	Fotos     *service.FotoService
	Aplicar   *service.AplicarViviendaService
	Sectores  *service.SectorService
}

func NewServer(viviendas *service.ViviendaService, fotos *service.FotoService, aplicar *service.AplicarViviendaService, sectores *service.SectorService) *Server {
	return &Server{
		Viviendas: viviendas,
		Fotos:     fotos,
		Aplicar:   aplicar,
		Sectores:  sectores,
	}
}
