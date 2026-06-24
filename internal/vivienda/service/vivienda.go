package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"strings"
)

type ViviendaService struct {
	repo storage.ViviendaRepository
}

func NuevaViviendaService(repo storage.ViviendaRepository) *ViviendaService {
	return &ViviendaService{repo: repo}
}

func (s *ViviendaService) Listar() []models.Vivienda {
	return s.repo.ListarViviendas()
}

func (s *ViviendaService) Obtener(id int) (models.Vivienda, error) {
	v, ok := s.repo.BuscarViviendaPorID(id)
	if !ok {
		return models.Vivienda{}, ErrNoEncontrado
	}
	return v, nil
}

func (s *ViviendaService) Crear(v models.Vivienda) (models.Vivienda, error) {
	if err := validacionVivienda(v); err != nil {
		return models.Vivienda{}, err
	}
	return s.repo.CrearVivienda(v), nil
}

func (s *ViviendaService) Actualizar(id int, v models.Vivienda) (models.Vivienda, error) {
	if err := validacionVivienda(v); err != nil {
		return models.Vivienda{}, err
	}
	actualizado, ok := s.repo.ActualizarVivienda(id, v)
	if !ok {
		return models.Vivienda{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ViviendaService) Borrar(id int) error {
	if !s.repo.BorrarVivienda(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionVivienda(v models.Vivienda) error {
	if strings.TrimSpace(v.Titulo) == "" {
		return ErrTituloVacio
	}
	return nil
}
