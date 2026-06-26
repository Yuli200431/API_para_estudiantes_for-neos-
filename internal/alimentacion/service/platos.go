package service

import (
	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/storage"
	"strings"
)

type PlatoService struct {
	repo storage.PlatoRepository
}

func NuevoPlatoService(repo storage.PlatoRepository) *PlatoService {
	return &PlatoService{repo: repo}
}

func (s *PlatoService) Listar() []models.Plato {
	return s.repo.ListarPlatos()
}

func (s *PlatoService) Obtener(id int) (models.Plato, error) {
	p, ok := s.repo.BuscarPlatoPorID(id)
	if !ok {
		return models.Plato{}, ErrNoEncontrado
	}
	return p, nil
}

func (s *PlatoService) Crear(p models.Plato) (models.Plato, error) {
	if err := validarPlato(p); err != nil {
		return models.Plato{}, err
	}
	return s.repo.CrearPlato(p), nil
}

func (s *PlatoService) Actualizar(id int, p models.Plato) (models.Plato, error) {
	if err := validarPlato(p); err != nil {
		return models.Plato{}, err
	}
	actualizado, ok := s.repo.ActualizarPlato(id, p)
	if !ok {
		return models.Plato{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *PlatoService) Borrar(id int) error {
	if !s.repo.BorrarPlato(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarPlato(p models.Plato) error {
	if strings.TrimSpace(p.NombrePlato) == "" {
		return ErrNombrePlatoVacio
	}
	return nil
}