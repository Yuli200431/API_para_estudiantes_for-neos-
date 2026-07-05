package handlers

import (
	"for-neos-api/internal/usuario/service"
)

type Server struct {
	Auth *service.AuthService
}

type Deps struct {
	Auth *service.AuthService
}

func NewServer(d Deps) *Server {
	return &Server{
		Auth: d.Auth,
	}
}
