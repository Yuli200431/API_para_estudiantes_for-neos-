package middleware

import (
	"net/http"
)

// RequiereRol verifica que el usuario autenticado tenga el rol esperado.
// Debe usarse DESPUÉS del middleware Auth, que ya puso el rol en el contexto.
func RequiereRol(rolEsperado string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, ok := r.Context().Value(claveRol).(string)
			if !ok || rol != rolEsperado {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"error": "No tienes permiso para acceder a este recurso"}`))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}