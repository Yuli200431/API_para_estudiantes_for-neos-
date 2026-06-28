package storage

import (
	"context"
	"database/sql"

	alimentacionModels "for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/storage/sqlcdb"
	transporteModels "for-neos-api/internal/transporte/models"
	viviendaModels "for-neos-api/internal/vivienda/models"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{
		q: sqlcdb.New(db),
	}
}

// =========================================================
// MAPEO sqlc <-> dominio (la "capa anticorrupción")
// =========================================================

// =========================================================
// VIVIENDA
// =========================================================
func aViviendaDominio(v sqlcdb.Vivienda) viviendaModels.Vivienda {
	garantia := v.Garantia
	luz := v.Luz
	agua := v.Agua
	amueblado := v.Amueblado
	internet := v.Internet
	mascotas := v.Mascotas
	bañoPrivado := v.BanoPrivado

	return viviendaModels.Vivienda{
		ViviendaID:         int(v.ViviendaID),
		Titulo:             v.Titulo,
		TipoVivienda:       v.TipoVivienda,
		Precio:             v.Precio,
		Garantia:           &garantia,
		PrecioGarantia:     v.PrecioGarantia,
		Direccion:          v.Direccion,
		Luz:                &luz,
		Agua:               &agua,
		Amueblado:          &amueblado,
		Internet:           &internet,
		BañoPrivado:        &bañoPrivado,
		NumeroHabitaciones: int(v.NumeroHabitaciones),
		Mascotas:           &mascotas,
		GeneroPreferido:    v.GeneroPreferido,
		ReglasConvivencia:  v.ReglasConvivencia,
		Telefono:           v.Telefono,
		Email:              v.Email,
		Estado:             v.Estado,
		Comentario:         v.Comentario,
		SectorID:           int(v.SectorID),
		ProveedorID:        int(v.ProveedorID),
	}
}

func aFotoDominio(f sqlcdb.Foto) viviendaModels.Foto {
	return viviendaModels.Foto{
		FotoID:     int(f.FotoID),
		URL:        f.Url,
		ViviendaID: int(f.ViviendaID),
	}
}

func aAplicarViviendaDominio(av sqlcdb.AplicarVivienda) viviendaModels.AplicarVivienda {
	return viviendaModels.AplicarVivienda{
		AplicarViviendaID: int(av.AplicarViviendaID),
		EstudianteID:      int(av.EstudianteID),
		ViviendaID:        int(av.ViviendaID),
		Estado:            av.Estado,
	}
}

func aSectorDominio(s sqlcdb.Sector) viviendaModels.Sector {
	return viviendaModels.Sector{
		SectorID: int(s.SectorID),
		Nombre:   s.Nombre,
	}
}

// =========================================================
// ALIMENTACION
// =========================================================

func aAlimentacionDominio(a sqlcdb.Alimentacion) alimentacionModels.Alimentacion {
	return alimentacionModels.Alimentacion{
		ID:              int(a.ID),
		NombreLocal:     a.NombreLocal,
		Descripcion:     a.Descripcion,
		Ubicacion:       a.Ubicacion,
		Direccion:       a.Direccion,
		HorarioApertura: a.HorarioApertura,
		HorarioCierre:   a.HorarioCierre,
		Telefono:        a.Telefono,
		TipoComida:      a.TipoComida,
		PrecioPromedio:  a.PrecioPromedio,
		ProviderID:      int(a.ProviderID),
	}
}

func aMenuDiarioDominio(m sqlcdb.MenuDiario) alimentacionModels.MenuDiario {
	return alimentacionModels.MenuDiario{
		ID:             int(m.ID),
		Fecha:          m.Fecha,
		AlimentacionID: int(m.AlimentacionID),
	}
}

func aPlatoDominio(p sqlcdb.Plato) alimentacionModels.Plato {
	return alimentacionModels.Plato{
		ID:           int(p.ID),
		NombrePlato:  p.NombrePlato,
		Descripcion:  p.Descripcion,
		Categoria:    p.Categoria,
		Precio:       p.Precio,
		MenuDiarioID: int(p.MenuDiarioID),
	}
}

func aResenaDominio(r sqlcdb.Resena) alimentacionModels.Resena {
	return alimentacionModels.Resena{
		ID:             int(r.ID),
		Comentario:     r.Comentario,
		Calificacion:   int(r.Calificacion),
		AlimentacionID: int(r.AlimentacionID),
	}
}

// =========================================================
// TRANSPORTE
// =========================================================

func aCooperativaDominio(c sqlcdb.Cooperativa) transporteModels.Cooperativa {
	return transporteModels.Cooperativa{
		ID:          uint(c.ID),
		Nombre:      c.Nombre,
		Telefono:    c.Telefono,
		Descripcion: c.Descripcion,
	}
}

func aParadaDominio(p sqlcdb.ParadasBus) transporteModels.ParadaBus {
	return transporteModels.ParadaBus{
		ID:           uint(p.ID),
		NombreParada: p.NombreParada,
		Direccion:    p.Direccion,
		Descripcion:  p.Descripcion,
	}
}

func aRutaDominio(r sqlcdb.RutaTransporte) transporteModels.RutaTransporte {
	return transporteModels.RutaTransporte{
		ID:              uint(r.ID),
		NombreLinea:     r.NombreLinea,
		FrecuenciaAprox: r.FrecuenciaAprox,
		Precio:          r.Precio,
		DescripcionRuta: r.DescripcionRuta,
		CooperativaID:   uint(r.CooperativaID),
		SectorOrigenID:  uint(r.SectorOrigenID),
		SectorDestinoID: uint(r.SectorDestinoID),
		ParadaBusID:     uint(r.ParadaBusID),
	}
}

// =========================================================
// VIVIENDA
// =========================================================

//Vivienda

func (a *AlmacenSQLC) ListarViviendas() []viviendaModels.Vivienda {
	filas, err := a.q.ListarViviendas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]viviendaModels.Vivienda, 0, len(filas))
	for _, f := range filas {
		out = append(out, aViviendaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarViviendaPorID(id int) (viviendaModels.Vivienda, bool) {
	f, err := a.q.BuscarViviendaPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return viviendaModels.Vivienda{}, false
	}
	return aViviendaDominio(f), true
}

func (a *AlmacenSQLC) CrearVivienda(v viviendaModels.Vivienda) viviendaModels.Vivienda {
	f, err := a.q.CrearVivienda(context.Background(), sqlcdb.CrearViviendaParams{
		Titulo:             v.Titulo,
		TipoVivienda:       v.TipoVivienda,
		Precio:             v.Precio,
		Garantia:           v.Garantia != nil,
		PrecioGarantia:     v.PrecioGarantia,
		Direccion:          v.Direccion,
		Luz:                v.Luz != nil,
		Agua:               v.Agua != nil,
		Amueblado:          v.Amueblado != nil,
		Internet:           v.Internet != nil,
		BanoPrivado:        v.BañoPrivado != nil,
		NumeroHabitaciones: int64(v.NumeroHabitaciones),
		Mascotas:           v.Mascotas != nil,
		GeneroPreferido:    v.GeneroPreferido,
		ReglasConvivencia:  v.ReglasConvivencia,
		Telefono:           v.Telefono,
		Email:              v.Email,
		Estado:             v.Estado,
		Comentario:         v.Comentario,
		SectorID:           int64(v.SectorID),
		ProveedorID:        int64(v.ProveedorID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return viviendaModels.Vivienda{}
	}
	return aViviendaDominio(f)
}

func (a *AlmacenSQLC) ActualizarVivienda(id int, datos viviendaModels.Vivienda) (viviendaModels.Vivienda, bool) {
	f, err := a.q.ActualizarVivienda(context.Background(), sqlcdb.ActualizarViviendaParams{
		Titulo:             datos.Titulo,
		TipoVivienda:       datos.TipoVivienda,
		Precio:             datos.Precio,
		Garantia:           datos.Garantia != nil,
		PrecioGarantia:     datos.PrecioGarantia,
		Direccion:          datos.Direccion,
		Luz:                datos.Luz != nil,
		Agua:               datos.Agua != nil,
		Amueblado:          datos.Amueblado != nil,
		Internet:           datos.Internet != nil,
		BanoPrivado:        datos.BañoPrivado != nil,
		NumeroHabitaciones: int64(datos.NumeroHabitaciones),
		Mascotas:           datos.Mascotas != nil,
		GeneroPreferido:    datos.GeneroPreferido,
		ReglasConvivencia:  datos.ReglasConvivencia,
		Telefono:           datos.Telefono,
		Email:              datos.Email,
		Estado:             datos.Estado,
		Comentario:         datos.Comentario,
		SectorID:           int64(datos.SectorID),
		ProveedorID:        int64(datos.ProveedorID),
		ViviendaID:         int64(id),
	})
	if err != nil {
		return viviendaModels.Vivienda{}, false
	}
	return aViviendaDominio(f), true
}

func (a *AlmacenSQLC) BorrarVivienda(id int) bool {
	filas, err := a.q.EliminarVivienda(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

//Foto

func (a *AlmacenSQLC) ListarFotos() []viviendaModels.Foto {
	filas, err := a.q.ListarFotos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]viviendaModels.Foto, 0, len(filas))
	for _, f := range filas {
		out = append(out, aFotoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarFotoPorID(id int) (viviendaModels.Foto, bool) {
	f, err := a.q.BuscarFotoPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return viviendaModels.Foto{}, false
	}
	return aFotoDominio(f), true
}

func (a *AlmacenSQLC) CrearFoto(foto viviendaModels.Foto) viviendaModels.Foto {
	nuevaFoto, err := a.q.CrearFoto(context.Background(), sqlcdb.CrearFotoParams{
		Url:        foto.URL,
		ViviendaID: int64(foto.ViviendaID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return viviendaModels.Foto{}
	}
	return aFotoDominio(nuevaFoto)
}

func (a *AlmacenSQLC) ActualizarFoto(id int, datos viviendaModels.Foto) (viviendaModels.Foto, bool) {
	f, err := a.q.ActualizarFoto(context.Background(), sqlcdb.ActualizarFotoParams{
		Url:        datos.URL,
		ViviendaID: int64(datos.ViviendaID),
		FotoID:     int64(id),
	})
	if err != nil {
		return viviendaModels.Foto{}, false
	}
	return aFotoDominio(f), true
}

func (a *AlmacenSQLC) BorrarFoto(id int) bool {
	filas, err := a.q.BorrarFoto(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// AplicarVivienda

func (a *AlmacenSQLC) ListarAplicarViviendas() []viviendaModels.AplicarVivienda {
	filas, err := a.q.ListarAplicarViviendas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]viviendaModels.AplicarVivienda, 0, len(filas))
	for _, f := range filas {
		out = append(out, aAplicarViviendaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarAplicarViviendaPorID(id int) (viviendaModels.AplicarVivienda, bool) {
	f, err := a.q.BuscarAplicarViviendaPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return viviendaModels.AplicarVivienda{}, false
	}
	return aAplicarViviendaDominio(f), true
}

func (a *AlmacenSQLC) CrearAplicarVivienda(av viviendaModels.AplicarVivienda) viviendaModels.AplicarVivienda {
	f, err := a.q.CrearAplicarVivienda(context.Background(), sqlcdb.CrearAplicarViviendaParams{
		EstudianteID: int64(av.EstudianteID),
		ViviendaID:   int64(av.ViviendaID),
		Estado:       av.Estado,
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return viviendaModels.AplicarVivienda{}
	}
	return aAplicarViviendaDominio(f)
}

func (a *AlmacenSQLC) ActualizarAplicarVivienda(id int, datos viviendaModels.AplicarVivienda) (viviendaModels.AplicarVivienda, bool) {
	f, err := a.q.ActualizarAplicarVivienda(context.Background(), sqlcdb.ActualizarAplicarViviendaParams{
		EstudianteID:      int64(datos.EstudianteID),
		ViviendaID:        int64(datos.ViviendaID),
		Estado:            datos.Estado,
		AplicarViviendaID: int64(id),
	})
	if err != nil {
		return viviendaModels.AplicarVivienda{}, false
	}
	return aAplicarViviendaDominio(f), true
}

func (a *AlmacenSQLC) BorrarAplicarVivienda(id int) bool {
	filas, err := a.q.BorrarAplicarVivienda(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// Sector

func (a *AlmacenSQLC) ListarSectores() []viviendaModels.Sector {
	filas, err := a.q.ListarSectores(context.Background())
	if err != nil {
		return nil
	}
	out := make([]viviendaModels.Sector, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSectorDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSectorPorID(id int) (viviendaModels.Sector, bool) {
	f, err := a.q.BuscarSectorPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return viviendaModels.Sector{}, false
	}
	return aSectorDominio(f), true
}

func (a *AlmacenSQLC) CrearSector(s viviendaModels.Sector) viviendaModels.Sector {
	f, err := a.q.CrearSector(context.Background(), s.Nombre)
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return viviendaModels.Sector{}
	}
	return aSectorDominio(f)
}

func (a *AlmacenSQLC) ActualizarSector(id int, datos viviendaModels.Sector) (viviendaModels.Sector, bool) {
	f, err := a.q.ActualizarSector(context.Background(), sqlcdb.ActualizarSectorParams{
		Nombre:   datos.Nombre,
		SectorID: int64(id),
	})
	if err != nil {
		return viviendaModels.Sector{}, false
	}
	return aSectorDominio(f), true
}

func (a *AlmacenSQLC) BorrarSector(id int) bool {
	filas, err := a.q.BorrarSector(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// ALIMENTACION
// =========================================================

//Alimentacion

func (a *AlmacenSQLC) ListarAlimentaciones() []alimentacionModels.Alimentacion {
	filas, err := a.q.ListarAlimentaciones(context.Background())
	if err != nil {
		return nil
	}
	out := make([]alimentacionModels.Alimentacion, 0, len(filas))
	for _, f := range filas {
		out = append(out, aAlimentacionDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarAlimentacionPorID(id int) (alimentacionModels.Alimentacion, bool) {
	f, err := a.q.BuscarAlimentacionPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return alimentacionModels.Alimentacion{}, false
	}
	return aAlimentacionDominio(f), true
}

func (a *AlmacenSQLC) CrearAlimentacion(alimentacion alimentacionModels.Alimentacion) alimentacionModels.Alimentacion {
	f, err := a.q.CrearAlimentacion(context.Background(), sqlcdb.CrearAlimentacionParams{
		NombreLocal:     alimentacion.NombreLocal,
		Descripcion:     alimentacion.Descripcion,
		Ubicacion:       alimentacion.Ubicacion,
		Direccion:       alimentacion.Direccion,
		HorarioApertura: alimentacion.HorarioApertura,
		HorarioCierre:   alimentacion.HorarioCierre,
		Telefono:        alimentacion.Telefono,
		TipoComida:      alimentacion.TipoComida,
		PrecioPromedio:  alimentacion.PrecioPromedio,
		ProviderID:      int64(alimentacion.ProviderID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return alimentacionModels.Alimentacion{}
	}
	return aAlimentacionDominio(f)
}

func (a *AlmacenSQLC) ActualizarAlimentacion(id int, datos alimentacionModels.Alimentacion) (alimentacionModels.Alimentacion, bool) {
	f, err := a.q.ActualizarAlimentacion(context.Background(), sqlcdb.ActualizarAlimentacionParams{
		NombreLocal:     datos.NombreLocal,
		Descripcion:     datos.Descripcion,
		Ubicacion:       datos.Ubicacion,
		Direccion:       datos.Direccion,
		HorarioApertura: datos.HorarioApertura,
		HorarioCierre:   datos.HorarioCierre,
		Telefono:        datos.Telefono,
		TipoComida:      datos.TipoComida,
		PrecioPromedio:  datos.PrecioPromedio,
		ProviderID:      int64(datos.ProviderID),
		ID:              int64(id),
	})
	if err != nil {
		return alimentacionModels.Alimentacion{}, false
	}
	return aAlimentacionDominio(f), true
}

func (a *AlmacenSQLC) BorrarAlimentacion(id int) bool {
	filas, err := a.q.BorrarAlimentacion(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// MenuDiario

func (a *AlmacenSQLC) ListarMenuDiarios() []alimentacionModels.MenuDiario {
	filas, err := a.q.ListarMenuDiarios(context.Background())
	if err != nil {
		return nil
	}
	out := make([]alimentacionModels.MenuDiario, 0, len(filas))
	for _, f := range filas {
		out = append(out, aMenuDiarioDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarMenuDiarioPorID(id int) (alimentacionModels.MenuDiario, bool) {
	f, err := a.q.BuscarMenuDiarioPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return alimentacionModels.MenuDiario{}, false
	}
	return aMenuDiarioDominio(f), true
}

func (a *AlmacenSQLC) CrearMenuDiario(m alimentacionModels.MenuDiario) alimentacionModels.MenuDiario {
	f, err := a.q.CrearMenuDiario(context.Background(), sqlcdb.CrearMenuDiarioParams{
		Fecha:          m.Fecha,
		AlimentacionID: int64(m.AlimentacionID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return alimentacionModels.MenuDiario{}
	}
	return aMenuDiarioDominio(f)
}

func (a *AlmacenSQLC) ActualizarMenuDiario(id int, datos alimentacionModels.MenuDiario) (alimentacionModels.MenuDiario, bool) {
	f, err := a.q.ActualizarMenuDiario(context.Background(), sqlcdb.ActualizarMenuDiarioParams{
		Fecha:          datos.Fecha,
		AlimentacionID: int64(datos.AlimentacionID),
		ID:             int64(id),
	})
	if err != nil {
		return alimentacionModels.MenuDiario{}, false
	}
	return aMenuDiarioDominio(f), true
}

func (a *AlmacenSQLC) BorrarMenuDiario(id int) bool {
	filas, err := a.q.BorrarMenuDiario(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// Plato

func (a *AlmacenSQLC) ListarPlatos() []alimentacionModels.Plato {
	filas, err := a.q.ListarPlatos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]alimentacionModels.Plato, 0, len(filas))
	for _, f := range filas {
		out = append(out, aPlatoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarPlatoPorID(id int) (alimentacionModels.Plato, bool) {
	f, err := a.q.BuscarPlatoPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return alimentacionModels.Plato{}, false
	}
	return aPlatoDominio(f), true
}

func (a *AlmacenSQLC) CrearPlato(p alimentacionModels.Plato) alimentacionModels.Plato {
	f, err := a.q.CrearPlato(context.Background(), sqlcdb.CrearPlatoParams{
		NombrePlato:  p.NombrePlato,
		Descripcion:  p.Descripcion,
		Categoria:    p.Categoria,
		Precio:       p.Precio,
		MenuDiarioID: int64(p.MenuDiarioID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return alimentacionModels.Plato{}
	}
	return aPlatoDominio(f)
}

func (a *AlmacenSQLC) ActualizarPlato(id int, datos alimentacionModels.Plato) (alimentacionModels.Plato, bool) {
	f, err := a.q.ActualizarPlato(context.Background(), sqlcdb.ActualizarPlatoParams{
		NombrePlato:  datos.NombrePlato,
		Descripcion:  datos.Descripcion,
		Categoria:    datos.Categoria,
		Precio:       datos.Precio,
		MenuDiarioID: int64(datos.MenuDiarioID),
		ID:           int64(id),
	})
	if err != nil {
		return alimentacionModels.Plato{}, false
	}
	return aPlatoDominio(f), true
}

func (a *AlmacenSQLC) BorrarPlato(id int) bool {
	filas, err := a.q.BorrarPlato(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// Resena

func (a *AlmacenSQLC) ListarResenas() []alimentacionModels.Resena {
	filas, err := a.q.ListarResenas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]alimentacionModels.Resena, 0, len(filas))
	for _, f := range filas {
		out = append(out, aResenaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarResenaPorID(id int) (alimentacionModels.Resena, bool) {
	f, err := a.q.BuscarResenaPorID(context.Background(), int64(id))
	if err != nil {
		return alimentacionModels.Resena{}, false
	}
	return aResenaDominio(f), true
}

func (a *AlmacenSQLC) CrearResena(r alimentacionModels.Resena) alimentacionModels.Resena {
	f, err := a.q.CrearResena(context.Background(), sqlcdb.CrearResenaParams{
		Comentario:     r.Comentario,
		Calificacion:   int64(r.Calificacion),
		AlimentacionID: int64(r.AlimentacionID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return alimentacionModels.Resena{}
	}
	return aResenaDominio(f)
}

func (a *AlmacenSQLC) ActualizarResena(id int, datos alimentacionModels.Resena) (alimentacionModels.Resena, bool) {
	f, err := a.q.ActualizarResena(context.Background(), sqlcdb.ActualizarResenaParams{
		Comentario:     datos.Comentario,
		Calificacion:   int64(datos.Calificacion),
		AlimentacionID: int64(datos.AlimentacionID),
		ID:             int64(id),
	})
	if err != nil {
		return alimentacionModels.Resena{}, false
	}
	return aResenaDominio(f), true
}

func (a *AlmacenSQLC) BorrarResena(id int) bool {
	filas, err := a.q.BorrarResena(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// Transporte
// =========================================================

//Cooperativa

func (a *AlmacenSQLC) ListarCooperativas() []transporteModels.Cooperativa {
	filas, err := a.q.ListarCooperativas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]transporteModels.Cooperativa, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCooperativaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCooperativaPorID(id uint) (transporteModels.Cooperativa, bool) {
	f, err := a.q.BuscarCooperativaPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return transporteModels.Cooperativa{}, false
	}
	return aCooperativaDominio(f), true
}

func (a *AlmacenSQLC) CrearCooperativa(c transporteModels.Cooperativa) transporteModels.Cooperativa {
	f, err := a.q.CrearCooperativa(context.Background(), sqlcdb.CrearCooperativaParams{
		Nombre:      c.Nombre,
		Telefono:    c.Telefono,
		Descripcion: c.Descripcion,
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return transporteModels.Cooperativa{}
	}
	return aCooperativaDominio(f)
}

func (a *AlmacenSQLC) ActualizarCooperativa(id uint, datos transporteModels.Cooperativa) (transporteModels.Cooperativa, bool) {
	f, err := a.q.ActualizarCooperativa(context.Background(), sqlcdb.ActualizarCooperativaParams{
		Nombre:      datos.Nombre,
		Telefono:    datos.Telefono,
		Descripcion: datos.Descripcion,
		ID:          int64(id),
	})
	if err != nil {
		return transporteModels.Cooperativa{}, false
	}
	return aCooperativaDominio(f), true
}

func (a *AlmacenSQLC) BorrarCooperativa(id uint) bool {
	filas, err := a.q.BorrarCooperativa(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// rutaTransporte

func (a *AlmacenSQLC) ListarRutas() []transporteModels.RutaTransporte {
	filas, err := a.q.ListarRutas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]transporteModels.RutaTransporte, 0, len(filas))
	for _, f := range filas {
		out = append(out, aRutaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarRutaPorID(id uint) (transporteModels.RutaTransporte, bool) {
	f, err := a.q.BuscarRutaPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return transporteModels.RutaTransporte{}, false
	}
	return aRutaDominio(f), true
}

func (a *AlmacenSQLC) CrearRuta(r transporteModels.RutaTransporte) transporteModels.RutaTransporte {
	f, err := a.q.CrearRuta(context.Background(), sqlcdb.CrearRutaParams{
		NombreLinea:     r.NombreLinea,
		FrecuenciaAprox: r.FrecuenciaAprox,
		Precio:          r.Precio,
		DescripcionRuta: r.DescripcionRuta,
		CooperativaID:   int64(r.CooperativaID),
		SectorOrigenID:  int64(r.SectorOrigenID),
		SectorDestinoID: int64(r.SectorDestinoID),
		ParadaBusID:     int64(r.ParadaBusID),
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return transporteModels.RutaTransporte{}
	}
	return aRutaDominio(f)
}

func (a *AlmacenSQLC) ActualizarRuta(id uint, datos transporteModels.RutaTransporte) (transporteModels.RutaTransporte, bool) {
	f, err := a.q.ActualizarRuta(context.Background(), sqlcdb.ActualizarRutaParams{
		NombreLinea:     datos.NombreLinea,
		FrecuenciaAprox: datos.FrecuenciaAprox,
		Precio:          datos.Precio,
		DescripcionRuta: datos.DescripcionRuta,
		CooperativaID:   int64(datos.CooperativaID),
		SectorOrigenID:  int64(datos.SectorOrigenID),
		SectorDestinoID: int64(datos.SectorDestinoID),
		ParadaBusID:     int64(datos.ParadaBusID),
		ID:              int64(id),
	})
	if err != nil {
		return transporteModels.RutaTransporte{}, false
	}
	return aRutaDominio(f), true
}

func (a *AlmacenSQLC) BorrarRuta(id uint) bool {
	filas, err := a.q.BorrarRuta(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// ParadaBus

func (a *AlmacenSQLC) ListarParadas() []transporteModels.ParadaBus {
	filas, err := a.q.ListarParadas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]transporteModels.ParadaBus, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParadaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParadaPorID(id uint) (transporteModels.ParadaBus, bool) {
	f, err := a.q.BuscarParadaPorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return transporteModels.ParadaBus{}, false
	}
	return aParadaDominio(f), true
}

func (a *AlmacenSQLC) CrearParada(p transporteModels.ParadaBus) transporteModels.ParadaBus {
	f, err := a.q.CrearParada(context.Background(), sqlcdb.CrearParadaParams{
		NombreParada: p.NombreParada,
		Direccion:    p.Direccion,
		Descripcion:  p.Descripcion,
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value. Ver nota en la guía docente.
		return transporteModels.ParadaBus{}
	}
	return aParadaDominio(f)
}

func (a *AlmacenSQLC) ActualizarParada(id uint, datos transporteModels.ParadaBus) (transporteModels.ParadaBus, bool) {
	f, err := a.q.ActualizarParada(context.Background(), sqlcdb.ActualizarParadaParams{
		NombreParada: datos.NombreParada,
		Direccion:    datos.Direccion,
		Descripcion:  datos.Descripcion,
		ID:           int64(id),
	})
	if err != nil {
		return transporteModels.ParadaBus{}, false
	}
	return aParadaDominio(f), true
}

func (a *AlmacenSQLC) BorrarParada(id uint) bool {
	filas, err := a.q.BorrarParada(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)
