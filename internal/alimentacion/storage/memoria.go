package storage

import (
	"sync"
	"API-Foraneos/internal/alimentacion/models"
)


type Memoria struct{
	alimentaciones []models.Alimentacion
	nextAlimentacionID int
    
	mu sync.Mutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		alimentaciones: []models.Alimentacion{},
		nextAlimentacionID: 1,
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