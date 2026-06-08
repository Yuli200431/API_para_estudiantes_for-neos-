package models

type Transporte struct {
	ID                 uint    `json:"id"`
	NombreLinea        string  `json:"nombre_linea"`
	Cooperativa        string  `json:"cooperativa"`
	SectorOrigen       string  `json:"sector_origen"`
	SectorDestino      string  `json:"sector_destino"`
	SectoresRecorridos string  `json:"sectores_recorridos"`
	FrecuenciaAprox    string  `json:"frecuencia_aprox"`
	Precio             float64 `json:"precio"`
	DescripcionRuta    string  `json:"descripcion_ruta"`
	ProviderID         uint    `json:"provider_id"`
}
