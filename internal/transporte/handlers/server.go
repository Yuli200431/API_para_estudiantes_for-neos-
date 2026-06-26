package handlers

import (
	"for-neos-api/internal/transporte/service"
)

type Server struct {
	Cooperativas *service.CooperativaService
	Paradas      *service.ParadaBusService
	Rutas        *service.RutaTransporteService
}

func NewServer(cooperativas *service.CooperativaService, paradas *service.ParadaBusService, rutas *service.RutaTransporteService,) *Server {
	return &Server{
		Cooperativas: cooperativas,
		Paradas:      paradas,
		Rutas:        rutas,
	}
}