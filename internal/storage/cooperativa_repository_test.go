package storage_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	transporteModels "for-neos-api/internal/transporte/models"
	"for-neos-api/internal/storage"
)

// abrirDBMemoria crea una base de datos SQLite en memoria, solo para el test.
// No usa el archivo for-neos-api.db real.
func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&transporteModels.Cooperativa{})
	require.NoError(t, err)

	return db
}

// TestCooperativaRepository_CrearYListar verifica que crear una cooperativa
// y luego listarla la refleje correctamente, usando GORM contra sqlite :memory:.
func TestCooperativaRepository_CrearYListar(t *testing.T) {
	db := abrirDBMemoria(t)
	almacen := storage.NuevoAlmacenSQLite(db)

	nueva := transporteModels.Cooperativa{
		Nombre:      "Coop Manta",
		Telefono:    "099123456",
		Descripcion: "Cooperativa de prueba",
	}

	creada := almacen.CrearCooperativa(nueva)
	require.NotZero(t, creada.ID, "GORM debió asignar un ID autogenerado")

	lista := almacen.ListarCooperativas()
	require.Len(t, lista, 1)
	require.Equal(t, "Coop Manta", lista[0].Nombre)

	// También probamos buscar por ID, que crear → buscar lo refleje.
	encontrada, ok := almacen.BuscarCooperativaPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "Coop Manta", encontrada.Nombre)
}