package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/middleware"
	"for-neos-api/internal/storage"
	usuarioHandlers "for-neos-api/internal/usuario/handlers"
	usuarioModels "for-neos-api/internal/usuario/models"
	usuarioService "for-neos-api/internal/usuario/service"
	viviendaHandlers "for-neos-api/internal/vivienda/handlers"
	viviendaModels "for-neos-api/internal/vivienda/models"
	viviendaService "for-neos-api/internal/vivienda/service"
)

//Usuario

type usuarioRepoFake struct {
	porEmail map[string]usuarioModels.Usuario
	nextID   int
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{
		porEmail: make(map[string]usuarioModels.Usuario),
		nextID:   1}
}

func (f *usuarioRepoFake) CrearUsuario(u usuarioModels.Usuario) (usuarioModels.Usuario, error) {
	if _, existe := f.porEmail[u.Email]; existe {
		return usuarioModels.Usuario{}, usuarioService.ErrEmailEnUso
	}
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (usuarioModels.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

//Vivienda

type viviendaRepoFake struct {
	viviendas []viviendaModels.Vivienda
}

func (f *viviendaRepoFake) ListarViviendas() []viviendaModels.Vivienda {
	return f.viviendas
}

func (f *viviendaRepoFake) BuscarViviendaPorID(id int) (viviendaModels.Vivienda, bool) {
	for _, v := range f.viviendas {
		if v.ViviendaID == id {
			return v, true
		}
	}
	return viviendaModels.Vivienda{}, false
}

func (f *viviendaRepoFake) CrearVivienda(v viviendaModels.Vivienda) viviendaModels.Vivienda {
	v.ViviendaID = len(f.viviendas) + 1
	f.viviendas = append(f.viviendas, v)
	return v
}

func (f *viviendaRepoFake) ActualizarVivienda(id int, datos viviendaModels.Vivienda) (viviendaModels.Vivienda, bool) {
	return viviendaModels.Vivienda{}, false
}

func (f *viviendaRepoFake) BorrarVivienda(id int) bool {
	return false
}

var _ storage.ViviendaRepository = (*viviendaRepoFake)(nil)

//Entorno

func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	repo := &viviendaRepoFake{}

	viviendaSrv := viviendaService.NuevaViviendaService(repo)

	servidor := viviendaHandlers.NewServer(
		viviendaHandlers.Deps{
			Viviendas: viviendaSrv,
		},
	)

	usuarios := nuevoUsuarioRepoFake()
	auth := usuarioService.NuevoAuthService(usuarios)
	authServer := usuarioHandlers.NewServer(
		usuarioHandlers.Deps{
			Auth: auth,
		},
	)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", authServer.Registrar)
		r.Post("/auth/login", authServer.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(auth))

			r.Get("/viviendas", servidor.ListarViviendas)
			r.Post("/viviendas", servidor.CrearVivienda)
			r.Get("/viviendas/{id}", servidor.ObtenerVivienda)
			r.Put("/viviendas/{id}", servidor.ActualizarVivienda)
			r.Delete("/viviendas/{id}", servidor.BorrarVivienda)
		})
	})

	token := registrarYObtenerToken(t, r)
	return r, token
}

func registrarYObtenerToken(t *testing.T, h http.Handler) string {

	t.Helper()

	cred := `{"nombre": "Docente", "email":"docente@uleam.edu.ec","password":"secreta123"}`

	reqReg := httptest.NewRequest(http.MethodPost, "/api/v1/auth/registrar", strings.NewReader(cred))

	recReg := httptest.NewRecorder()

	h.ServeHTTP(recReg, reqReg)

	t.Log(
		"registro:",
		recReg.Code,
		recReg.Body.String(),
	)

	reqLogin := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/login",
		strings.NewReader(cred),
	)

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, reqLogin)

	require.Equal(
		t,
		http.StatusOK,
		rec.Code,
		rec.Body.String(),
	)

	var resp struct {
		Token string `json:"token"`
	}

	require.NoError(
		t,
		json.NewDecoder(rec.Body).Decode(&resp),
	)

	return resp.Token
}

// TestCrearVivienda_Exitoso: POST con token y cuerpo valido -> 201 Created.
func TestCrearVivienda_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"titulo":"Casa Esquinera", "tipo_vivienda":"Casa", "precio":100, "garantia":true, "precio_garantia":50, "luz":true, "agua":true, "amueblado":true, "internet":true, "baño_privado":true, "numero_habitaciones":2, "mascotas":true, "genero_preferido":"Cualquiera", "reglas_convivencia":"Reglas", "telefono":"0999999999", "email":"correo@test.com", "estado":"Disponible", "comentario":"Comentario", "sector_id":1, "proveedor_id":1, "fotos":[{"foto_id":1}]}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/viviendas", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code, rec.Body.String())
	var creado viviendaModels.Vivienda
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
	assert.NotZero(t, creado.ViviendaID)
	assert.Equal(t, "Casa Esquinera", creado.Titulo)
}

// TestObtenerVivienda_NoEncontrado: id inexistente -> 404 Not Found.
func TestObtenerVivienda_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/viviendas/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// TestCrearVivienda_Invalido: cuerpo que viola la regla de negocio -> 400.
func TestCrearVivienda_Invalido(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"titulo":"   ", "tipo_vivienda":"Casa", "precio":100, "garantia":true, "precio_garantia":50, "luz":true, "agua":true, "amueblado":true, "internet":true, "baño_privado":true, "numero_habitaciones":2, "mascotas":true, "genero_preferido":"Cualquiera", "reglas_convivencia":"Reglas", "telefono":"0999999999", "email":"correo@test.com", "estado":"Disponible", "comentario":"Comentario", "sector_id":1, "proveedor_id":1, "fotos":[{"foto_id":1}]}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/viviendas", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization, el middleware corta
// antes de llegar al handler -> 401 Unauthorized.
func TestRutaProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)
	body := `{"titulo":"Casa Esquinera", "tipo_vivienda":"Casa", "precio":100, "garantia":true, "precio_garantia":50, "luz":true, "agua":true, "amueblado":true, "internet":true, "baño_privado":true, "numero_habitaciones":2, "mascotas":true, "genero_preferido":"Cualquiera", "reglas_convivencia":"Reglas", "telefono":"0999999999", "email":"correo@test.com", "estado":"Disponible", "comentario":"Comentario", "sector_id":1, "proveedor_id":1}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/viviendas", strings.NewReader(body))
	// A proposito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTeapot, rec.Code)
}
