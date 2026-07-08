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

type paradaBusRepoMock struct {
	mock.Mock
}

func (m *paradaBusRepoMock) ListarParadas() []models.ParadaBus {
	args := m.Called()
	return args.Get(0).([]models.ParadaBus)
}

func (m *paradaBusRepoMock) BuscarParadaPorID(id uint) (models.ParadaBus, bool) {
	args := m.Called(id)
	return args.Get(0).(models.ParadaBus), args.Bool(1)
}

func (m *paradaBusRepoMock) CrearParada(p models.ParadaBus) models.ParadaBus {
	args := m.Called(p)
	return args.Get(0).(models.ParadaBus)
}

func (m *paradaBusRepoMock) ActualizarParada(id uint, p models.ParadaBus) (models.ParadaBus, bool) {
	args := m.Called(id, p)
	return args.Get(0).(models.ParadaBus), args.Bool(1)
}

func (m *paradaBusRepoMock) BorrarParada(id uint) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.ParadaBusRepository = (*paradaBusRepoMock)(nil)

func TestParadaBusService_Crear(t *testing.T) {
	paradaValida := models.ParadaBus{
		NombreParada: "Terminal Terrestre",
		Direccion:    "Av. 4 de Noviembre",
		Descripcion:  "Parada principal",
	}

	casos := []struct {
		nombre        string
		entrada       models.ParadaBus
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "parada valida",
			entrada:       paradaValida,
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre:        "nombre vacio -> error, no se persiste",
			entrada:       models.ParadaBus{NombreParada: "", Direccion: "Av. 4", Descripcion: "desc"},
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
		{
			nombre:        "direccion vacia -> error, no se persiste",
			entrada:       models.ParadaBus{NombreParada: "Terminal", Direccion: "", Descripcion: "desc"},
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(paradaBusRepoMock)
			if c.debePersistir {
				guardada := c.entrada
				guardada.ID = 1
				repo.On("CrearParada", c.entrada).Return(guardada)
			}
			svc := service.NuevaParadaBusService(repo)

			creada, err := svc.CrearParada(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearParada")
			} else {
				require.NoError(t, err)
				assert.Equal(t, uint(1), creada.ID)
				repo.AssertCalled(t, "CrearParada", c.entrada)
			}
		})
	}
}

func TestParadaBusService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(paradaBusRepoMock)
	repo.On("BuscarParadaPorID", uint(999)).Return(models.ParadaBus{}, false)

	svc := service.NuevaParadaBusService(repo)

	_, err := svc.ObtenerParadas(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}