package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/storage"
)

// --- MOCK DEL REPOSITORIO MENU DIARIO ---
type menuDiarioRepoMock struct {
	mock.Mock
}

func (m *menuDiarioRepoMock) ListarMenuDiarios() []models.MenuDiario {
	args := m.Called()
	return args.Get(0).([]models.MenuDiario)
}

func (m *menuDiarioRepoMock) BuscarMenuDiarioPorID(id int) (models.MenuDiario, bool) {
	args := m.Called(id)
	return args.Get(0).(models.MenuDiario), args.Bool(1)
}

func (m *menuDiarioRepoMock) CrearMenuDiario(menu models.MenuDiario) models.MenuDiario {
	args := m.Called(menu)
	return args.Get(0).(models.MenuDiario)
}

func (m *menuDiarioRepoMock) ActualizarMenuDiario(id int, menu models.MenuDiario) (models.MenuDiario, bool) {
	args := m.Called(id, menu)
	return args.Get(0).(models.MenuDiario), args.Bool(1)
}

func (m *menuDiarioRepoMock) BorrarMenuDiario(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.MenuDiarioRepository = (*menuDiarioRepoMock)(nil)


// --- DATOS DE PRUEBA ---
func menuDiarioValido() models.MenuDiario {
	return models.MenuDiario{
		Fecha:          "2026-07-08",
		AlimentacionID: 1,
	}
}


// ==========================================================
// 1. CREAR
// ==========================================================
func TestMenuDiarioService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       models.MenuDiario
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "menu diario valido",
			entrada:       menuDiarioValido(),
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre: "fecha vacia",
			entrada: func() models.MenuDiario {
				m := menuDiarioValido()
				m.Fecha = "   "
				return m
			}(),
			errEsperado:   service.ErrFechaVacia,
			debePersistir: false,
		},
		{
			nombre: "alimentacion id cero",
			entrada: func() models.MenuDiario {
				m := menuDiarioValido()
				m.AlimentacionID = 0
				return m
			}(),
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
	}


	for _, c := range casos {

		t.Run(c.nombre, func(t *testing.T) {

			repo := new(menuDiarioRepoMock)


			if c.debePersistir {

				guardado := c.entrada
				guardado.ID = 100

				repo.On("CrearMenuDiario", c.entrada).
					Return(guardado)
			}


			svc := service.NuevoMenuDiarioService(repo)

			creado, err := svc.Crear(c.entrada)


			if c.errEsperado != nil {

				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearMenuDiario")

			} else {

				require.NoError(t, err)
				require.Equal(t, 100, creado.ID)

				repo.AssertCalled(t, "CrearMenuDiario", c.entrada)
			}

		})
	}
}


// ==========================================================
// 2. OBTENER
// ==========================================================
func TestMenuDiarioService_Obtener(t *testing.T) {


	t.Run("obtener exitoso", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		esperado := menuDiarioValido()
		esperado.ID = 5


		repo.On("BuscarMenuDiarioPorID", 5).
			Return(esperado, true)


		svc := service.NuevoMenuDiarioService(repo)


		resultado, err := svc.Obtener(5)


		require.NoError(t, err)
		require.Equal(t, esperado, resultado)

		repo.AssertExpectations(t)

	})


	t.Run("menu no encontrado", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		repo.On("BuscarMenuDiarioPorID", 999).
			Return(models.MenuDiario{}, false)


		svc := service.NuevoMenuDiarioService(repo)


		_, err := svc.Obtener(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)

		repo.AssertExpectations(t)

	})
}



// ==========================================================
// 3. LISTAR
// ==========================================================
func TestMenuDiarioService_Listar(t *testing.T) {


	repo := new(menuDiarioRepoMock)


	lista := []models.MenuDiario{

		{
			ID: 1,
			Fecha: "2026-07-08",
			AlimentacionID: 1,
		},

		{
			ID: 2,
			Fecha: "2026-07-09",
			AlimentacionID: 1,
		},
	}


	repo.On("ListarMenuDiarios").
		Return(lista)


	svc := service.NuevoMenuDiarioService(repo)


	resultado := svc.Listar()


	require.Len(t, resultado, 2)
	require.Equal(t, "2026-07-08", resultado[0].Fecha)


	repo.AssertExpectations(t)

}



// ==========================================================
// 4. ACTUALIZAR
// ==========================================================
func TestMenuDiarioService_Actualizar(t *testing.T) {


	t.Run("actualizar exitoso", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)

		m := menuDiarioValido()


		repo.On("ActualizarMenuDiario", 1, m).
			Return(m, true)


		svc := service.NuevoMenuDiarioService(repo)


		resultado, err := svc.Actualizar(1, m)


		require.NoError(t, err)
		require.Equal(t, m, resultado)


		repo.AssertExpectations(t)

	})



	t.Run("error por validacion", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		m := menuDiarioValido()

		m.Fecha = ""


		svc := service.NuevoMenuDiarioService(repo)


		_, err := svc.Actualizar(1, m)


		require.ErrorIs(t, err, service.ErrFechaVacia)


		repo.AssertNotCalled(t, "ActualizarMenuDiario")

	})



	t.Run("id no encontrado", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		m := menuDiarioValido()


		repo.On("ActualizarMenuDiario", 404, m).
			Return(models.MenuDiario{}, false)



		svc := service.NuevoMenuDiarioService(repo)


		_, err := svc.Actualizar(404, m)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}



// ==========================================================
// 5. BORRAR
// ==========================================================
func TestMenuDiarioService_Borrar(t *testing.T) {


	t.Run("borrar exitoso", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		repo.On("BorrarMenuDiario", 1).
			Return(true)


		svc := service.NuevoMenuDiarioService(repo)


		err := svc.Borrar(1)


		require.NoError(t, err)


		repo.AssertExpectations(t)

	})



	t.Run("menu no encontrado", func(t *testing.T) {


		repo := new(menuDiarioRepoMock)


		repo.On("BorrarMenuDiario", 999).
			Return(false)



		svc := service.NuevoMenuDiarioService(repo)


		err := svc.Borrar(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}