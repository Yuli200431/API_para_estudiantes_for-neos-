package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/transporte/models"
	"strings"
)

type CooperativaService struct {
	repo storage.CooperativaRepository
}

func NuevaCooperaticaService(repo storage.CooperativaRepository) *CooperativaService {
	return &CooperativaService{repo: repo}
}

func (s *CooperativaService) Listar() []models.Cooperativa {
	return s.repo.ListarCooperativas()
}

func (s *CooperativaService) Obtener(id int) (models.Cooperativa, error) {
	a, ok := s.repo.BuscarCooperativaPorID(uint(id))
	if !ok {
		return models.Cooperativa{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *CooperativaService) Crear(a models.Cooperativa) (models.Cooperativa, error) {
	if err := validacionCooperativa(a); err != nil {
		return models.Cooperativa{}, err
	}
	return s.repo.CrearCooperativa(a), nil
}

func (s *CooperativaService) Actualizar(id int, a models.Cooperativa) (models.Cooperativa, error) {
	if err := validacionCooperativa(a); err != nil {
		return models.Cooperativa{}, err
	}
	actualizado, ok := s.repo.ActualizarCooperativa(uint(id), a)
	if !ok {
		return models.Cooperativa{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *CooperativaService) Borrar(id int) error {
	if !s.repo.BorrarCooperativa(uint(id)) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionCooperativa(a models.Cooperativa) error {
	if strings.TrimSpace(a.Nombre) == "" {
		return ErrEstadoVacio
	}
	if strings.TrimSpace(a.Telefono) == "" {
		return ErrEstadoVacio
	}
	return nil
}
