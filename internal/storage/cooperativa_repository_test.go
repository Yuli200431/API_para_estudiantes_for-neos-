package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	transporteModels "for-neos-api/internal/transporte/models"
)

func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error abriendo sqlite: %v", err)
	}
	err = gdb.AutoMigrate(&transporteModels.Cooperativa{})
	if err != nil {
		t.Fatalf("error migrando: %v", err)
	}
	return gdb
}

func TestSQLite_CrearYBuscarCooperativa(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBMemoria(t))

	creada := repo.CrearCooperativa(transporteModels.Cooperativa{
		Nombre:      "Coop Manta",
		Telefono:    "099123456",
		Descripcion: "Cooperativa de prueba",
	})

	if creada.ID == 0 {
		t.Fatalf("esperaba ID generado")
	}

	encontrada, ok := repo.BuscarCooperativaPorID(creada.ID)
	if !ok {
		t.Fatalf("no se encontró la cooperativa")
	}
	if encontrada.Nombre != "Coop Manta" {
		t.Errorf("nombre=%q esperaba=%q", encontrada.Nombre, "Coop Manta")
	}

	lista := repo.ListarCooperativas()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 cooperativa")
	}
}

func TestSQLite_BuscarCooperativaInexistente(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBMemoria(t))

	_, ok := repo.BuscarCooperativaPorID(999)
	if ok {
		t.Fatalf("se esperaba no encontrar la cooperativa")
	}
}

func TestSQLite_ActualizarCooperativa(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBMemoria(t))

	creada := repo.CrearCooperativa(transporteModels.Cooperativa{
		Nombre:      "Coop Manta",
		Telefono:    "099123456",
		Descripcion: "Cooperativa de prueba",
	})

	actualizada, ok := repo.ActualizarCooperativa(creada.ID, transporteModels.Cooperativa{
		Nombre:      "Coop Costa Azul",
		Telefono:    "099999999",
		Descripcion: "Actualizada",
	})

	if !ok {
		t.Fatal("no se pudo actualizar la cooperativa")
	}
	if actualizada.Nombre != "Coop Costa Azul" {
		t.Errorf("nombre=%q esperaba=%q", actualizada.Nombre, "Coop Costa Azul")
	}
}

func TestSQLite_BorrarCooperativa(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBMemoria(t))

	creada := repo.CrearCooperativa(transporteModels.Cooperativa{
		Nombre:      "Coop Manta",
		Telefono:    "099123456",
		Descripcion: "Cooperativa de prueba",
	})

	repo.BorrarCooperativa(creada.ID)

	_, ok := repo.BuscarCooperativaPorID(creada.ID)
	if ok {
		t.Fatalf("se esperaba no encontrar la cooperativa después de borrarla")
	}
}

func TestSQLite_ListarCooperativas(t *testing.T) {
	repo := NuevoAlmacenSQLite(abrirDBMemoria(t))

	repo.CrearCooperativa(transporteModels.Cooperativa{Nombre: "Coop 1", Telefono: "099111111", Descripcion: "desc1"})
	repo.CrearCooperativa(transporteModels.Cooperativa{Nombre: "Coop 2", Telefono: "099222222", Descripcion: "desc2"})

	lista := repo.ListarCooperativas()
	if len(lista) != 2 {
		t.Fatalf("esperaba 2 cooperativas, obtuvo %d", len(lista))
	}
}