package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/alimentacion/service"
	"for-neos-api/internal/storage"
)

// --- MOCK DEL REPOSITORIO RESEÑA ---
type resenaRepoMock struct {
	mock.Mock
}

func (m *resenaRepoMock) ListarResenas() []models.Resena {
	args := m.Called()
	return args.Get(0).([]models.Resena)
}

func (m *resenaRepoMock) BuscarResenaPorID(id int) (models.Resena, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Resena), args.Bool(1)
}

func (m *resenaRepoMock) CrearResena(r models.Resena) models.Resena {
	args := m.Called(r)
	return args.Get(0).(models.Resena)
}

func (m *resenaRepoMock) ActualizarResena(id int, r models.Resena) (models.Resena, bool) {
	args := m.Called(id, r)
	return args.Get(0).(models.Resena), args.Bool(1)
}

func (m *resenaRepoMock) BorrarResena(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var _ storage.ResenaRepository = (*resenaRepoMock)(nil)


// --- DATOS DE PRUEBA ---
func resenaValida() models.Resena {
	return models.Resena{
		Comentario:     "Excelente comida y muy buena atención.",
		Calificacion:   5,
		AlimentacionID: 1,
	}
}


// ==========================================================
// 1. CREAR
// ==========================================================
func TestResenaService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       models.Resena
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "resena valida",
			entrada:       resenaValida(),
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre: "comentario vacio",
			entrada: func() models.Resena {
				r := resenaValida()
				r.Comentario = "   "
				return r
			}(),
			errEsperado:   service.ErrComentarioVacio,
			debePersistir: false,
		},
		{
			nombre: "calificacion menor o igual a cero",
			entrada: func() models.Resena {
				r := resenaValida()
				r.Calificacion = 0
				return r
			}(),
			errEsperado:   service.ErrCalificacionVacia,
			debePersistir: false,
		},
		{
			nombre: "alimentacion id cero",
			entrada: func() models.Resena {
				r := resenaValida()
				r.AlimentacionID = 0
				return r
			}(),
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
	}


	for _, c := range casos {

		t.Run(c.nombre, func(t *testing.T) {

			repo := new(resenaRepoMock)


			if c.debePersistir {

				guardado := c.entrada
				guardado.ID = 25

				repo.On("CrearResena", c.entrada).
					Return(guardado)
			}


			svc := service.NuevaResenaService(repo)


			creado, err := svc.Crear(c.entrada)


			if c.errEsperado != nil {

				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearResena")

			} else {

				require.NoError(t, err)

				require.Equal(t, 25, creado.ID)

				repo.AssertCalled(t, "CrearResena", c.entrada)

			}

		})
	}
}


// ==========================================================
// 2. OBTENER
// ==========================================================
func TestResenaService_Obtener(t *testing.T) {


	t.Run("obtener resena exitoso", func(t *testing.T) {


		repo := new(resenaRepoMock)


		esperado := resenaValida()
		esperado.ID = 7


		repo.On("BuscarResenaPorID", 7).
			Return(esperado, true)


		svc := service.NuevaResenaService(repo)


		resultado, err := svc.Obtener(7)


		require.NoError(t, err)

		require.Equal(t, esperado, resultado)


		repo.AssertExpectations(t)

	})


	t.Run("resena no encontrada", func(t *testing.T) {


		repo := new(resenaRepoMock)


		repo.On("BuscarResenaPorID", 999).
			Return(models.Resena{}, false)


		svc := service.NuevaResenaService(repo)


		_, err := svc.Obtener(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}


// ==========================================================
// 3. LISTAR
// ==========================================================
func TestResenaService_Listar(t *testing.T) {


	repo := new(resenaRepoMock)


	lista := []models.Resena{

		{
			ID:           1,
			Comentario:  "Rico",
			Calificacion: 4,
		},

		{
			ID:           2,
			Comentario:  "Malo",
			Calificacion: 2,
		},
	}


	repo.On("ListarResenas").
		Return(lista)


	svc := service.NuevaResenaService(repo)


	resultado := svc.Listar()


	require.Len(t, resultado, 2)

	require.Equal(t, "Rico", resultado[0].Comentario)


	repo.AssertExpectations(t)

}


// ==========================================================
// 4. ACTUALIZAR
// ==========================================================
func TestResenaService_Actualizar(t *testing.T) {


	t.Run("actualizar resena exitoso", func(t *testing.T) {


		repo := new(resenaRepoMock)


		r := resenaValida()


		repo.On("ActualizarResena", 1, r).
			Return(r, true)


		svc := service.NuevaResenaService(repo)


		resultado, err := svc.Actualizar(1, r)


		require.NoError(t, err)

		require.Equal(t, r, resultado)


		repo.AssertExpectations(t)

	})



	t.Run("actualizar fallo por validacion", func(t *testing.T) {


		repo := new(resenaRepoMock)


		r := resenaValida()

		r.Comentario = ""


		svc := service.NuevaResenaService(repo)


		_, err := svc.Actualizar(1, r)


		require.ErrorIs(t, err, service.ErrComentarioVacio)


		repo.AssertNotCalled(t, "ActualizarResena")

	})



	t.Run("actualizar id no encontrado", func(t *testing.T) {


		repo := new(resenaRepoMock)


		r := resenaValida()


		repo.On("ActualizarResena", 404, r).
			Return(models.Resena{}, false)



		svc := service.NuevaResenaService(repo)


		_, err := svc.Actualizar(404, r)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}


// ==========================================================
// 5. BORRAR
// ==========================================================
func TestResenaService_Borrar(t *testing.T) {


	t.Run("borrar resena exitoso", func(t *testing.T) {


		repo := new(resenaRepoMock)


		repo.On("BorrarResena", 1).
			Return(true)


		svc := service.NuevaResenaService(repo)


		err := svc.Borrar(1)


		require.NoError(t, err)


		repo.AssertExpectations(t)

	})



	t.Run("resena no encontrada", func(t *testing.T) {


		repo := new(resenaRepoMock)


		repo.On("BorrarResena", 999).
			Return(false)


		svc := service.NuevaResenaService(repo)


		err := svc.Borrar(999)


		require.ErrorIs(t, err, service.ErrNoEncontrado)


		repo.AssertExpectations(t)

	})

}