package storage

import (
	"for-neos-api/internal/vivienda/models"
	transporteModels "for-neos-api/internal/transporte/models"
	alimentacionModels "for-neos-api/internal/alimentacion/models"
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

//Trnasporte

type CooperativaRepository interface {
	ListarCooperativas() []transporteModels.Cooperativa
	ObtenerCooperativaPorID(id uint) (transporteModels.Cooperativa, bool)
	AgregarCooperativa(c transporteModels.Cooperativa) transporteModels.Cooperativa
	ActualizarCooperativa(id uint, c transporteModels.Cooperativa) (transporteModels.Cooperativa, bool)
	EliminarCooperativa(id uint) bool
}

type ParadaBusRepository interface {
	ListarParadas() []transporteModels.ParadaBus
	ObtenerParadaPorID(id uint) (transporteModels.ParadaBus, bool)
	AgregarParada(p transporteModels.ParadaBus) transporteModels.ParadaBus
	ActualizarParada(id uint, p transporteModels.ParadaBus) (transporteModels.ParadaBus, bool)
	EliminarParada(id uint) bool
}

type RutaTransporteRepository interface {
	ListarRutas() []transporteModels.RutaTransporte
	ObtenerRutaPorID(id uint) (transporteModels.RutaTransporte, bool)
	AgregarRuta(r transporteModels.RutaTransporte) transporteModels.RutaTransporte
	ActualizarRuta(id uint, r transporteModels.RutaTransporte) (transporteModels.RutaTransporte, bool)
	EliminarRuta(id uint) bool
}

//Alimentacion

type AlimentacionRepository interface {
	ListarAlimentaciones() []alimentacionModels.Alimentacion
	BuscarAlimentacionPorID(id int) (alimentacionModels.Alimentacion, bool)
	CrearAlimentacion(a alimentacionModels.Alimentacion) alimentacionModels.Alimentacion
	ActualizarAlimentacion(id int, a alimentacionModels.Alimentacion) (alimentacionModels.Alimentacion, bool)
	BorrarAlimentacion(id int) bool
}

type MenuDiarioRepository interface {
	ListarMenuDiarios() []alimentacionModels.MenuDiario
	BuscarMenuDiarioPorID(id int) (alimentacionModels.MenuDiario, bool)
	CrearMenuDiario(m alimentacionModels.MenuDiario) alimentacionModels.MenuDiario
	ActualizarMenuDiario(id int, m alimentacionModels.MenuDiario) (alimentacionModels.MenuDiario, bool)
	BorrarMenuDiario(id int) bool
}

type PlatoRepository interface {
	ListarPlatos() []alimentacionModels.Plato
	BuscarPlatoPorID(id int) (alimentacionModels.Plato, bool)
	CrearPlato(p alimentacionModels.Plato) alimentacionModels.Plato
	ActualizarPlato(id int, p alimentacionModels.Plato) (alimentacionModels.Plato, bool)
	BorrarPlato(id int) bool
}

type ResenaRepository interface {
	ListarResenas() []alimentacionModels.Resena
	BuscarResenaPorID(id int) (alimentacionModels.Resena, bool)
	CrearResena(r alimentacionModels.Resena) alimentacionModels.Resena
	ActualizarResena(id int, r alimentacionModels.Resena) (alimentacionModels.Resena, bool)
	BorrarResena(id int) bool
}

//var _ Almacen = (*AlmacenStorage)(nil) NO ELIMINAR SE UTILIZARA MAS ADELANTE
