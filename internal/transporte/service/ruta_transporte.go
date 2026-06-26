package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/transporte/models"
	"strings"
)

type RutaTransporteService struct {
	repo storage.RutaTransporteRepository
}

func NuevaRutaService(repo storage.RutaTransporteRepository) *RutaTransporteService{
	return &RutaTransporteService{repo:repo}
}

func (s *RutaTransporteService) ListarRuta() []models.RutaTransporte {
	return s.repo.ListarRutas()
}

func (s *RutaTransporteService) ObtenerRutas(id int) (models.RutaTransporte, error) {
	a, ok := s.repo.ObtenerRutaPorID(uint(id))
	if !ok {
		return models.RutaTransporte{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *RutaTransporteService) CrearRuta(a models.RutaTransporte) (models.RutaTransporte, error) {
	if err := validacionRutaTransporte(a); err != nil {
		return models.RutaTransporte{}, err
	}
	return s.repo.AgregarRuta(a), nil
}

func (s *RutaTransporteService) ActualizarRuta(id int, a models.RutaTransporte) (models.RutaTransporte, error) {
	if err := validacionRutaTransporte(a); err != nil {
		return models.RutaTransporte{}, err
	}
	actualizado, ok := s.repo.ActualizarRuta(uint(id), a)
	if !ok {
		return models.RutaTransporte{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *RutaTransporteService) BorrarRuta(id int) error {
	if !s.repo.EliminarRuta(uint(id)) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionRutaTransporte(a models.RutaTransporte) error {
	if strings.TrimSpace(a.NombreLinea) == "" {
		return ErrEstadoVacio
	}
	return nil
}