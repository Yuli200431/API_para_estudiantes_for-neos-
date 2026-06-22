package handlers

import (
	"for-neos-api/internal/usuario/service"
)

type Server struct {
	Auth *service.AuthService
}

func NewServer(auth *service.AuthService) *Server {
	return &Server{
		Auth: auth,
	}
}
