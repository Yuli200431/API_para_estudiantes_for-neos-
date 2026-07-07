package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/transporte/models"
	"for-neos-api/internal/transporte/service"
)

// cooperativaRepoMock es un doble de prueba de storage.CooperativaRepository.
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

// TestCooperativaService_Crear comprueba la regla de negocio (validacionCooperativa)
// de forma aislada, sin tocar la base de datos real.
func TestCooperativaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Cooperativa
		debeFallar    bool
		debePersistir bool
	}{
		{
			nombre:        "nombre vacio -> error, no se persiste",
			entrada:       models.Cooperativa{Nombre: "   ", Telefono: "099"},
			debeFallar:    true,
			debePersistir: false,
		},
		{
			nombre:        "cooperativa valida -> sin error, se persiste",
			entrada:       models.Cooperativa{Nombre: "Coop Manta", Telefono: "099123"},
			debeFallar:    false,
			debePersistir: true,
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

			if c.debeFallar {
				require.Error(t, err)
				repo.AssertNotCalled(t, "CrearCooperativa")
			} else {
				require.NoError(t, err)
				require.Equal(t, uint(1), creada.ID)
				repo.AssertCalled(t, "CrearCooperativa", c.entrada)
			}
		})
	}
}