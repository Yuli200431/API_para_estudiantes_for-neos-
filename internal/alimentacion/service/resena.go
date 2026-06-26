package service

import (
	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/storage"
	"strings"
)

type ResenaService struct {
	repo storage.ResenaRepository
}

func NuevaResenaService(repo storage.ResenaRepository) *ResenaService {
	return &ResenaService{repo: repo}
}

func (s *ResenaService) Listar() []models.Resena {
	return s.repo.ListarResenas()
}

func (s *ResenaService) Obtener(id int) (models.Resena, error) {
	r, ok := s.repo.BuscarResenaPorID(id)
	if !ok {
		return models.Resena{}, ErrNoEncontrado
	}
	return r, nil
}

func (s *ResenaService) Crear(r models.Resena) (models.Resena, error) {
	if err := validarResena(r); err != nil {
		return models.Resena{}, err
	}
	return s.repo.CrearResena(r), nil
}

func (s *ResenaService) Actualizar(id int, r models.Resena) (models.Resena, error) {
	if err := validarResena(r); err != nil {
		return models.Resena{}, err
	}
	actualizado, ok := s.repo.ActualizarResena(id, r)
	if !ok {
		return models.Resena{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ResenaService) Borrar(id int) error {
	if !s.repo.BorrarResena(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarResena(r models.Resena) error {
	if strings.TrimSpace(r.Comentario) == "" {
		return ErrComentarioVacio
	}
	return nil
}