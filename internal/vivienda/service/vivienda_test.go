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

type mockViviendaRepo struct {
	mock.Mock
}

func (m *mockViviendaRepo) ListarViviendas() []models.Vivienda {
	args := m.Called()
	return args.Get(0).([]models.Vivienda)
}

func (m *mockViviendaRepo) BuscarViviendaPorID(id int) (models.Vivienda, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Vivienda), args.Bool(1)
}

func (m *mockViviendaRepo) CrearVivienda(v models.Vivienda) models.Vivienda {
	args := m.Called(v)
	return args.Get(0).(models.Vivienda)
}

func (m *mockViviendaRepo) ActualizarVivienda(id int, datos models.Vivienda) (models.Vivienda, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Vivienda), args.Bool(1)
}

func (m *mockViviendaRepo) BorrarVivienda(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.ViviendaRepository = (*mockViviendaRepo)(nil)

//Primer Test

func TestViviendaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Vivienda
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "titulo vacio",
			entrada:       models.Vivienda{Titulo: "   "},
			errEsperado:   service.ErrTituloVacio,
			debePersistir: false,
		},
		{
			nombre:        "titulo valida",
			entrada:       models.Vivienda{Titulo: "titulo", ViviendaID: 1},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(mockViviendaRepo)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ViviendaID = 42
				repo.On("CrearVivienda", c.entrada).Return(guardado)
			}
			svc := service.NuevaViviendaService(repo)

			creado, err := svc.Crear(c.entrada)
			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearVivienda")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.ViviendaID, "el service debe devolver el producto que entrego el repo")
				repo.AssertCalled(t, "CrearVivienda", c.entrada)
			}
		})

	}
}

func TestViviendaService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockViviendaRepo)
	repo.On("BuscarViviendaPorID", 999).Return(models.Vivienda{}, false)
	svc := service.NuevaViviendaService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
