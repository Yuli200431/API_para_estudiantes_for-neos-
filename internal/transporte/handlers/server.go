package handlers

import (
	"for-neos-api/internal/transporte/service"
)

type Server struct {
	Cooperativas *service.CooperativaService
	Paradas      *service.ParadaBusService
	Rutas        *service.RutaTransporteService
}

type Deps struct {
	Cooperativas *service.CooperativaService
	Paradas      *service.ParadaBusService
	Rutas        *service.RutaTransporteService
}

func NewServer(d Deps) *Server {
	return &Server{
		Cooperativas: d.Cooperativas,
		Paradas:      d.Paradas,
		Rutas:        d.Rutas,
	}
}
