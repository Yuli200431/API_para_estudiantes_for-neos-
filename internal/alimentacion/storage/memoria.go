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
    
	mu sync.Mutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		alimentaciones: []models.Alimentacion{},
		nextAlimentacionID: 1,

		menuDiarios: []models.MenuDiario{},
		nextMenuDiarioID: 1,
	
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