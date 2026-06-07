package storage

import (
	"sync"

	"for-neos-api/internal/vivienda/models"
)

// Base de datos en memoria. Definir lo que se va a guardar.
type Memoria struct {
	viviendas      []models.Vivienda
	nextViviendaID int

	fotos      []models.Foto
	nextFotoID int

	sectores []models.Sector

	aplicarviviendas       []models.AplicarVivienda
	nextAplicarViviendasID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		viviendas:      []models.Vivienda{},
		nextViviendaID: 1,

		fotos:      []models.Foto{},
		nextFotoID: 1,

		sectores: []models.Sector{},

		aplicarviviendas:       []models.AplicarVivienda{},
		nextAplicarViviendasID: 1,
	}
}

// VIVIENDAS

func (m *Memoria) SeedViviendas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.viviendas = []models.Vivienda{
		{ViviendaID: 1, Titulo: "Casa Amarilla", TipoVivienda: "Casa", Precio: 120, Estado: "Disponible", SectorID: 1, ProveedorID: 1},
		{ViviendaID: 2, Titulo: "Suite de Lujo", TipoVivienda: "Departamento", Precio: 200, Estado: "Ocupado", SectorID: 1, ProveedorID: 1},
		{ViviendaID: 3, Titulo: "Departamento en el Centro", TipoVivienda: "Departamento", Precio: 150, Estado: "Disponible", SectorID: 2, ProveedorID: 2},
	}
	m.nextViviendaID = 4
}

// ListarViviendas devuelve todas las viviendas en memoria.
func (m *Memoria) ListarViviendas() []models.Vivienda {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Vivienda, len(m.viviendas))
	copy(copia, m.viviendas)
	return copia
}

// BuscarViviendaPorID devuelve la vivienda con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarViviendaPorID(id int) (models.Vivienda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, v := range m.viviendas {
		if v.ViviendaID == id {
			return v, true
		}
	}
	return models.Vivienda{}, false
}

// CrearVivienda agrega una vivienda nueva y devuelve la vivienda con ID asignado.
func (m *Memoria) CrearVivienda(v models.Vivienda) models.Vivienda {
	m.mu.Lock()
	defer m.mu.Unlock()

	v.ViviendaID = m.nextViviendaID
	m.nextViviendaID++
	m.viviendas = append(m.viviendas, v)
	return v
}

// ActualizarVivienda reemplaza la vivienda con el ID dado.
func (m *Memoria) ActualizarVivienda(id int, datos models.Vivienda) (models.Vivienda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.viviendas {
		if c.ViviendaID == id {
			datos.ViviendaID = id
			m.viviendas[i] = datos
			return datos, true
		}
	}
	return models.Vivienda{}, false
}

// BorrarVivienda elimina la vivienda con el ID dado.
func (m *Memoria) BorrarVivienda(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.viviendas {
		if p.ViviendaID == id {
			m.viviendas = append(m.viviendas[:i], m.viviendas[i+1:]...)
			return true
		}
	}
	return false
}

// SECTORES

// SeedSectores carga categorías iniciales que coinciden con CategoriaID de los productos pre-cargados.
func (m *Memoria) SeedSectores() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sectores = []models.Sector{
		{SectorID: 1, Nombre: "Santa Martha"},
		{SectorID: 2, Nombre: "La Epoca"},
		{SectorID: 3, Nombre: "Los Electricos"},
	}
}

// FOTOS

// SeedFotos carga fotos iniciales que coinciden con ViviendaID de las viviendas pre-cargadas.
func (m *Memoria) SeedFotos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.fotos = []models.Foto{
		{FotoID: 1, URL: "https://example.com/foto1.jpg", ViviendaID: 1},
		{FotoID: 2, URL: "https://example.com/foto2.jpg", ViviendaID: 2},
		{FotoID: 3, URL: "https://example.com/foto3.jpg", ViviendaID: 3},
	}
	m.nextFotoID = 4
}

// ListarFotos devuelve todas las fotos en memoria.
func (m *Memoria) ListarFotos() []models.Foto {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Foto, len(m.fotos))
	copy(copia, m.fotos)
	return copia
}

// BuscarFotoPorID devuelve la foto con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarFotoPorID(id int) (models.Foto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, f := range m.fotos {
		if f.FotoID == id {
			return f, true
		}
	}
	return models.Foto{}, false
}

// CrearFoto agrega una foto nueva y devuelve la foto con ID asignado.
func (m *Memoria) CrearFoto(f models.Foto) models.Foto {
	m.mu.Lock()
	defer m.mu.Unlock()

	f.FotoID = m.nextFotoID
	m.nextFotoID++
	m.fotos = append(m.fotos, f)
	return f
}

// ActualizarFoto reemplaza la foto con el ID dado.
func (m *Memoria) ActualizarFoto(id int, datos models.Foto) (models.Foto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, f := range m.fotos {
		if f.FotoID == id {
			datos.FotoID = id
			m.fotos[i] = datos
			return datos, true
		}
	}
	return models.Foto{}, false
}

// BorrarFoto elimina la foto con el ID dado.
func (m *Memoria) BorrarFoto(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, f := range m.fotos {
		if f.FotoID == id {
			m.fotos = append(m.fotos[:i], m.fotos[i+1:]...)
			return true
		}
	}
	return false
}
