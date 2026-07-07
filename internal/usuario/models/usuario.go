package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Nombre       string    `json:"nombre" gorm:"not full;uniqueIndex"`
	Email        string    `json:"email" gorm:"not full;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not full"`
	CreadoEn     time.Time `json:"creado_en"`
	Rol          string    `gorm:"default:estudiante"`
}
