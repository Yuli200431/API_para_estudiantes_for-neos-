package service

import "errors"

var (
	ErrEmailVacio            = errors.New("El campo Email es obligatorio")
	ErrPasswordVacio         = errors.New("La contraseña es obligatoria")
	ErrNombreVacio           = errors.New("El campo Nombre no puede estar vacío")
	ErrNoEncontrado          = errors.New("No se encontró el producto")
	ErrEmailEnUso            = errors.New("El campo Email ya está en uso")
	ErrCredencialesInvalidas = errors.New("Email o contraseña no son válidos")
)
