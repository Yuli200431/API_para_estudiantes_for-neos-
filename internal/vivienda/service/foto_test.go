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

type mockFotoRepo struct {
	mock.Mock
}

func (m *mockFotoRepo) ListarFotos() []models.Foto {
	args := m.Called()
	return args.Get(0).([]models.Foto)
}

func (m *mockFotoRepo) BuscarFotoPorID(id int) (models.Foto, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Foto), args.Bool(1)
}

func (m *mockFotoRepo) CrearFoto(f models.Foto) models.Foto {
	args := m.Called(f)
	return args.Get(0).(models.Foto)
}

func (m *mockFotoRepo) ActualizarFoto(id int, datos models.Foto) (models.Foto, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Foto), args.Bool(1)
}

func (m *mockFotoRepo) BorrarFoto(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.FotoRepository = (*mockFotoRepo)(nil)

//Primer Test

func TestFotoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Foto
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "url vacia",
			entrada:       models.Foto{URL: "   "},
			errEsperado:   service.ErrURLVacio,
			debePersistir: false,
		},
		{
			nombre:        "vivienda vacia",
			entrada:       models.Foto{ViviendaID: 0, URL: "https://example.com/image.jpg"},
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
		{
			nombre:        "foto valida",
			entrada:       models.Foto{URL: "https://example.com/image.jpg", ViviendaID: 1},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(mockFotoRepo)
			if c.debePersistir {
				guardado := c.entrada
				guardado.FotoID = 42
				repo.On("CrearFoto", c.entrada).Return(guardado)
			}
			svc := service.NuevaFotoService(repo)

			creado, err := svc.Crear(c.entrada)
			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearFoto")
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.FotoID, "el service debe devolver el producto que entrego el repo")
				repo.AssertCalled(t, "CrearFoto", c.entrada)
			}
		})
	}
}
func TestFotoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockFotoRepo)
	repo.On("BuscarFotoPorID", 999).Return(models.Foto{}, false)
	svc := service.NuevaFotoService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
	repo.AssertExpectations(t)
}
