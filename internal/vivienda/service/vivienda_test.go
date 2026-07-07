// Aqui se prueban las reglas de negocio de las funciones del servicio.
// Se verifica que el servicio valide los datos antes de llamar al repositorio.
package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"  //probar cosas
	"github.com/stretchr/testify/mock"    //repositorio mock
	"github.com/stretchr/testify/require" //prueba cosa, pero si falla no sigue ejecutando las pruebas

	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"for-neos-api/internal/vivienda/service"
)

// Se crea la estructura mock (repositorio), donde tendrá todas las funcionas del paquete mock.Mock.
type mockViviendaRepo struct {
	mock.Mock
}

//Aquí se simula los metodos de las distintas fucniones del repositorio. Sin necesidad de conectarnos
//a la base de datos real. Todo para saber si el servicona funciona bien. Todo basado en el repositorio original.

// Función ListarViviendas, devuelve la lista de viviendas.
func (m *mockViviendaRepo) ListarViviendas() []models.Vivienda {
	////Obtiene el valor que fue configurado previamente en el test.
	args := m.Called()
	//Retorna el primer valor de la lista Vivienda.
	return args.Get(0).([]models.Vivienda)
}

// Funcion BuscarViviendasPorId, devuelve la vivienda con el id.
func (m *mockViviendaRepo) BuscarViviendaPorID(id int) (models.Vivienda, bool) {
	args := m.Called(id)
	//devuelve el primer valor configurado
	return args.Get(0).(models.Vivienda), args.Bool(1)
}

// Funcion CrearVivienda, crea una vivienda.
func (m *mockViviendaRepo) CrearVivienda(v models.Vivienda) models.Vivienda {
	args := m.Called(v)
	return args.Get(0).(models.Vivienda)
}

// Funcion ActualizarVivienda
func (m *mockViviendaRepo) ActualizarVivienda(id int, datos models.Vivienda) (models.Vivienda, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Vivienda), args.Bool(1)
}

// Funcion BorrarVivienda.
func (m *mockViviendaRepo) BorrarVivienda(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Verifica que todas las funciones del repositorio esten.
var _ storage.ViviendaRepository = (*mockViviendaRepo)(nil)

// Primer Test
// Sirve para probar el metodo de crear
func TestViviendaService_Crear(t *testing.T) {

	luz := true
	agua := true
	amueblado := true
	internet := true
	bañoPrivado := true
	mascotas := true
	garantia := true

	//Casos de prueba.
	//Casos de prueba.
	viviendaValida := models.Vivienda{
		Titulo:             "titulo",
		TipoVivienda:       "casa",
		Precio:             100,
		Garantia:           &garantia,
		PrecioGarantia:     50,
		Luz:                &luz,
		Agua:               &agua,
		Amueblado:          &amueblado,
		Internet:           &internet,
		BañoPrivado:        &bañoPrivado,
		NumeroHabitaciones: 2,
		Mascotas:           &mascotas,
		GeneroPreferido:    "cualquiera",
		ReglasConvivencia:  "reglas",
		Telefono:           "0999999999",
		Email:              "correo@test.com",
		Estado:             "Disponible",
		Comentario:         "comentario",
		ViviendaID:         1,
		SectorID:           1,
		ProveedorID:        1,
		Fotos: []models.Foto{
			{FotoID: 1},
		},
	}

	casos := []struct {
		nombre        string
		entrada       models.Vivienda
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre:        "vivienda valida",
			entrada:       viviendaValida,
			errEsperado:   nil,
			debePersistir: true,
		},
		{
			nombre: "titulo vacio",
			entrada: models.Vivienda{
				Titulo: "",
			},
			errEsperado:   service.ErrTituloVacio,
			debePersistir: false,
		},
		{
			nombre: "precio vacio",
			entrada: func() models.Vivienda {
				v := viviendaValida
				v.Precio = 0
				return v
			}(),
			errEsperado:   service.ErrPrecioVacio,
			debePersistir: false,
		},
		{
			nombre: "sector no encontrado",
			entrada: func() models.Vivienda {
				v := viviendaValida
				v.SectorID = 0
				return v
			}(),
			errEsperado:   service.ErrNoEncontrado,
			debePersistir: false,
		},
	}

	//Se recorre los casos de prueba. Subtest con su nombre
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			//Repositorio falso.
			repo := new(mockViviendaRepo)
			//Si se guarda en el repositorio, se devuelve una vivienda con el Id generado.
			if c.debePersistir {
				guardado := c.entrada
				guardado.ViviendaID = 42
				repo.On("CrearVivienda", c.entrada).Return(guardado)
			}

			//Se crea el servicio utilizando el repositorio falso.
			svc := service.NuevaViviendaService(repo)

			//Se llama el metodo Crear dandole la vivienda que se pasa.
			creado, err := svc.Crear(c.entrada)
			//Si error ocurre
			if c.errEsperado != nil {
				//Comprueba que el service devolvió exactamente el error esperado.
				require.ErrorIs(t, err, c.errEsperado)
				//Verifica que el repositorio nunca intentó guardar la vivienda.
				repo.AssertNotCalled(t, "CrearVivienda")

				//Si no se esperaba ningún error, se verifica que la vivienda se creó correctamente.
			} else {
				//Comprueba que el servicio devuelve la vivienda creada.
				require.NoError(t, err)
				assert.Equal(t, 42, creado.ViviendaID, "el service debe devolver el producto que entrego el repo")
				//Verifica que el repositorio guardó la vivienda.
				repo.AssertCalled(t, "CrearVivienda", c.entrada)
			}
		})

	}
}

// Segundo Test
// Sirve para probar que el servicio devuelva el error si no existe la vivienda.
func TestViviendaService_Obtener_NoEncontrado(t *testing.T) {
	//Repositorio falso.
	repo := new(mockViviendaRepo)
	//Si el ID es 999, no existe la vivienda.
	repo.On("BuscarViviendaPorID", 999).Return(models.Vivienda{}, false)

	//Se crea el servicio utilizando el repositorio falso.
	svc := service.NuevaViviendaService(repo)

	//Se llama al servicio con el ID 999. Solo se guarda el error.
	_, err := svc.Obtener(999)

	//Comprueba que el error es el esperado.
	require.ErrorIs(t, err, service.ErrNoEncontrado)
	//Verifica que el mock fue usado como esperábamos
	repo.AssertExpectations(t)
}
