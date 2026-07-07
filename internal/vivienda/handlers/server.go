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

type Deps struct {
	Viviendas *service.ViviendaService
	Fotos     *service.FotoService
	Aplicar   *service.AplicarViviendaService
	Sectores  *service.SectorService
}

func NewServer(d Deps) *Server {

	return &Server{
		Viviendas: d.Viviendas,
		Fotos:     d.Fotos,
		Aplicar:   d.Aplicar,
		Sectores:  d.Sectores,
	}
}
