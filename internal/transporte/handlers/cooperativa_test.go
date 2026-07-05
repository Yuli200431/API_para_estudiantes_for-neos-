package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/middleware"
	"for-neos-api/internal/transporte/handlers"
	"for-neos-api/internal/transporte/models"
	"for-neos-api/internal/transporte/service"
	usuarioModels "for-neos-api/internal/usuario/models"
	usuarioService "for-neos-api/internal/usuario/service"
)

//Fake repo de cooperativas

type fakeCooperativaRepo struct {
	datos  []models.Cooperativa
	nextID uint
}

func nuevoFakeRepo() *fakeCooperativaRepo {
	return &fakeCooperativaRepo{datos: []models.Cooperativa{}, nextID: 1}
}

func (f *fakeCooperativaRepo) ListarCooperativas() []models.Cooperativa { return f.datos }

func (f *fakeCooperativaRepo) BuscarCooperativaPorID(id uint) (models.Cooperativa, bool) {
	for _, c := range f.datos {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cooperativa{}, false
}

func (f *fakeCooperativaRepo) CrearCooperativa(c models.Cooperativa) models.Cooperativa {
	c.ID = f.nextID
	f.nextID++
	f.datos = append(f.datos, c)
	return c
}

func (f *fakeCooperativaRepo) ActualizarCooperativa(id uint, c models.Cooperativa) (models.Cooperativa, bool) {
	for i, x := range f.datos {
		if x.ID == id {
			c.ID = id
			f.datos[i] = c
			return c, true
		}
	}
	return models.Cooperativa{}, false
}

func (f *fakeCooperativaRepo) BorrarCooperativa(id uint) bool {
	for i, x := range f.datos {
		if x.ID == id {
			f.datos = append(f.datos[:i], f.datos[i+1:]...)
			return true
		}
	}
	return false
}

//Fake repo de usuarios

type fakeUsuarioRepo struct {
	porEmail map[string]usuarioModels.Usuario
	nextID   int
}

func nuevoFakeUsuarioRepo() *fakeUsuarioRepo {
	return &fakeUsuarioRepo{porEmail: map[string]usuarioModels.Usuario{}, nextID: 1}
}

func (f *fakeUsuarioRepo) CrearUsuario(u usuarioModels.Usuario) (usuarioModels.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *fakeUsuarioRepo) BuscarUsuarioPorEmail(email string) (usuarioModels.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

//Armado del entorno de prueba

// construirEntorno crea el router con el middleware.Auth REAL (no simulado),
// usando un repo de cooperativas fake y un repo de usuarios fake.
// Devuelve el router y un token válido ya emitido por login real.
func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	repo := nuevoFakeRepo()
	svc := service.NuevaCooperaticaService(repo)
	srv := handlers.NewServer(
		handlers.Deps{
			Cooperativas: svc,
		},
	)

	usuarios := nuevoFakeUsuarioRepo()
	authSvc := usuarioService.NuevoAuthService(usuarios)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc)) // middleware real, no simulado
		r.Get("/api/v1/cooperativa", srv.ListarCooperativas)
		r.Post("/api/v1/cooperativa", srv.AgregarCooperativa)
	})

	// Registramos un usuario y obtenemos su token con Login real.
	_, err := authSvc.Registrar("Docente", "docente@uleam.edu.ec", "secreta123")
	require.NoError(t, err)

	token, err := authSvc.Login("docente@uleam.edu.ec", "secreta123")
	require.NoError(t, err)

	return r, token
}

// TestAgregarCooperativa_ConToken_201: con token válido, POST responde 201.
func TestAgregarCooperativa_ConToken_201(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"nombre":"Coop Test","telefono":"099123","descripcion":"desc"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/cooperativa", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada models.Cooperativa
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))
	require.Equal(t, "Coop Test", creada.Nombre)
}

// TestListarCooperativas_SinToken_401: sin header Authorization,
// el middleware real corta antes de llegar al handler -> 401.
func TestListarCooperativas_SinToken_401(t *testing.T) {
	h, _ := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cooperativa", nil)
	// A propósito: NO se envía el header Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
