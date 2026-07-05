package handlers

import (
	"for-neos-api/internal/alimentacion/service"
)

type Server struct {
	Alimentacion *service.AlimentacionService
	MenuDiario   *service.MenuDiarioService
	Plato        *service.PlatoService
	Resena       *service.ResenaService
}

type Deps struct {
	Alimentacion *service.AlimentacionService
	MenuDiario   *service.MenuDiarioService
	Plato        *service.PlatoService
	Resena       *service.ResenaService
}

func NewServer(d Deps) *Server {
	return &Server{
		Alimentacion: d.Alimentacion,
		MenuDiario:   d.MenuDiario,
		Plato:        d.Plato,
		Resena:       d.Resena,
	}
}
