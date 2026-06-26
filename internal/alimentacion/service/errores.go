package service

import "errors"

var (
	ErrNoEncontrado     = errors.New("no encontrado")
	ErrNombreVacio      = errors.New("nombre vacio")
	ErrDescripcionVacio = errors.New("descripcion vacia")
	ErrPrecioVacio      = errors.New("precio vacio")
	ErrEstadoVacio      = errors.New("estado vacio")
	ErrFechaVacia       = errors.New("fecha vacia")
	ErrNombrePlatoVacio = errors.New("nombre del plato vacio")
	ErrComentarioVacio  = errors.New("comentario vacio")
)