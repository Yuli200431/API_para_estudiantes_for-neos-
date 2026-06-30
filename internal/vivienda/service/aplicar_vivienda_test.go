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

type mockAplicarRepo struct {
	mock.Mock
}

func (m *mockAplicarRepo) ListarAplicarViviendas() []models.AplicarVivienda {
	args := m.Called()
	return args.Get(0).([]models.AplicarVivienda)
}

func (m *mockAplicarRepo) BuscarAplicarViviendaPorID(id int) (models.AplicarVivienda, bool) {
	args := m.Called(id)
	return args.Get(0).(models.AplicarVivienda), args.Bool(1)
}

func (m *mockAplicarRepo) CrearAplicarVivienda(a models.AplicarVivienda) models.AplicarVivienda {
	args := m.Called(a)
	return args.Get(0).(models.AplicarVivienda)
}

func (m *mockAplicarRepo) ActualizarAplicarVivienda(id int, datos models.AplicarVivienda) (models.AplicarVivienda, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.AplicarVivienda), args.Bool(1)
}

func (m *mockAplicarRepo) BorrarAplicarVivienda(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.AplicarViviendaRepository = (*mockAplicarRepo)(nil)

// Primer Test
func TestAplicarViviendaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.AplicarVivienda
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "vivienda vacia",
			entrada:       models.AplicarVivienda{Estado: "Pendiente", ViviendaID: 0},
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
		{
			nombre:        "estado vacio",
			entrada:       models.AplicarVivienda{Estado: " "},
			errEsperado:   service.ErrEstadoVacio,
			debePersistir: false,
		},
		{
			nombre:        "aplicar valido",
			entrada:       models.AplicarVivienda{ViviendaID: 1, Estado: "aprobado"},
			errEsperado:   nil,
			debePersistir: true,
		},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(mockAplicarRepo)
			if c.debePersistir {
				guardado := c.entrada
				guardado.AplicarViviendaID = 42
				repo.On("CrearAplicarVivienda", c.entrada).Return(guardado)
			}
			svc := service.NuevaAplicarViviendaService(repo)
			creado, err := svc.Crear(c.entrada)
			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearAplicarVivienda")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.AplicarViviendaID, "el service debe devolver el producto que entrego el repo")
				repo.AssertCalled(t, "CrearAplicarVivienda", c.entrada)
			}
		})
	}
}
func TestAplicarViviendaService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockAplicarRepo)
	repo.On("BuscarAplicarViviendaPorID", 999).Return(models.AplicarVivienda{}, false)
	svc := service.NuevaAplicarViviendaService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
