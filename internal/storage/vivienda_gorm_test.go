package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"for-neos-api/internal/vivienda/models"
)

func TestSQLite_ListarFotos(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("error abriendo sqlite: %v", err)
	}

	err = gdb.AutoMigrate(&models.Foto{})

	if err != nil {
		t.Fatalf("error migrando: %v", err)
	}

	repo := NuevoAlmacenSQLite(gdb)

	repo.CrearFoto(models.Foto{FotoID: 1, ViviendaID: 1})
	lista := repo.ListarFotos()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 foto")
	}
}
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

func TestSQLite_ActualizarVivienda(t *testing.T) {
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

	actualizada, ok := repo.ActualizarVivienda(creada.ViviendaID, models.Vivienda{
		Titulo: "Casa Prueba 2",
	})

	if !ok {
		t.Fatal("No se pudo actualizar la vivienda")
	}

	if actualizada.Titulo != "Casa Prueba 2" {
		t.Errorf("titulo=%q esperaba=%q", actualizada.Titulo, "Casa Prueba 2")
	}
}

func TestSQLite_BorrarVivienda(t *testing.T) {
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

	repo.BorrarVivienda(creada.ViviendaID)

	_, ok := repo.BuscarViviendaPorID(creada.ViviendaID)

	if ok {
		t.Fatalf("se esperaba no encontrar la vivienda")
	}
}
