package storage

import (
	"sync"

	"for-neos-api/internal/transporte/models"
)

type MemoriaTransporte struct {
	rutas  []models.Transporte
	nextID uint
	mu     sync.Mutex
}

func NuevaMemoriaTransporte() *MemoriaTransporte {
	return &MemoriaTransporte{
		rutas:  []models.Transporte{},
		nextID: 1,
	}
}

func (m *MemoriaTransporte) SeedTransportes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []models.Transporte{
		{ID: 1, NombreLinea: "Línea 1", Cooperativa: "Cooperativa A", SectorOrigen: "Costa Azul", SectorDestino: "Universidad", SectoresRecorridos: "Sector 1, Sector 3, Sector 2", FrecuenciaAprox: "Cada 10 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 1 con el Sector 2 pasando por el Sector 3", ProviderID: 1},
		{ID: 2, NombreLinea: "Línea 2", Cooperativa: "Cooperativa B", SectorOrigen: "Terminal", SectorDestino: "Coliseo", SectoresRecorridos: "Sector 2, Sector 4, Sector 3", FrecuenciaAprox: "Cada 15 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 2 con el Sector 3 pasando por el Sector 4", ProviderID: 2},
		{ID: 3, NombreLinea: "Línea 3", Cooperativa: "Cooperativa C", SectorOrigen: "Cuba", SectorDestino: "Santa Martha", SectoresRecorridos: "Sector 1, Sector 5, Sector 4", FrecuenciaAprox: "Cada 12 minutos", Precio: 0.40, DescripcionRuta: "Ruta que conecta el Sector 1 con el Sector 4 pasando por el Sector 5", ProviderID: 3},
	}
}

// Listar todas las rutas
func (m *MemoriaTransporte) ListarRutas() []models.Transporte {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Transporte, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

func (m *MemoriaTransporte) ObtenerRutaPorID(id uint) (models.Transporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.rutas {
		if p.ID == id {
			return p, true
		}
	}
	return models.Transporte{}, false
}

func (m *MemoriaTransporte) AgregarRuta(ruta models.Transporte) models.Transporte {
	m.mu.Lock()
	defer m.mu.Unlock()

	ruta.ID = m.nextID
	m.nextID++
	m.rutas = append(m.rutas, ruta)
	return ruta
}
