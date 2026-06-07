package models

type AplicarVivienda struct {
	AplicarViviendaID int    `json:"aplicar_vivienda_id"`
	EstudianteID      int    `json:"estudiante_id"`
	ViviendaID        int    `json:"vivienda_id"`
	Estado            string `json:"estado"`
}
