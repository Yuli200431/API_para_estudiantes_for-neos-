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

type cooperativaRepoMock struct {
	mock.Mock
}

func (m *cooperativaRepoMock) ListarCooperativas() []models.Cooperativa {
	args := m.Called()
	return args.Get(0).([]models.Cooperativa)
}

func (m *cooperativaRepoMock) BuscarCooperativaPorID(id uint) (models.Cooperativa, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Cooperativa), args.Bool(1)
}

func (m *cooperativaRepoMock) CrearCooperativa(c models.Cooperativa) models.Cooperativa {
	args := m.Called(c)
	return args.Get(0).(models.Cooperativa)
}

func (m *cooperativaRepoMock) ActualizarCooperativa(id uint, c models.Cooperativa) (models.Cooperativa, bool) {
	args := m.Called(id, c)
	return args.Get(0).(models.Cooperativa), args.Bool(1)
}

func (m *cooperativaRepoMock) BorrarCooperativa(id uint) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Red de seguridad: el mock DEBE cumplir el contrato de la interfaz.
var _ storage.CooperativaRepository = (*cooperativaRepoMock)(nil)

func TestCooperativaService_Crear(t *testing.T) {
	cooperativaValida := models.Cooperativa{
		Nombre:      "Coop Manta",
		Telefono:    "099123456",
		Descripcion: "Cooperativa de prueba",
	}

	casos := []struct {
		nombre        string
		entrada       models.Cooperativa
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "cooperativa valida",
			entrada:       cooperativaValida,
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre:        "nombre vacio -> error, no se persiste",
			entrada:       models.Cooperativa{Nombre: "", Telefono: "099", Descripcion: "desc"},
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
		{
			nombre: "telefono vacio -> error, no se persiste",
			entrada: func() models.Cooperativa {
				c := cooperativaValida
				c.Telefono = ""
				return c
			}(),
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(cooperativaRepoMock)
			if c.debePersistir {
				guardada := c.entrada
				guardada.ID = 1
				repo.On("CrearCooperativa", c.entrada).Return(guardada)
			}
			svc := service.NuevaCooperaticaService(repo)

			creada, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearCooperativa")
			} else {
				require.NoError(t, err)
				assert.Equal(t, uint(1), creada.ID)
				repo.AssertCalled(t, "CrearCooperativa", c.entrada)
			}
		})
	}
}

func TestCooperativaService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(cooperativaRepoMock)
	repo.On("BuscarCooperativaPorID", uint(999)).Return(models.Cooperativa{}, false)

	svc := service.NuevaCooperaticaService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
