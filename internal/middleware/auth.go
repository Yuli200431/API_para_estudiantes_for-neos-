package middleware

import (
	"context"
	"net/http"
	"strings"

	service "for-neos-api/internal/usuario/service"
)

type claveContexto string

const claveUsuarioID claveContexto = "usuarioID"
const claveRol claveContexto = "rol"

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				responderNoAutorizado(w)
				return
			}
			usuarioID, rol, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}

			ctx := context.WithValue(r.Context(), claveUsuarioID, usuarioID)
			ctx = context.WithValue(ctx, claveRol, rol)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error": "Token de autenticación requerido"}`))
}
