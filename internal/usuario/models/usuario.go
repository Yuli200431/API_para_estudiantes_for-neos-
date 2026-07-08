package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Nombre       string    `json:"nombre" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not full"`
	CreadoEn     time.Time `json:"creado_en"`
	Rol          string    `json:"rol" gorm:"not null;default:estudiante"`
}
