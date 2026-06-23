package models

type ParadaBus struct {
	ID          uint   `json:"id"`
	NombreParada      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	Descripcion string `json:"descripcion"`
}