package storage

import (
	"gorm.io/gorm"

	alimentacionModels "for-neos-api/internal/alimentacion/models"
)

// AlimentacionGORM implementa AlimentacionRepository usando GORM sobre SQLite.
type AlimentacionGORM struct {
	db *gorm.DB
}

func NuevoAlimentacionGORM(db *gorm.DB) *AlimentacionGORM {
	return &AlimentacionGORM{db: db}
}

func (a *AlimentacionGORM) ListarAlimentaciones() []alimentacionModels.Alimentacion {
	var alimentaciones []alimentacionModels.Alimentacion
	a.db.Find(&alimentaciones)
	return alimentaciones
}

func (a *AlimentacionGORM) BuscarAlimentacionPorID(id int) (alimentacionModels.Alimentacion, bool) {
	var al alimentacionModels.Alimentacion
	if err := a.db.First(&al, id).Error; err != nil {
		return alimentacionModels.Alimentacion{}, false
	}
	return al, true
}

func (a *AlimentacionGORM) CrearAlimentacion(al alimentacionModels.Alimentacion) alimentacionModels.Alimentacion {
	a.db.Create(&al)
	return al
}

func (a *AlimentacionGORM) ActualizarAlimentacion(id int, datos alimentacionModels.Alimentacion) (alimentacionModels.Alimentacion, bool) {
	var existente alimentacionModels.Alimentacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return alimentacionModels.Alimentacion{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlimentacionGORM) BorrarAlimentacion(id int) bool {
	res := a.db.Delete(&alimentacionModels.Alimentacion{}, id)
	return res.RowsAffected > 0
}

var _ AlimentacionRepository = (*AlimentacionGORM)(nil)