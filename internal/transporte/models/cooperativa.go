package models

type Cooperativa struct {
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Telefono    string `json:"telefono"`
	Descripcion string `json:"descripcion"`
}

