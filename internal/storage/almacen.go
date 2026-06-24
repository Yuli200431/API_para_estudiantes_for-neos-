package storage

import (
	"for-neos-api/internal/vivienda/models"
)

type ViviendaRepository interface {
	ListarViviendas() []models.Vivienda
	BuscarViviendaPorID(id int) (models.Vivienda, bool)
	CrearVivienda(v models.Vivienda) models.Vivienda
	ActualizarVivienda(id int, datos models.Vivienda) (models.Vivienda, bool)
	BorrarVivienda(id int) bool
}

type FotoRepository interface {
	ListarFotos() []models.Foto
	BuscarFotoPorID(id int) (models.Foto, bool)
	CrearFoto(f models.Foto) models.Foto
	ActualizarFoto(id int, datos models.Foto) (models.Foto, bool)
	BorrarFoto(id int) bool
}

type AplicarViviendaRepository interface {
	ListarAplicarViviendas() []models.AplicarVivienda
	BuscarAplicarViviendaPorID(id int) (models.AplicarVivienda, bool)
	CrearAplicarVivienda(a models.AplicarVivienda) models.AplicarVivienda
	ActualizarAplicarVivienda(id int, datos models.AplicarVivienda) (models.AplicarVivienda, bool)
	BorrarAplicarVivienda(id int) bool
}

type SectorRepository interface {
	ListarSectores() []models.Sector
	BuscarSectorPorID(id int) (models.Sector, bool)
	CrearSector(e models.Sector) models.Sector
	ActualizarSector(id int, datos models.Sector) (models.Sector, bool)
	BorrarSector(id int) bool
}

type Almacen interface {
	ViviendaRepository
	FotoRepository
	AplicarViviendaRepository
	SectorRepository
}

//var _ Almacen = (*AlmacenStorage)(nil) NO ELIMINAR SE UTILIZARA MAS ADELANTE
