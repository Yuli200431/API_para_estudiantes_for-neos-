package storage

import (
	"API-Foraneos/internal/alimentacion/models"
)

var Alimentaciones []models.Alimentacion
var NextID = 1

type Memoria struct{
	alimentaciones []models.Alimentacion
	nextAlimentacionID int
    
	mu sync.Mutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		alimentaciones: []models.Alimentacion{},
		nextAlimentacionID: 1,
	}
} 
