package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/handlers"
	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/middleware"
	usuarioModels "for-neos-api/internal/usuario/models"
	usuarioService "for-neos-api/internal/usuario/service"
)

// alimentacionRepoFake: repositorio de alimentacion en memoria solo para
// estos tests de handler (no toca la base de datos real).
type alimentacionRepoFake struct {
	mu     sync.Mutex
	datos  map[int]models.Alimentacion
	nextID int
}

func nuevoAlimentacionRepoFake() *alimentacionRepoFake {
	return &alimentacionRepoFake{datos: map[int]models.Alimentacion{}, nextID: 1}
}

func (f *alimentacionRepoFake) ListarAlimentaciones() []models.Alimentacion {
	f.mu.Lock()
	defer f.mu.Unlock()
	lista := make([]models.Alimentacion, 0, len(f.datos))
	for _, a := range f.datos {
		lista = append(lista, a)
	}
	return lista
}

func (f *alimentacionRepoFake) BuscarAlimentacionPorID(id int) (models.Alimentacion, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	a, ok := f.datos[id]
	return a, ok
}

func (f *alimentacionRepoFake) CrearAlimentacion(a models.Alimentacion) models.Alimentacion {
	f.mu.Lock()
	defer f.mu.Unlock()
	a.ID = f.nextID
	f.nextID++
	f.datos[a.ID] = a
	return a
}

func (f *alimentacionRepoFake) ActualizarAlimentacion(id int, datos models.Alimentacion) (models.Alimentacion, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.datos[id]; !ok {
		return models.Alimentacion{}, false
	}
	datos.ID = id
	f.datos[id] = datos
	return datos, true
}

func (f *alimentacionRepoFake) BorrarAlimentacion(id int) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.datos[id]; !ok {
		return false
	}
	delete(f.datos, id)
	return true
}

// usuarioRepoFake: repositorio de usuarios en memoria para login/register.
type usuarioRepoFake struct {
	porEmail map[string]usuarioModels.Usuario
	nextID   int
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{porEmail: map[string]usuarioModels.Usuario{}, nextID: 1}
}

func (f *usuarioRepoFake) CrearUsuario(u usuarioModels.Usuario) (usuarioModels.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (usuarioModels.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

// construirEntorno arma el router con el middleware de auth REAL pero con
// repos en memoria (fakes), igual al patron del proyecto de referencia.
func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	repoAlimentacion := nuevoAlimentacionRepoFake()
	alimentacionSvc := service.NuevaAlimentacionService(repoAlimentacion)

	usuarios := nuevoUsuarioRepoFake()
	authSvc := usuarioService.NuevoAuthService(usuarios)

	srv := handlers.NewServer(alimentacionSvc, nil, nil, nil)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/alimentaciones", srv.ListarAlimentaciones)
			r.Post("/alimentaciones", srv.CrearAlimentacion)
			r.Get("/alimentaciones/{id}", srv.BuscarAlimentacionesPorID)
		})
	})

	token := registrarYObtenerToken(t, authSvc)
	return r, token
}

// registrarYObtenerToken crea un usuario y genera un token valido directo
// con el service (no por HTTP) para simplificar el setup del test.
func registrarYObtenerToken(t *testing.T, authSvc *usuarioService.AuthService) string {
	t.Helper()
	_, err := authSvc.Registrar("Docente", "docente@uleam.edu.ec", "secreta123")
	require.NoError(t, err)

	token, err := authSvc.Login("docente@uleam.edu.ec", "secreta123")
	require.NoError(t, err)
	require.NotEmpty(t, token)
	return token
}

// TestCrearAlimentacion_Exitoso: POST con token y cuerpo valido -> 201 Created.
func TestCrearAlimentacion_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"nombre_local":"Comedor ULEAM","descripcion":"Comida economica","precio_promedio":2.5}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/alimentaciones", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	var creado models.Alimentacion
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	assert.NotZero(t, creado.ID)
	assert.Equal(t, "Comedor ULEAM", creado.NombreLocal)
}

// TestObtenerAlimentacion_NoEncontrado: id inexistente -> 404 Not Found.
func TestObtenerAlimentacion_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/alimentaciones/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization -> 401 Unauthorized.
func TestRutaProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)
	body := `{"nombre_local":"Comedor ULEAM","precio_promedio":2.5}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/alimentaciones", strings.NewReader(body))
	// A proposito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}