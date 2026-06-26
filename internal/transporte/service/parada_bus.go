package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/transporte/models"
	"strings"
)

type ParadaBusService struct {
	repo storage.ParadaBusRepository
}

func NuevaParadaBusService(repo storage.ParadaBusRepository) *ParadaBusService{
	return &ParadaBusService{repo:repo}
}

func (s *ParadaBusService) ListarParadas() []models.ParadaBus {
	return s.repo.ListarParadas()
}

func (s *ParadaBusService) ObtenerParadas(id int) (models.ParadaBus, error) {
	a, ok := s.repo.ObtenerParadaPorID(uint(id))
	if !ok {
		return models.ParadaBus{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *ParadaBusService) CrearParada(a models.ParadaBus) (models.ParadaBus, error) {
	if err := validacionParadaBus(a); err != nil {
		return models.ParadaBus{}, err
	}
	return s.repo.AgregarParada(a), nil
}

func (s *ParadaBusService) ActualizarParadas(id int, a models.ParadaBus) (models.ParadaBus, error) {
	if err := validacionParadaBus(a); err != nil {
		return models.ParadaBus{}, err
	}
	actualizado, ok := s.repo.ActualizarParada(uint(id), a)
	if !ok {
		return models.ParadaBus{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ParadaBusService) BorrarParadas(id int) error {
	if !s.repo.EliminarParada(uint(id)) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionParadaBus(a models.ParadaBus) error {
	if strings.TrimSpace(a.NombreParada) == "" {
		return ErrEstadoVacio
	}
	return nil
}