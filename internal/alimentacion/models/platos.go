package models

type Plato struct {
	ID           int     `json:"id"`
	NombrePlato  string  `json:"nombre_plato"`
	Descripcion  string  `json:"descripcion"`
	Categoria    string  `json:"categoria"`
	Precio       float64 `json:"precio"`
	MenuDiarioID int     `json:"menu_diario_id"`
}