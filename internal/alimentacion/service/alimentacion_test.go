package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/storage"
)

// alimentacionRepoMock es un doble de prueba de storage.AlimentacionRepository.
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

// Chequeo en tiempo de compilacion: el mock debe cumplir el contrato.
var _ storage.AlimentacionRepository = (*alimentacionRepoMock)(nil)

// TestAlimentacionService_Crear prueba la regla de negocio: nombre vacio se rechaza.
func TestAlimentacionService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Alimentacion
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "nombre vacio -> ErrNombreVacio",
			entrada:       models.Alimentacion{NombreLocal: "   ", Descripcion: "Comida casera"},
			errEsperado:   service.ErrNombreVacio,
			debePersistir: false,
		},
		{
			nombre:        "alimentacion valida -> sin error y se persiste",
			entrada:       models.Alimentacion{NombreLocal: "Comedor ULEAM", Descripcion: "Comida economica", PrecioPromedio: 2.50},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(alimentacionRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 42
				repo.On("CrearAlimentacion", c.entrada).Return(guardado)
			}
			svc := service.NuevaAlimentacionService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearAlimentacion")
			} else {
				require.NoError(t, err)
				require.Equal(t, 42, creado.ID)
				repo.AssertCalled(t, "CrearAlimentacion", c.entrada)
			}
		})
	}
}

// TestAlimentacionService_Obtener_NoEncontrado prueba que un ID inexistente
// se traduce en ErrNoEncontrado.
func TestAlimentacionService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(alimentacionRepoMock)
	repo.On("BuscarAlimentacionPorID", 999).Return(models.Alimentacion{}, false)
	svc := service.NuevaAlimentacionService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}