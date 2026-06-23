package models

type RutaTransporte struct {
	ID              uint    `json:"id"`
	NombreLinea     string  `json:"nombre_linea"`
	FrecuenciaAprox string  `json:"frecuencia_aprox"`
	Precio          float64 `json:"precio"`
	DescripcionRuta string  `json:"descripcion_ruta"`

	CooperativaID uint `json:"cooperativa_id"`

	SectorOrigenID  uint `json:"sector_origen_id"`
	SectorDestinoID uint `json:"sector_destino_id"`

	ParadaBusID uint `json:"parada_bus_id"`
}
