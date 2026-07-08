package storage

import (
	"gorm.io/gorm"

	alimentacionModels "for-neos-api/internal/alimentacion/models"
	transporteModels "for-neos-api/internal/transporte/models"
	viviendaModels "for-neos-api/internal/vivienda/models"

	usuarioModels "for-neos-api/internal/usuario/models"

	"log"

	"golang.org/x/crypto/bcrypt"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// VIVIENDA
// =========================================================

func (a *AlmacenSQLite) ListarViviendas() []viviendaModels.Vivienda {
	var viviendas []viviendaModels.Vivienda
	a.db.Find(&viviendas)
	return viviendas
}

func (a *AlmacenSQLite) BuscarViviendaPorID(id int) (viviendaModels.Vivienda, bool) {
	var v viviendaModels.Vivienda
	if err := a.db.First(&v, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return viviendaModels.Vivienda{}, false
	}
	return v, true
}

func (a *AlmacenSQLite) CrearVivienda(v viviendaModels.Vivienda) viviendaModels.Vivienda {
	a.db.Create(&v) // GORM rellena el ID autogenerado en &v
	return v
}

func (a *AlmacenSQLite) ActualizarVivienda(id int, datos viviendaModels.Vivienda) (viviendaModels.Vivienda, bool) {
	var existente viviendaModels.Vivienda
	if err := a.db.First(&existente, id).Error; err != nil {
		return viviendaModels.Vivienda{}, false
	}
	datos.ViviendaID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarVivienda(id int) bool {
	res := a.db.Delete(&viviendaModels.Vivienda{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// FOTO
// =========================================================

func (a *AlmacenSQLite) ListarFotos() []viviendaModels.Foto {
	var fotos []viviendaModels.Foto
	a.db.Find(&fotos)
	return fotos
}

func (a *AlmacenSQLite) BuscarFotoPorID(id int) (viviendaModels.Foto, bool) {
	var f viviendaModels.Foto
	if err := a.db.First(&f, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return viviendaModels.Foto{}, false
	}
	return f, true
}

func (a *AlmacenSQLite) CrearFoto(f viviendaModels.Foto) viviendaModels.Foto {
	a.db.Create(&f) // GORM rellena el ID autogenerado en &f
	return f
}

func (a *AlmacenSQLite) ActualizarFoto(id int, datos viviendaModels.Foto) (viviendaModels.Foto, bool) {
	var existente viviendaModels.Foto
	if err := a.db.First(&existente, id).Error; err != nil {
		return viviendaModels.Foto{}, false
	}
	datos.FotoID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarFoto(id int) bool {
	res := a.db.Delete(&viviendaModels.Foto{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// APLICAR VIVIENDA
// =========================================================

func (a *AlmacenSQLite) ListarAplicarViviendas() []viviendaModels.AplicarVivienda {
	var aplicarViviendas []viviendaModels.AplicarVivienda
	a.db.Find(&aplicarViviendas)
	return aplicarViviendas
}

func (a *AlmacenSQLite) BuscarAplicarViviendaPorID(id int) (viviendaModels.AplicarVivienda, bool) {
	var av viviendaModels.AplicarVivienda
	if err := a.db.First(&av, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return viviendaModels.AplicarVivienda{}, false
	}
	return av, true
}

func (a *AlmacenSQLite) CrearAplicarVivienda(av viviendaModels.AplicarVivienda) viviendaModels.AplicarVivienda {
	if err := a.db.Create(&av).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return viviendaModels.AplicarVivienda{}
	}
	return av
}

func (a *AlmacenSQLite) ActualizarAplicarVivienda(id int, datos viviendaModels.AplicarVivienda) (viviendaModels.AplicarVivienda, bool) {
	var existente viviendaModels.AplicarVivienda
	if err := a.db.First(&existente, id).Error; err != nil {
		return viviendaModels.AplicarVivienda{}, false
	}
	datos.AplicarViviendaID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarAplicarVivienda(id int) bool {
	res := a.db.Delete(&viviendaModels.AplicarVivienda{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SECTOR
// =========================================================

func (a *AlmacenSQLite) ListarSectores() []viviendaModels.Sector {
	var sectores []viviendaModels.Sector
	a.db.Find(&sectores)
	return sectores
}

func (a *AlmacenSQLite) BuscarSectorPorID(id int) (viviendaModels.Sector, bool) {
	var s viviendaModels.Sector
	if err := a.db.First(&s, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return viviendaModels.Sector{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSector(s viviendaModels.Sector) viviendaModels.Sector {
	a.db.Create(&s) // GORM rellena el ID autogenerado en &s
	return s
}

func (a *AlmacenSQLite) ActualizarSector(id int, datos viviendaModels.Sector) (viviendaModels.Sector, bool) {
	var existente viviendaModels.Sector
	if err := a.db.First(&existente, id).Error; err != nil {
		return viviendaModels.Sector{}, false
	}
	datos.SectorID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSector(id int) bool {
	res := a.db.Delete(&viviendaModels.Sector{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// ALIMENTACION
// =========================================================

func (a *AlmacenSQLite) ListarAlimentaciones() []alimentacionModels.Alimentacion {
	var alimentaciones []alimentacionModels.Alimentacion
	a.db.Find(&alimentaciones)
	return alimentaciones
}

func (a *AlmacenSQLite) BuscarAlimentacionPorID(id int) (alimentacionModels.Alimentacion, bool) {
	var alimentacion alimentacionModels.Alimentacion
	if err := a.db.First(&a, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return alimentacionModels.Alimentacion{}, false
	}
	return alimentacion, true
}

func (a *AlmacenSQLite) CrearAlimentacion(alimentacion alimentacionModels.Alimentacion) alimentacionModels.Alimentacion {
	a.db.Create(&a) // GORM rellena el ID autogenerado en &a
	return alimentacion
}

func (a *AlmacenSQLite) ActualizarAlimentacion(id int, datos alimentacionModels.Alimentacion) (alimentacionModels.Alimentacion, bool) {
	var existente alimentacionModels.Alimentacion
	if err := a.db.First(&existente, id).Error; err != nil {
		return alimentacionModels.Alimentacion{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarAlimentacion(id int) bool {
	res := a.db.Delete(&alimentacionModels.Alimentacion{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// MENU DIARIO
// =========================================================

func (a *AlmacenSQLite) ListarMenuDiarios() []alimentacionModels.MenuDiario {
	var menuDiarios []alimentacionModels.MenuDiario
	a.db.Find(&menuDiarios)
	return menuDiarios
}

func (a *AlmacenSQLite) BuscarMenuDiarioPorID(id int) (alimentacionModels.MenuDiario, bool) {
	var m alimentacionModels.MenuDiario
	if err := a.db.First(&m, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return alimentacionModels.MenuDiario{}, false
	}
	return m, true
}

func (a *AlmacenSQLite) CrearMenuDiario(m alimentacionModels.MenuDiario) alimentacionModels.MenuDiario {
	a.db.Create(&m) // GORM rellena el ID autogenerado en &m
	return m
}

func (a *AlmacenSQLite) ActualizarMenuDiario(id int, datos alimentacionModels.MenuDiario) (alimentacionModels.MenuDiario, bool) {
	var existente alimentacionModels.MenuDiario
	if err := a.db.First(&existente, id).Error; err != nil {
		return alimentacionModels.MenuDiario{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarMenuDiario(id int) bool {
	res := a.db.Delete(&alimentacionModels.MenuDiario{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// PLATO
// =========================================================

func (a *AlmacenSQLite) ListarPlatos() []alimentacionModels.Plato {
	var platos []alimentacionModels.Plato
	a.db.Find(&platos)
	return platos
}

func (a *AlmacenSQLite) BuscarPlatoPorID(id int) (alimentacionModels.Plato, bool) {
	var p alimentacionModels.Plato
	if err := a.db.First(&p, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return alimentacionModels.Plato{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearPlato(p alimentacionModels.Plato) alimentacionModels.Plato {
	a.db.Create(&p) // GORM rellena el ID autogenerado en &p
	return p
}

func (a *AlmacenSQLite) ActualizarPlato(id int, datos alimentacionModels.Plato) (alimentacionModels.Plato, bool) {
	var existente alimentacionModels.Plato
	if err := a.db.First(&existente, id).Error; err != nil {
		return alimentacionModels.Plato{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarPlato(id int) bool {
	res := a.db.Delete(&alimentacionModels.Plato{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// RESENA
// =========================================================

func (a *AlmacenSQLite) ListarResenas() []alimentacionModels.Resena {
	var resenas []alimentacionModels.Resena
	a.db.Find(&resenas)
	return resenas
}

func (a *AlmacenSQLite) BuscarResenaPorID(id int) (alimentacionModels.Resena, bool) {
	var r alimentacionModels.Resena
	if err := a.db.First(&r, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return alimentacionModels.Resena{}, false
	}
	return r, true
}

func (a *AlmacenSQLite) CrearResena(r alimentacionModels.Resena) alimentacionModels.Resena {
	a.db.Create(&r) // GORM rellena el ID autogenerado en &r
	return r
}

func (a *AlmacenSQLite) ActualizarResena(id int, datos alimentacionModels.Resena) (alimentacionModels.Resena, bool) {
	var existente alimentacionModels.Resena
	if err := a.db.First(&existente, id).Error; err != nil {
		return alimentacionModels.Resena{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarResena(id int) bool {
	res := a.db.Delete(&alimentacionModels.Resena{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// COOPERATIVA
// =========================================================

func (a *AlmacenSQLite) ListarCooperativas() []transporteModels.Cooperativa {
	var cooperativas []transporteModels.Cooperativa
	a.db.Find(&cooperativas)
	return cooperativas
}

func (a *AlmacenSQLite) BuscarCooperativaPorID(id uint) (transporteModels.Cooperativa, bool) {
	var c transporteModels.Cooperativa
	if err := a.db.First(&c, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return transporteModels.Cooperativa{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCooperativa(c transporteModels.Cooperativa) transporteModels.Cooperativa {
	a.db.Create(&c) // GORM rellena el ID autogenerado en &c
	return c
}

func (a *AlmacenSQLite) ActualizarCooperativa(id uint, datos transporteModels.Cooperativa) (transporteModels.Cooperativa, bool) {
	var existente transporteModels.Cooperativa
	if err := a.db.First(&existente, id).Error; err != nil {
		return transporteModels.Cooperativa{}, false
	}
	datos.ID = uint(id)
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarCooperativa(id uint) bool {
	res := a.db.Delete(&transporteModels.Cooperativa{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// RUTAS DE TRANSPORTE
// =========================================================

func (a *AlmacenSQLite) ListarRutas() []transporteModels.RutaTransporte {
	var rutas []transporteModels.RutaTransporte
	a.db.Find(&rutas)
	return rutas
}

func (a *AlmacenSQLite) BuscarRutaPorID(id uint) (transporteModels.RutaTransporte, bool) {
	var r transporteModels.RutaTransporte
	if err := a.db.First(&r, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return transporteModels.RutaTransporte{}, false
	}
	return r, true
}

func (a *AlmacenSQLite) CrearRuta(r transporteModels.RutaTransporte) transporteModels.RutaTransporte {
	a.db.Create(&r) // GORM rellena el ID autogenerado en &r
	return r
}

func (a *AlmacenSQLite) ActualizarRuta(id uint, datos transporteModels.RutaTransporte) (transporteModels.RutaTransporte, bool) {
	var existente transporteModels.RutaTransporte
	if err := a.db.First(&existente, id).Error; err != nil {
		return transporteModels.RutaTransporte{}, false
	}
	datos.ID = uint(id)
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarRuta(id uint) bool {
	res := a.db.Delete(&transporteModels.RutaTransporte{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// PARADAS BUS
// =========================================================

func (a *AlmacenSQLite) ListarParadas() []transporteModels.ParadaBus {
	var paradas []transporteModels.ParadaBus
	a.db.Find(&paradas)
	return paradas
}

func (a *AlmacenSQLite) BuscarParadaPorID(id uint) (transporteModels.ParadaBus, bool) {
	var p transporteModels.ParadaBus
	if err := a.db.First(&p, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return transporteModels.ParadaBus{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParada(p transporteModels.ParadaBus) transporteModels.ParadaBus {
	a.db.Create(&p) // GORM rellena el ID autogenerado en &p
	return p
}

func (a *AlmacenSQLite) ActualizarParada(id uint, datos transporteModels.ParadaBus) (transporteModels.ParadaBus, bool) {
	var existente transporteModels.ParadaBus
	if err := a.db.First(&existente, id).Error; err != nil {
		return transporteModels.ParadaBus{}, false
	}
	datos.ID = uint(id)
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarParada(id uint) bool {
	res := a.db.Delete(&transporteModels.ParadaBus{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarSiVacio() {
	// =========
	// USUARIO ADMIN — siempre se verifica, independiente de otros datos
	// =========
	var u int64
	a.db.Model(&usuarioModels.Usuario{}).Where("rol = ?", "admin").Count(&u)
	if u == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		result := a.db.Save(&usuarioModels.Usuario{
			Nombre:       "Administrador",
			Email:        "administrador@forneos.com",
			PasswordHash: string(hash),
			Rol:          "admin",
		})
		log.Printf("Admin creado: %v, error: %v", result.RowsAffected, result.Error)
	}
	var n int64

	// Revisamos una tabla principal
	a.db.Model(&viviendaModels.Sector{}).Count(&n)

	if n > 0 {
		return
	}

	// =========
	// VIVIENDA
	// =========

	sectores := []viviendaModels.Sector{
		{SectorID: 1, Nombre: "Centro"},
		{SectorID: 2, Nombre: "Universidad"},
	}

	a.db.Create(&sectores)

	viviendas := []viviendaModels.Vivienda{
		{
			ViviendaID: 1, Titulo: "Casa Central", SectorID: 1,
		},
	}

	a.db.Create(&viviendas)

	// =========
	// TRANSPORTE
	// =========

	cooperativas := []transporteModels.Cooperativa{
		{
			ID:     1,
			Nombre: "Ruta Universitaria",
		},
	}

	a.db.Create(&cooperativas)

	// =========
	// ALIMENTACION
	// =========

	alimentaciones := []alimentacionModels.Alimentacion{
		{
			ID:          1,
			NombreLocal: "Pan",
		},
	}

	a.db.Create(&alimentaciones)
}

var _ Almacen = (*AlmacenSQLite)(nil)
