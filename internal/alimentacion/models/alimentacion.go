package models

type Alimentacion struct {
	ID              int     `json:"id" gorm:"primaryKey"`
	NombreLocal     string  `json:"nombre_local" gorm:"not null"`
	Descripcion     string  `json:"descripcion"`
	Ubicacion       string  `json:"ubicacion"`
	Direccion       string  `json:"direccion"`
	HorarioApertura string  `json:"horario_apertura"`
	HorarioCierre   string  `json:"horario_cierre"`
	Telefono        string  `json:"telefono"`
	TipoComida      string  `json:"tipo_comida"`
	PrecioPromedio  float64 `json:"precio_promedio"`
	ProviderID      int     `json:"provider_id"`
}