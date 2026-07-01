package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"for-neos-api/internal/vivienda/service"
)

type mockSectorRepo struct {
	mock.Mock
}

func (m *mockSectorRepo) ListarSectores() []models.Sector {
	args := m.Called()
	return args.Get(0).([]models.Sector)
}

func (m *mockSectorRepo) BuscarSectorPorID(id int) (models.Sector, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Sector), args.Bool(1)
}

func (m *mockSectorRepo) CrearSector(e models.Sector) models.Sector {
	args := m.Called(e)
	return args.Get(0).(models.Sector)
}

func (m *mockSectorRepo) ActualizarSector(id int, datos models.Sector) (models.Sector, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Sector), args.Bool(1)
}

func (m *mockSectorRepo) BorrarSector(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.SectorRepository = (*mockSectorRepo)(nil)

// Primer Test
func TestSectorService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Sector
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "nombre vacio",
			entrada:       models.Sector{Nombre: "   "},
			errEsperado:   service.ErrNombreVacio,
			debePersistir: false,
		},
		{
			nombre:        "sector valido",
			entrada:       models.Sector{Nombre: "nombre"},
			errEsperado:   nil,
			debePersistir: true,
		},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(mockSectorRepo)
			if c.debePersistir {
				guardado := c.entrada
				guardado.SectorID = 42
				repo.On("CrearSector", c.entrada).Return(guardado)
			}
			svc := service.NuevaSectorService(repo)
			creado, err := svc.Crear(c.entrada)
			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearSector")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.SectorID, "el service debe devolver el producto que entrego el repo")
				repo.AssertCalled(t, "CrearSector", c.entrada)
			}
		})
	}
}
func TestSectorService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockSectorRepo)
	repo.On("BuscarSectorPorID", 999).Return(models.Sector{}, false)
	svc := service.NuevaSectorService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
