package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/storage"
	"for-neos-api/internal/transporte/models"
	"for-neos-api/internal/transporte/service"
)

type rutaTransporteRepoMock struct {
	mock.Mock
}

func (m *rutaTransporteRepoMock) ListarRutas() []models.RutaTransporte {
	args := m.Called()
	return args.Get(0).([]models.RutaTransporte)
}

func (m *rutaTransporteRepoMock) BuscarRutaPorID(id uint) (models.RutaTransporte, bool) {
	args := m.Called(id)
	return args.Get(0).(models.RutaTransporte), args.Bool(1)
}

func (m *rutaTransporteRepoMock) CrearRuta(r models.RutaTransporte) models.RutaTransporte {
	args := m.Called(r)
	return args.Get(0).(models.RutaTransporte)
}

func (m *rutaTransporteRepoMock) ActualizarRuta(id uint, r models.RutaTransporte) (models.RutaTransporte, bool) {
	args := m.Called(id, r)
	return args.Get(0).(models.RutaTransporte), args.Bool(1)
}

func (m *rutaTransporteRepoMock) BorrarRuta(id uint) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.RutaTransporteRepository = (*rutaTransporteRepoMock)(nil)

func TestRutaTransporteService_Crear(t *testing.T) {
	rutaValida := models.RutaTransporte{
		NombreLinea:     "Línea 1",
		FrecuenciaAprox: "Cada 10 minutos",
		Precio:          0.40,
		DescripcionRuta: "Ruta principal",
		CooperativaID:   1,
		SectorOrigenID:  1,
		SectorDestinoID: 2,
		ParadaBusID:     1,
	}

	casos := []struct {
		nombre        string
		entrada       models.RutaTransporte
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "ruta valida",
			entrada:       rutaValida,
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre:        "nombre linea vacio -> error, no se persiste",
			entrada:       models.RutaTransporte{NombreLinea: ""},
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(rutaTransporteRepoMock)
			if c.debePersistir {
				guardada := c.entrada
				guardada.ID = 1
				repo.On("CrearRuta", c.entrada).Return(guardada)
			}
			svc := service.NuevaRutaService(repo)

			creada, err := svc.CrearRuta(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearRuta")
			} else {
				require.NoError(t, err)
				assert.Equal(t, uint(1), creada.ID)
				repo.AssertCalled(t, "CrearRuta", c.entrada)
			}
		})
	}
}

func TestRutaTransporteService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(rutaTransporteRepoMock)
	repo.On("BuscarRutaPorID", uint(999)).Return(models.RutaTransporte{}, false)

	svc := service.NuevaRutaService(repo)

	_, err := svc.ObtenerRutas(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}