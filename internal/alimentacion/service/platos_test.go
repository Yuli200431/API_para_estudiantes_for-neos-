package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/storage"
)

// --- MOCK DEL REPOSITORIO PLATO ---
type platoRepoMock struct {
	mock.Mock
}

func (m *platoRepoMock) ListarPlatos() []models.Plato {
	args := m.Called()
	return args.Get(0).([]models.Plato)
}

func (m *platoRepoMock) BuscarPlatoPorID(id int) (models.Plato, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Plato), args.Bool(1)
}

func (m *platoRepoMock) CrearPlato(p models.Plato) models.Plato {
	args := m.Called(p)
	return args.Get(0).(models.Plato)
}

func (m *platoRepoMock) ActualizarPlato(id int, p models.Plato) (models.Plato, bool) {
	args := m.Called(id, p)
	return args.Get(0).(models.Plato), args.Bool(1)
}

func (m *platoRepoMock) BorrarPlato(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.PlatoRepository = (*platoRepoMock)(nil)


// --- DATOS DE PRUEBA ---
func platoValido() models.Plato {
	return models.Plato{
		NombrePlato:  "Encebollado",
		Descripcion:  "Tradicional con yuca y pescado",
		Categoria:    "Sopas / Desayunos",
		Precio:       3.50,
		MenuDiarioID: 1,
	}
}


// ==========================================================
// 1. CREAR
// ==========================================================
func TestPlatoService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       models.Plato
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "plato valido",
			entrada:       platoValido(),
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre: "nombre de plato vacio",
			entrada: func() models.Plato {
				p := platoValido()
				p.NombrePlato = "   "
				return p
			}(),
			errEsperado:   service.ErrNombrePlatoVacio,
			debePersistir: false,
		},
		{
			nombre: "descripcion vacia",
			entrada: func() models.Plato {
				p := platoValido()
				p.Descripcion = ""
				return p
			}(),
			errEsperado:   service.ErrDescripcionVacio,
			debePersistir: false,
		},
		{
			nombre: "categoria vacia",
			entrada: func() models.Plato {
				p := platoValido()
				p.Categoria = ""
				return p
			}(),
			errEsperado:   service.ErrCategoriaVacia,
			debePersistir: false,
		},
		{
			nombre: "precio menor o igual a cero",
			entrada: func() models.Plato {
				p := platoValido()
				p.Precio = 0
				return p
			}(),
			errEsperado:   service.ErrPrecioVacio,
			debePersistir: false,
		},
		{
			nombre: "menu diario id cero",
			entrada: func() models.Plato {
				p := platoValido()
				p.MenuDiarioID = 0
				return p
			}(),
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
	}


	for _, c := range casos {

		t.Run(c.nombre, func(t *testing.T) {

			repo := new(platoRepoMock)


			if c.debePersistir {

				guardado := c.entrada
				guardado.ID = 77

				repo.On("CrearPlato", c.entrada).
					Return(guardado)
			}


			svc := service.NuevoPlatoService(repo)


			creado, err := svc.Crear(c.entrada)


			if c.errEsperado != nil {

				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearPlato")

			} else {

				require.NoError(t, err)

				require.Equal(t, 77, creado.ID)

				repo.AssertCalled(t, "CrearPlato", c.entrada)
			}
		})
	}
}



// ==========================================================
// 2. OBTENER
// ==========================================================
func TestPlatoService_Obtener(t *testing.T) {


	t.Run("obtener plato exitoso", func(t *testing.T) {


		repo := new(platoRepoMock)


		esperado := platoValido()
		esperado.ID = 12


		repo.On("BuscarPlatoPorID", 12).
			Return(esperado, true)


		svc := service.NuevoPlatoService(repo)


		resultado, err := svc.Obtener(12)


		require.NoError(t, err)

		require.Equal(t, esperado, resultado)


		repo.AssertExpectations(t)

	})


	t.Run("plato no encontrado", func(t *testing.T) {


		repo := new(platoRepoMock)


		repo.On("BuscarPlatoPorID", 999).
			Return(models.Plato{}, false)


		svc := service.NuevoPlatoService(repo)


		_, err := svc.Obtener(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})
}



// ==========================================================
// 3. LISTAR
// ==========================================================
func TestPlatoService_Listar(t *testing.T) {


	repo := new(platoRepoMock)


	lista := []models.Plato{

		{
			ID: 1,
			NombrePlato: "Seco de Pollo",
		},

		{
			ID: 2,
			NombrePlato: "Ceviche",
		},
	}


	repo.On("ListarPlatos").
		Return(lista)


	svc := service.NuevoPlatoService(repo)


	resultado := svc.Listar()


	require.Len(t, resultado, 2)

	require.Equal(t, "Seco de Pollo", resultado[0].NombrePlato)


	repo.AssertExpectations(t)

}



// ==========================================================
// 4. ACTUALIZAR
// ==========================================================
func TestPlatoService_Actualizar(t *testing.T) {


	t.Run("actualizar plato exitoso", func(t *testing.T) {


		repo := new(platoRepoMock)


		p := platoValido()


		repo.On("ActualizarPlato", 1, p).
			Return(p, true)


		svc := service.NuevoPlatoService(repo)


		resultado, err := svc.Actualizar(1, p)


		require.NoError(t, err)

		require.Equal(t, p, resultado)


		repo.AssertExpectations(t)

	})


	t.Run("actualizar fallo por validacion", func(t *testing.T) {


		repo := new(platoRepoMock)


		p := platoValido()

		p.NombrePlato = ""


		svc := service.NuevoPlatoService(repo)


		_, err := svc.Actualizar(1, p)


		require.ErrorIs(t, err, service.ErrNombrePlatoVacio)


		repo.AssertNotCalled(t, "ActualizarPlato")

	})


	t.Run("actualizar id no encontrado", func(t *testing.T) {


		repo := new(platoRepoMock)


		p := platoValido()


		repo.On("ActualizarPlato", 404, p).
			Return(models.Plato{}, false)



		svc := service.NuevoPlatoService(repo)


		_, err := svc.Actualizar(404, p)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}



// ==========================================================
// 5. BORRAR
// ==========================================================
func TestPlatoService_Borrar(t *testing.T) {


	t.Run("borrar plato exitoso", func(t *testing.T) {


		repo := new(platoRepoMock)


		repo.On("BorrarPlato", 1).
			Return(true)


		svc := service.NuevoPlatoService(repo)


		err := svc.Borrar(1)


		require.NoError(t, err)


		repo.AssertExpectations(t)

	})


	t.Run("plato no encontrado", func(t *testing.T) {


		repo := new(platoRepoMock)


		repo.On("BorrarPlato", 999).
			Return(false)


		svc := service.NuevoPlatoService(repo)


		err := svc.Borrar(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}