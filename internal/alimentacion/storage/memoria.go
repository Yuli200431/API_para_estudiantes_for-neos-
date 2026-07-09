package storage

import (
	"sync"
	"for-neos-api/internal/alimentacion/models"
	
)

type Memoria struct{
	alimentaciones []models.Alimentacion
	nextAlimentacionID int

	menuDiarios []models.MenuDiario
	nextMenuDiarioID int

	platos []models.Plato
	nextPlatoID int

	resenas []models.Resena
	nextResenaID int
    
	mu sync.Mutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		alimentaciones: []models.Alimentacion{},
		nextAlimentacionID: 1,

		menuDiarios: []models.MenuDiario{},
		nextMenuDiarioID: 1,

		platos: []models.Plato{},
		nextPlatoID: 1,

		resenas: []models.Resena{},
		nextResenaID: 1,
	
	}
} 

func (m *Memoria) SeedAlimentaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.alimentaciones = []models.Alimentacion{
		{
			ID:              1,
			NombreLocal:     "Comedor ULEAM",
			Descripcion:     "Comida económica para estudiantes",
			Ubicacion:       "Dentro de la ULEAM",
			Direccion:       "Campus principal",
			HorarioApertura: "07:00",
			HorarioCierre:   "17:00",
			Telefono:        "0939039152",
			TipoComida:      "Casera",
			PrecioPromedio:  2.50,
			ProviderID:      1,
		},
		{
			ID:              2,
			NombreLocal:     "La Sazón Manabita",
			Descripcion:     "Almuerzos y platos típicos",
			Ubicacion:       "Manta Centro",
			Direccion:       "Av. 4 y Calle 12",
			HorarioApertura: "08:00",
			HorarioCierre:   "18:00",
			Telefono:        "0939037890",
			TipoComida:      "Manabita",
			PrecioPromedio:  3.00,
			ProviderID:      2,
		},
	}


	m.nextAlimentacionID = 3
}

//Listar 
func (m *Memoria) ListarAlimentaciones() []models.Alimentacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	copia := make([]models.Alimentacion, len(m.alimentaciones))
	copy(copia, m.alimentaciones)
	return copia
}
//buscar por ID
func (m *Memoria) BuscarAlimentacionPorID(id int) (models.Alimentacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, a := range m.alimentaciones {
		if a.ID == id {
			return a, true
		}	
	}
	return models.Alimentacion{}, false
}
//crearAlimentacion
func (m *Memoria) CrearAlimentacion(alimentacion models.Alimentacion) models.Alimentacion {
	m.mu.Lock()
	defer m.mu.Unlock()

	alimentacion.ID = m.nextAlimentacionID
	m.nextAlimentacionID++
	m.alimentaciones = append(m.alimentaciones, alimentacion)
	return alimentacion
}
//actualizar
func (m *Memoria) ActualizarAlimentacion(id int, alimentacion models.Alimentacion) (models.Alimentacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.alimentaciones {
		if a.ID == id {
			alimentacion.ID = id
			m.alimentaciones[i] = alimentacion
			return alimentacion, true
		}	
	}
	return models.Alimentacion{}, false
}
//borrar
func (m *Memoria) BorrarAlimentacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, a := range m.alimentaciones {
		if a.ID == id {
			m.alimentaciones = append(m.alimentaciones[:i], m.alimentaciones[i+1:]...)
			return true
		}
	}
	return false
}

/// MENU DIARIOS ///

// Seed de MenuDiarios
func (m *Memoria) SeedMenuDiarios() {
	m.mu.Lock()
	defer m.mu.Unlock()	
	m.menuDiarios = []models.MenuDiario{
		{
			ID:             1,	
			Fecha:          "2024-10-01",
			AlimentacionID: 1,
		},
		{
			ID:             2,
			Fecha:          "2024-10-02",
			AlimentacionID: 2,
		},
	}
	m.nextMenuDiarioID = 3
}

// ListarMenuDiarios devuelve todos los menús diarios
func (m *Memoria) ListarMenuDiarios() []models.MenuDiario {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.MenuDiario, len(m.menuDiarios))
	copy(copia, m.menuDiarios)
	return copia
}

// BuscarMenuDiarioPorID devuelve un menú diario por su ID
func (m *Memoria) BuscarMenuDiarioPorID(id int) (models.MenuDiario, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()	
	for _, md := range m.menuDiarios {
		if md.ID == id {
			return md, true
		}
	}
	return models.MenuDiario{}, false
}

// CrearMenuDiario crea un nuevo menú diario
func (m *Memoria) CrearMenuDiario(menuDiario models.MenuDiario) models.MenuDiario {
	m.mu.Lock()
	defer m.mu.Unlock()
	menuDiario.ID = m.nextMenuDiarioID
	m.nextMenuDiarioID++
	m.menuDiarios = append(m.menuDiarios, menuDiario)
	return menuDiario
}

// ActualizarMenuDiario actualiza un menú diario por su ID
func (m *Memoria) ActualizarMenuDiario(id int, menuDiario models.MenuDiario) (models.MenuDiario, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, md := range m.menuDiarios {
		if md.ID == id {
			menuDiario.ID = id
			m.menuDiarios[i] = menuDiario
			return menuDiario, true
		}
	}
	return models.MenuDiario{}, false
}

// BorrarMenuDiario borra un menú diario por su ID
func (m *Memoria) BorrarMenuDiario(id int) bool{
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, md := range m.menuDiarios {
		if md.ID == id {
			m.menuDiarios = append(m.menuDiarios[:i], m.menuDiarios[i+1:]...)
			return true
		}
	}
	return false
}

//metodos de platos

// Seed de Platos
func (m *Memoria) SeedPlatos() {
	m.mu.Lock()
	defer m.mu.Unlock()	
	m.platos = []models.Plato{
		{
			ID:           1,
			NombrePlato:  "Arroz con Pollo",
			Descripcion:  "Arroz amarillo con pollo guisado",
			Categoria:    "Plato Principal",
			Precio:       2.50,
			MenuDiarioID: 1,
		},
		{
			ID:           2,
			NombrePlato:  "Ensalada de Frutas",
			Descripcion:  "Mezcla de frutas frescas",
			Categoria:    "Postre",
			Precio:       1.50,
			MenuDiarioID: 1,
		},
		{
			ID:           3,	
			NombrePlato:  "Sopa de Menestra",
			Descripcion:  "Sopa de lentejas con verduras",
			Categoria:    "Entrada",
			Precio:       1.00,	
			MenuDiarioID: 2,
		},
	}
	m.nextPlatoID = 4
}

// ListarPlatos devuelve todos los platos
func (m *Memoria) ListarPlatos() []models.Plato {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Plato, len(m.platos))
	copy(copia, m.platos)
	return copia
}

// BuscarPlatoPorID devuelve un plato por su ID
func (m *Memoria) BuscarPlatoPorID(id int) (models.Plato, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.platos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Plato{}, false
}

// CrearPlato crea un nuevo plato
func (m *Memoria) CrearPlato(plato models.Plato) models.Plato {
	m.mu.Lock()
	defer m.mu.Unlock()

	plato.ID = m.nextPlatoID
	m.nextPlatoID++
	m.platos = append(m.platos, plato)
	return plato
}

// ActualizarPlato actualiza un plato por su ID
func (m *Memoria) ActualizarPlato(id int, plato models.Plato) (models.Plato, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.platos {
		if p.ID == id {
			plato.ID = id
			m.platos[i] = plato
			return plato, true
		}
	}
	return models.Plato{}, false
}

// BorrarPlato borra un plato por su ID
func (m *Memoria) BorrarPlato(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.platos {
		if p.ID == id {
			m.platos = append(m.platos[:i], m.platos[i+1:]...)
			return true
		}
	}
	return false
}				

//metodos de resenas

// Seed de Resenas
func (m *Memoria) SeedResenas() {
	m.mu.Lock()
	defer m.mu.Unlock()	
	m.resenas = []models.Resena{
		{
			ID:             1,
			Comentario:     "Muy buena",
			Calificacion:   5,
			AlimentacionID: 1,
		},
		{
			ID:             2,
			Comentario:     "Buena",
			Calificacion:   4,
			AlimentacionID: 2,
		},
		{
			ID:             3,
			Comentario:     "Mala",
			Calificacion:   3,
			AlimentacionID: 3,
		},
	}
	m.nextResenaID = 4
}

// ListarResenas devuelve todos los resenas
func (m *Memoria) ListarResenas() []models.Resena {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Resena, len(m.resenas))
	copy(copia, m.resenas)
	return copia
}

// BuscarResenaPorID devuelve un resena por su ID
func (m *Memoria) BuscarResenaPorID(id int) (models.Resena, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, r := range m.resenas {
		if r.ID == id {
			return r, true
		}
	}			
	return models.Resena{}, false
	
}

// CrearResena crea un nuevo resena
func (m *Memoria) CrearResena(resena models.Resena) models.Resena {
	m.mu.Lock()
	defer m.mu.Unlock()

	resena.ID = m.nextResenaID
	m.nextResenaID++
	m.resenas = append(m.resenas, resena)
	return resena
}

// ActualizarResena actualiza un resena por su ID
func (m *Memoria) ActualizarResena(id int, resena models.Resena) (models.Resena, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, r := range m.resenas {
		if r.ID == id {
			resena.ID = id
			m.resenas[i] = resena
			return resena, true
		}
	}
	return models.Resena{}, false
}

// BorrarResena borra un resena por su ID
func (m *Memoria) BorrarResena(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, r := range m.resenas {
		if r.ID == id {
			m.resenas = append(m.resenas[:i], m.resenas[i+1:]...)
			return true
		}
	}
	return false
	
}	
