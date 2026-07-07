package storage_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	alimentacionModels "for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/storage"
)

// abrirDBPrueba abre una base de datos SQLite en memoria, solo para este
// test. Se crea y se destruye en cada ejecucion, no toca el archivo .db real.
func abrirDBPrueba(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&alimentacionModels.Alimentacion{})
	require.NoError(t, err)

	return db
}

// TestAlimentacionGORM_CrearYBuscar prueba que crear un registro y
// buscarlo despues por ID realmente persiste en la base de datos.
func TestAlimentacionGORM_CrearYBuscar(t *testing.T) {
	db := abrirDBPrueba(t)
	repo := storage.NuevoAlimentacionGORM(db)

	creado := repo.CrearAlimentacion(alimentacionModels.Alimentacion{
		NombreLocal:    "Comedor ULEAM",
		Descripcion:    "Comida economica para estudiantes",
		PrecioPromedio: 2.50,
	})
	require.NotZero(t, creado.ID, "esperaba un ID asignado por GORM")

	encontrado, ok := repo.BuscarAlimentacionPorID(creado.ID)
	require.True(t, ok, "deberia encontrar el registro recien creado")
	require.Equal(t, "Comedor ULEAM", encontrado.NombreLocal)
}

// TestAlimentacionGORM_Listar prueba que listar refleja lo creado.
func TestAlimentacionGORM_Listar(t *testing.T) {
	db := abrirDBPrueba(t)
	repo := storage.NuevoAlimentacionGORM(db)

	repo.CrearAlimentacion(alimentacionModels.Alimentacion{NombreLocal: "Comedor A", PrecioPromedio: 2.0})
	repo.CrearAlimentacion(alimentacionModels.Alimentacion{NombreLocal: "Comedor B", PrecioPromedio: 3.0})

	lista := repo.ListarAlimentaciones()
	require.Len(t, lista, 2)
}

// TestAlimentacionGORM_BuscarInexistente: id que no existe -> ok=false.
func TestAlimentacionGORM_BuscarInexistente(t *testing.T) {
	db := abrirDBPrueba(t)
	repo := storage.NuevoAlimentacionGORM(db)

	_, ok := repo.BuscarAlimentacionPorID(999)
	require.False(t, ok)
}