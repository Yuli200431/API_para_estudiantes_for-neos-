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


// ===============================
// FAKE REPOSITORY ALIMENTACION
// ===============================

//sirve para crear un repositorio de alimentacion//
type alimentacionRepoFake struct {
	mu     sync.Mutex
	datos  map[int]models.Alimentacion
	nextID int
}

//crea un repositorio de alimentacion vacio//
func nuevoAlimentacionRepoFake() *alimentacionRepoFake {
	return &alimentacionRepoFake{
		datos:  make(map[int]models.Alimentacion),
		nextID: 1,
	}
}

//ListarAlimentaciones devuelve la lista de alimentaciones//
func (f *alimentacionRepoFake) ListarAlimentaciones() []models.Alimentacion {

	f.mu.Lock()
	defer f.mu.Unlock()

	lista := []models.Alimentacion{}

	for _, a := range f.datos {
		lista = append(lista, a)
	}

	return lista
}

//BuscarAlimentacionPorID devuelve la alimentacion con el ID dado//
func (f *alimentacionRepoFake) BuscarAlimentacionPorID(id int) (models.Alimentacion, bool) {

	f.mu.Lock()
	defer f.mu.Unlock()

	a, ok := f.datos[id]

	return a, ok
}

//CrearAlimentacion agrega una alimentacion nueva y devuelve la alimentacion con ID asignado//
func (f *alimentacionRepoFake) CrearAlimentacion(a models.Alimentacion) models.Alimentacion {

	f.mu.Lock()
	defer f.mu.Unlock()

	a.ID = f.nextID
	f.nextID++

	f.datos[a.ID] = a

	return a
}

//ActualizarAlimentacion reemplaza la alimentacion con el ID dado//
func (f *alimentacionRepoFake) ActualizarAlimentacion(id int, datos models.Alimentacion) (models.Alimentacion, bool) {

	f.mu.Lock()
	defer f.mu.Unlock()

	_, existe := f.datos[id]

	if !existe {
		return models.Alimentacion{}, false
	}

	datos.ID = id
	f.datos[id] = datos

	return datos, true
}

//BorrarAlimentacion elimina la alimentacion con el ID dado//
func (f *alimentacionRepoFake) BorrarAlimentacion(id int) bool {

	f.mu.Lock()
	defer f.mu.Unlock()

	_, existe := f.datos[id]

	if !existe {
		return false
	}

	delete(f.datos, id)

	return true
}



// ===============================
// FAKE USUARIO
// ===============================
//sirve para crear un repositorio de usuario//
type usuarioRepoFake struct {
	porEmail map[string]usuarioModels.Usuario
	nextID int
}

//crea un repositorio de usuario vacio//
func nuevoUsuarioRepoFake() *usuarioRepoFake {

	return &usuarioRepoFake{
		porEmail: make(map[string]usuarioModels.Usuario),
		nextID:1,
	}
}

//CrearUsuario agrega un usuario nuevo y devuelve el usuario con ID asignado//
func (f *usuarioRepoFake) CrearUsuario(u usuarioModels.Usuario)(usuarioModels.Usuario,error){

	u.ID=f.nextID
	f.nextID++

	f.porEmail[u.Email]=u

	return u,nil
}

//BuscarUsuarioPorEmail devuelve el usuario con el email dado//
func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string)(usuarioModels.Usuario,bool){

	u,ok:=f.porEmail[email]

	return u,ok
}



// ===============================
// ENTORNO
// ===============================

func construirEntorno(t *testing.T)(http.Handler,string){

	t.Helper()


	repo:=nuevoAlimentacionRepoFake()

	srvService:=service.NuevaAlimentacionService(repo)


	usuarios:=nuevoUsuarioRepoFake()

	auth:=usuarioService.NuevoAuthService(usuarios)


	srv:=handlers.NewServer(
		handlers.Deps{
			Alimentacion:srvService,
		},
	)


	r:=chi.NewRouter()


	r.Route("/api/v1",func(r chi.Router){

		r.Group(func(r chi.Router){

			r.Use(middleware.Auth(auth))


			r.Get("/alimentaciones",srv.ListarAlimentaciones)

			r.Post("/alimentaciones",srv.CrearAlimentacion)

			r.Get("/alimentaciones/{id}",srv.BuscarAlimentacionesPorID)

			r.Put("/alimentaciones/{id}",srv.ActualizarAlimentacion)

			r.Delete("/alimentaciones/{id}",srv.BorrarAlimentacion)

		})

	})


	token:=registrarYObtenerToken(t,auth)


	return r,token
}



func registrarYObtenerToken(t *testing.T,auth *usuarioService.AuthService)string{

	t.Helper()


	_,err:=auth.Registrar(
		"Docente",
		"docente@uleam.edu.ec",
		"secreta123",
	)

	require.NoError(t,err)


	token,err:=auth.Login(
		"docente@uleam.edu.ec",
		"secreta123",
	)


	require.NoError(t,err)

	return token
}



// ===============================
// BODY VALIDO
// ===============================

func alimentacionBody()string{

	return `{
	"nombre_local":"Comedor ULEAM",
	"descripcion":"Comida economica",
	"ubicacion":"Manta",
	"direccion":"Av Universitaria",
	"horario_apertura":"08:00",
	"horario_cierre":"14:00",
	"telefono":"0999999999",
	"tipo_comida":"Ecuatoriana",
	"precio_promedio":2.5,
	"provider_id":1
	}`
}



// ===============================
// TEST CREAR
// ===============================

func TestCrearAlimentacion_Exitoso(t *testing.T){

	h,token:=construirEntorno(t)


	req:=httptest.NewRequest(
		http.MethodPost,
		"/api/v1/alimentaciones",
		strings.NewReader(alimentacionBody()),
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	require.Equal(
		t,
		http.StatusCreated,
		rec.Code,
	)


	var creado models.Alimentacion

	require.NoError(
		t,
		json.NewDecoder(rec.Body).Decode(&creado),
	)


	assert.NotZero(t,creado.ID)

}




func TestCrearAlimentacion_Invalido(t *testing.T){

	h,token:=construirEntorno(t)


	body:=`{
	"nombre_local":"",
	"precio_promedio":2
	}`


	req:=httptest.NewRequest(
		http.MethodPost,
		"/api/v1/alimentaciones",
		strings.NewReader(body),
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	assert.Equal(
		t,
		http.StatusBadRequest,
		rec.Code,
	)

}



// ===============================
// LISTAR
// ===============================

func TestListarAlimentaciones_Exitoso(t *testing.T){

	h,token:=construirEntorno(t)


	req:=httptest.NewRequest(
		http.MethodGet,
		"/api/v1/alimentaciones",
		nil,
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

}



// ===============================
// OBTENER
// ===============================

func TestObtenerAlimentacion_NoEncontrado(t *testing.T){

	h,token:=construirEntorno(t)


	req:=httptest.NewRequest(
		http.MethodGet,
		"/api/v1/alimentaciones/999",
		nil,
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	assert.Equal(
		t,
		http.StatusNotFound,
		rec.Code,
	)

}



// ===============================
// ACTUALIZAR
// ===============================

func TestActualizarAlimentacion_Exitoso(t *testing.T){

	h,token:=construirEntorno(t)


	crear:=httptest.NewRequest(
		http.MethodPost,
		"/api/v1/alimentaciones",
		strings.NewReader(alimentacionBody()),
	)

	crear.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	recCrear:=httptest.NewRecorder()

	h.ServeHTTP(recCrear,crear)



	actualizar:=httptest.NewRequest(
		http.MethodPut,
		"/api/v1/alimentaciones/1",
		strings.NewReader(alimentacionBody()),
	)


	actualizar.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,actualizar)


	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

}



// ===============================
// BORRAR
// ===============================

func TestBorrarAlimentacion_Exitoso(t *testing.T){

	h,token:=construirEntorno(t)


	crear:=httptest.NewRequest(
		http.MethodPost,
		"/api/v1/alimentaciones",
		strings.NewReader(alimentacionBody()),
	)


	crear.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	recCrear:=httptest.NewRecorder()

	h.ServeHTTP(recCrear,crear)



	req:=httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/alimentaciones/1",
		nil,
	)


	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	assert.Equal(
		t,
		http.StatusNoContent,
		rec.Code,
	)

}



// ===============================
// SIN TOKEN
// ===============================

func TestRutaProtegida_SinToken(t *testing.T){

	h,_:=construirEntorno(t)


	req:=httptest.NewRequest(
		http.MethodGet,
		"/api/v1/alimentaciones",
		nil,
	)


	rec:=httptest.NewRecorder()


	h.ServeHTTP(rec,req)


	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)

}