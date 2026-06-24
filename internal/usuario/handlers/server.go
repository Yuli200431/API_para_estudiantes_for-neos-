package handlers

import (
	"for-neos-api/internal/usuario/service"
)

type Server struct {
	Auth *service.AuthService
}

func NewServer(a *service.AuthService) *Server {
	return &Server{
		Auth: a,
	}
}
