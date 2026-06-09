package storage

import (
	"sync"

	"for-neos-api/internal/transporte/models"
)

type MemoriaTransporte struct {
	rutas  []models.RutaTransporte
	nextID uint

	cooperativas      []models.Cooperativa
	nextCooperativaID uint

	sectores     []models.Sector
	nextSectorID uint

	mu sync.Mutex
}

func NuevaMemoriaTransporte() *MemoriaTransporte {
	return &MemoriaTransporte{
		rutas:             []models.RutaTransporte{},
		nextID:            1,
		cooperativas:      []models.Cooperativa{},
		nextCooperativaID: 1,
		sectores:          []models.Sector{},
		nextSectorID:      1,
		mu:                sync.Mutex{},
	}
}

//RUTAS DE TRANSPORTE

// SeedTransportes agrega rutas de transporte de ejemplo a la memoria para pruebas iniciales
func (m *MemoriaTransporte) SeedTransportes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []models.RutaTransporte{
		{ID: 1, NombreLinea: "Línea 1", CooperativaID: 1, SectorOrigenID: 1, SectorDestinoID: 2, FrecuenciaAprox: "Cada 10 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 1 con el Sector 2 pasando por el Sector 3", ProviderID: 1},
		{ID: 2, NombreLinea: "Línea 2", CooperativaID: 2, SectorOrigenID: 2, SectorDestinoID: 3, FrecuenciaAprox: "Cada 15 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 2 con el Sector 3 pasando por el Sector 4", ProviderID: 2},
		{ID: 3, NombreLinea: "Línea 3", CooperativaID: 3, SectorOrigenID: 1, SectorDestinoID: 4, FrecuenciaAprox: "Cada 12 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 1 con el Sector 4 pasando por el Sector 5", ProviderID: 3},
	}
	m.nextID = uint(4)
}

// Listar todas las rutas
func (m *MemoriaTransporte) ListarRutas() []models.RutaTransporte {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.RutaTransporte, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

// Obtener una ruta por ID
func (m *MemoriaTransporte) ObtenerRutaPorID(id uint) (models.RutaTransporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.rutas {
		if p.ID == id {
			return p, true
		}
	}
	return models.RutaTransporte{}, false
}

func (m *MemoriaTransporte) AgregarRuta(ruta models.RutaTransporte) models.RutaTransporte {
	m.mu.Lock()
	defer m.mu.Unlock()

	ruta.ID = m.nextID
	m.nextID++
	m.rutas = append(m.rutas, ruta)
	return ruta
}

func (m *MemoriaTransporte) ActualizarRuta(id uint, rutaActualizada models.RutaTransporte) (models.RutaTransporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.rutas {
		if p.ID == id {
			rutaActualizada.ID = id
			m.rutas[i] = rutaActualizada
			return rutaActualizada, true
		}
	}
	return models.RutaTransporte{}, false
}

func (m *MemoriaTransporte) EliminarRuta(id uint) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.rutas {
		if p.ID == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return true
		}
	}
	return false
}
