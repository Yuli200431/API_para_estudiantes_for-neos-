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

func NewServer(
	alimentacion *service.AlimentacionService,
	menuDiario *service.MenuDiarioService,
	plato *service.PlatoService,
	resena *service.ResenaService,
) *Server {
	return &Server{
		Alimentacion: alimentacion,
		MenuDiario:   menuDiario,
		Plato:        plato,
		Resena:       resena,
	}
}