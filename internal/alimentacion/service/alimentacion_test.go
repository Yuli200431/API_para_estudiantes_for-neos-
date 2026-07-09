package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/storage"
)

// --- MOCK DEL REPOSITORIO ---
type alimentacionRepoMock struct {
	mock.Mock
}

func (m *alimentacionRepoMock) ListarAlimentaciones() []models.Alimentacion {
	args := m.Called()
	return args.Get(0).([]models.Alimentacion)
}

func (m *alimentacionRepoMock) BuscarAlimentacionPorID(id int) (models.Alimentacion, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Alimentacion), args.Bool(1)
}

func (m *alimentacionRepoMock) CrearAlimentacion(a models.Alimentacion) models.Alimentacion {
	args := m.Called(a)
	return args.Get(0).(models.Alimentacion)
}

func (m *alimentacionRepoMock) ActualizarAlimentacion(id int, a models.Alimentacion) (models.Alimentacion, bool) {
	args := m.Called(id, a)
	return args.Get(0).(models.Alimentacion), args.Bool(1)
}

func (m *alimentacionRepoMock) BorrarAlimentacion(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.AlimentacionRepository = (*alimentacionRepoMock)(nil)

// --- DATOS DE PRUEBA ---
func alimentacionValida() models.Alimentacion {
	return models.Alimentacion{
		NombreLocal:    "Comedor ULEAM",
		Descripcion:    "Comida economica",
		Ubicacion:      "Manta",
		Direccion:      "Av. Universitaria",
		Telefono:       "0999999999",
		TipoComida:     "Ecuatoriana",
		PrecioPromedio: 2.50,
		ProviderID:     1,
	}
}

// ==========================================================
// 1. CREAR
// ==========================================================
func TestAlimentacionService_Crear(t *testing.T) {

	repo := new(alimentacionRepoMock)

	entrada := alimentacionValida()

	resultado := entrada
	resultado.ID = 42

	repo.On("CrearAlimentacion", entrada).
		Return(resultado)

	svc := service.NuevaAlimentacionService(repo)

	creado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, 42, creado.ID)

	repo.AssertExpectations(t)
}


// ==========================================================
// 2. OBTENER
// ==========================================================
func TestAlimentacionService_Obtener(t *testing.T) {

	t.Run("obtener exitoso", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		esperado := alimentacionValida()
		esperado.ID = 10

		repo.On("BuscarAlimentacionPorID", 10).
			Return(esperado, true)

		svc := service.NuevaAlimentacionService(repo)

		resultado, err := svc.Obtener(10)

		require.NoError(t, err)
		require.Equal(t, esperado, resultado)

		repo.AssertExpectations(t)
	})


	t.Run("no encontrado", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		repo.On("BuscarAlimentacionPorID", 999).
			Return(models.Alimentacion{}, false)

		svc := service.NuevaAlimentacionService(repo)

		_, err := svc.Obtener(999)

		require.ErrorIs(t, err, service.ErrNoEncontrado)

		repo.AssertExpectations(t)
	})
}


// ==========================================================
// 3. LISTAR
// ==========================================================
func TestAlimentacionService_Listar(t *testing.T) {

	repo := new(alimentacionRepoMock)

	lista := []models.Alimentacion{
		{
			ID:          1,
			NombreLocal: "Comedor 1",
		},
		{
			ID:          2,
			NombreLocal: "Comedor 2",
		},
	}


	repo.On("ListarAlimentaciones").
		Return(lista)


	svc := service.NuevaAlimentacionService(repo)

	resultado := svc.Listar()


	require.Len(t, resultado, 2)
	require.Equal(t, "Comedor 1", resultado[0].NombreLocal)

	repo.AssertExpectations(t)
}


// ==========================================================
// 4. ACTUALIZAR
// ==========================================================
func TestAlimentacionService_Actualizar(t *testing.T) {


	t.Run("actualizacion exitosa", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		a := alimentacionValida()

		repo.On("ActualizarAlimentacion", 1, a).
			Return(a, true)


		svc := service.NuevaAlimentacionService(repo)


		resultado, err := svc.Actualizar(1, a)


		require.NoError(t, err)
		require.Equal(t, a, resultado)

		repo.AssertExpectations(t)

	})


	t.Run("error por validacion", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		a := alimentacionValida()
		a.NombreLocal = ""


		svc := service.NuevaAlimentacionService(repo)


		_, err := svc.Actualizar(1, a)


		require.ErrorIs(t, err, service.ErrNombreVacio)

		repo.AssertNotCalled(t, "ActualizarAlimentacion")

	})


	t.Run("id no encontrado", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		a := alimentacionValida()


		repo.On("ActualizarAlimentacion", 555, a).
			Return(models.Alimentacion{}, false)


		svc := service.NuevaAlimentacionService(repo)


		_, err := svc.Actualizar(555, a)


		require.ErrorIs(t, err, service.ErrNoEncontrado)

		repo.AssertExpectations(t)

	})
}


// ==========================================================
// 5. BORRAR
// ==========================================================
func TestAlimentacionService_Borrar(t *testing.T) {


	t.Run("borrado exitoso", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		repo.On("BorrarAlimentacion", 1).
			Return(true)


		svc := service.NuevaAlimentacionService(repo)


		err := svc.Borrar(1)


		require.NoError(t, err)

		repo.AssertExpectations(t)

	})


	t.Run("no encontrado", func(t *testing.T) {

		repo := new(alimentacionRepoMock)

		repo.On("BorrarAlimentacion", 999).
			Return(false)


		svc := service.NuevaAlimentacionService(repo)


		err := svc.Borrar(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)

		repo.AssertExpectations(t)

	})
}