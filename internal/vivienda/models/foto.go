package models

type Foto struct {
	FotoID     int    `json:"foto_id"`
	URL        string `json:"url"`
	ViviendaID int    `json:"vivienda_id"`
}
