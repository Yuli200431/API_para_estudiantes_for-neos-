package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/vivienda/models"
	"strings"
)

type SectorService struct {
	repo storage.SectorRepository
}

func NuevaSectorService(repo storage.SectorRepository) *SectorService {
	return &SectorService{repo: repo}
}

func (s *SectorService) Listar() []models.Sector {
	return s.repo.ListarSectores()
}

func (s *SectorService) Obtener(id int) (models.Sector, error) {
	e, ok := s.repo.BuscarSectorPorID(id)
	if !ok {
		return models.Sector{}, ErrNoEncontrado
	}
	return e, nil
}

func (s *SectorService) Crear(e models.Sector) (models.Sector, error) {
	if err := validacionSector(e); err != nil {
		return models.Sector{}, err
	}
	return s.repo.CrearSector(e), nil
}

func (s *SectorService) Actualizar(id int, e models.Sector) (models.Sector, error) {
	if err := validacionSector(e); err != nil {
		return models.Sector{}, err
	}
	actualizado, ok := s.repo.ActualizarSector(id, e)
	if !ok {
		return models.Sector{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *SectorService) Borrar(id int) error {
	if !s.repo.BorrarSector(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionSector(e models.Sector) error {
	if strings.TrimSpace(e.Nombre) == "" {
		return ErrNombreVacio
	}
	return nil
}
