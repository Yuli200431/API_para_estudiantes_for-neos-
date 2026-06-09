package models

type Resena struct {
	ID             int     `json:"id"`
	Comentario     string  `json:"comentario"`
	Calificacion   int     `json:"calificacion"` // 1–5 estrellas
	AlimentacionID int     `json:"alimentacion_id"`
}