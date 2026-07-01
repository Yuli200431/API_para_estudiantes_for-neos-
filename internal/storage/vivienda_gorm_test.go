package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"for-neos-api/internal/vivienda/models"
)

func TestSQLite_CrearYBuscarVivienda(t *testing.T) {

	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("error abriendo sqlite: %v", err)
	}

	err = gdb.AutoMigrate(&models.Vivienda{})

	if err != nil {
		t.Fatalf("error migrando: %v", err)
	}

	repo := NuevoAlmacenSQLite(gdb)

	creada := repo.CrearVivienda(models.Vivienda{Titulo: "Casa Prueba"})

	if creada.ViviendaID == 0 {
		t.Fatalf("esperaba ID generado")
	}

	encontrada, ok := repo.BuscarViviendaPorID(creada.ViviendaID)
	if !ok {
		t.Fatalf("no se encontro la vivienda")
	}
	if encontrada.Titulo != "Casa Prueba" {
		t.Errorf("titulo=%q esperaba=%q", encontrada.Titulo, "Casa Prueba")
	}

	lista := repo.ListarViviendas()

	if len(lista) != 1 {
		t.Fatalf("esperaba 1 vivienda")
	}
}

func TestSQLite_BuscarViviendaInexistente(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("error abriendo sqlite: %v", err)
	}

	err = gdb.AutoMigrate(&models.Vivienda{})

	if err != nil {
		t.Fatalf("error migrando: %v", err)
	}

	repo := NuevoAlmacenSQLite(gdb)

	_, ok := repo.BuscarViviendaPorID(999)

	if ok {
		t.Fatalf("se esperaba no encontrar la vivienda")
	}
}
