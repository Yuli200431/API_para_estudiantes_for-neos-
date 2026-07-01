package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"strings"
)

type AplicarViviendaService struct {
	repo storage.AplicarViviendaRepository
}

func NuevaAplicarViviendaService(repo storage.AplicarViviendaRepository) *AplicarViviendaService {
	return &AplicarViviendaService{repo: repo}
}

func (s *AplicarViviendaService) Listar() []models.AplicarVivienda {
	return s.repo.ListarAplicarViviendas()
}

func (s *AplicarViviendaService) Obtener(id int) (models.AplicarVivienda, error) {
	a, ok := s.repo.BuscarAplicarViviendaPorID(id)
	if !ok {
		return models.AplicarVivienda{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *AplicarViviendaService) Crear(a models.AplicarVivienda) (models.AplicarVivienda, error) {
	if err := validacionAplicarVivienda(a); err != nil {
		return models.AplicarVivienda{}, err
	}
	return s.repo.CrearAplicarVivienda(a), nil
}

func (s *AplicarViviendaService) Actualizar(id int, a models.AplicarVivienda) (models.AplicarVivienda, error) {
	if err := validacionAplicarVivienda(a); err != nil {
		return models.AplicarVivienda{}, err
	}
	actualizado, ok := s.repo.ActualizarAplicarVivienda(id, a)
	if !ok {
		return models.AplicarVivienda{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *AplicarViviendaService) Borrar(id int) error {
	if !s.repo.BorrarAplicarVivienda(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionAplicarVivienda(a models.AplicarVivienda) error {
	if strings.TrimSpace(a.Estado) == "" {
		return ErrEstadoVacio
	}
	if a.ViviendaID == 0 {
		return ErrNoEncontrado
	}
	return nil
}
