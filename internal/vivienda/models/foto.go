package models

type Foto struct {
	FotoID     int    `gorm:"primaryKey" json:"foto_id"`
	URL        string `json:"url"`
	ViviendaID int    `json:"vivienda_id"`
}
