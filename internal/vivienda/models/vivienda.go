package models

type Vivienda struct {
	ViviendaID         int     `json:"vivienda_id"`
	Titulo             string  `json:"titulo"`
	TipoVivienda       string  `json:"tipo_vivienda"`
	Precio             int     `json:"precio"`
	Garantia           bool    `json:"garantia"`
	PrecioGarantia     float64 `json:"precio_garantia"`
	Direccion          string  `json:"direccion"`
	Luz                bool    `json:"luz"`
	Agua               bool    `json:"agua"`
	Amueblado          bool    `json:"amueblado"`
	Internet           bool    `json:"internet"`
	BañoPrivado        bool    `json:"baño_privado"`
	NumeroHabitaciones int     `json:"numero_habitaciones"`
	Mascotas           bool    `json:"mascotas"`
	GeneroPreferido    string  `json:"genero_preferido"`
	ReglasConvivencia  string  `json:"reglas_convivencia"`
	Telefono           string  `json:"telefono"`
	Email              string  `json:"email"`
	Estado             string  `json:"estado"`
	Comentario         string  `json:"comentario"`

	SectorID    int    `json:"sector_id"`
	Fotos       []Foto `json:"fotos"`
	ProveedorID int    `json:"proveedor_id"`
}
