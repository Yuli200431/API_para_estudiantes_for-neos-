package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"strings"
)

type FotoService struct {
	repo storage.FotoRepository
}

func NuevaFotoService(repo storage.FotoRepository) *FotoService {
	return &FotoService{repo: repo}
}

func (s *FotoService) Listar() []models.Foto {
	return s.repo.ListarFotos()
}

func (s *FotoService) Obtener(id int) (models.Foto, error) {
	f, ok := s.repo.BuscarFotoPorID(id)
	if !ok {
		return models.Foto{}, ErrNoEncontrado
	}
	return f, nil
}

func (s *FotoService) Crear(f models.Foto) (models.Foto, error) {
	if err := validacionFoto(f); err != nil {
		return models.Foto{}, err
	}
	return s.repo.CrearFoto(f), nil
}

func (s *FotoService) Actualizar(id int, f models.Foto) (models.Foto, error) {
	if err := validacionFoto(f); err != nil {
		return models.Foto{}, err
	}
	actualizado, ok := s.repo.ActualizarFoto(id, f)
	if !ok {
		return models.Foto{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *FotoService) Borrar(id int) error {
	if !s.repo.BorrarFoto(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionFoto(f models.Foto) error {
	if strings.TrimSpace(f.URL) == "" {
		return ErrURLVacio
	}
	return nil
}
