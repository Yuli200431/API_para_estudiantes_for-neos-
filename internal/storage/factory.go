package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	alimentacionModels "for-neos-api/internal/alimentacion/models"
	transporteModels "for-neos-api/internal/transporte/models"
	usuarioModels "for-neos-api/internal/usuario/models"
	viviendaModels "for-neos-api/internal/vivienda/models"
)

// Recursos agrupa todo lo que la capa de almacenamiento expone a la aplicacion.
type Recursos struct {
	Almacen      Almacen
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

// Inicializar centraliza TODO el plumbing de almacenamiento (patron Factory).
func Inicializar(rutaDB, backend string) (*Recursos, error) {
	// 1. GORM es el DUENO DEL ESQUEMA: abre, migra y siembra.
	gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("abrir GORM: %w", err)
	}
	if err := gdb.AutoMigrate(
		&viviendaModels.Vivienda{},
		&viviendaModels.Foto{},
		&viviendaModels.AplicarVivienda{},
		&viviendaModels.Sector{},
		&alimentacionModels.Alimentacion{},
		&alimentacionModels.MenuDiario{},
		&alimentacionModels.Plato{},
		&alimentacionModels.Resena{},
		&transporteModels.Cooperativa{},
		&transporteModels.RutaTransporte{},
		&transporteModels.ParadaBus{},
		&usuarioModels.Usuario{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}
	almacenGorm := NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend.
	var almacen Almacen
	var sdb *sql.DB
	backendUsado := "gorm"
	switch backend {
	case "sqlc":
		sdb, err = sql.Open("sqlite", rutaDB)
		if err != nil {
			return nil, fmt.Errorf("abrir sql.DB para sqlc: %w", err)
		}
		almacen = NuevoAlmacenSQLC(sdb)
		backendUsado = "sqlc"
	default:
		almacen = almacenGorm
	}

	// 3. Usuarios siempre en GORM.
	usuarios := NewUsuarioGORM(gdb)

	// 4. Cierre ordenado.
	cerrar := func() error {
		if sdb != nil {
			if err := sdb.Close(); err != nil {
				return err
			}
		}
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Almacen:      almacen,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}
