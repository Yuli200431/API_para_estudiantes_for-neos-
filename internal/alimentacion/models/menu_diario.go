package models

type MenuDiario struct {
	ID             int    `json:"id"`
	Fecha          string `json:"fecha"`
	AlimentacionID int    `json:"alimentacion_id"`
}