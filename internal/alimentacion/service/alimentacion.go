package service

import (
	"for-neos-api/internal/alimentacion/models"
	"for-neos-api/internal/storage"
	"strings"
)

type AlimentacionService struct {
	repo storage.AlimentacionRepository
}

func NuevaAlimentacionService(repo storage.AlimentacionRepository) *AlimentacionService {
	return &AlimentacionService{repo: repo}
}

func (s *AlimentacionService) Listar() []models.Alimentacion {
	return s.repo.ListarAlimentaciones()
}

func (s *AlimentacionService) Obtener(id int) (models.Alimentacion, error) {
	a, ok := s.repo.BuscarAlimentacionPorID(id)
	if !ok {
		return models.Alimentacion{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *AlimentacionService) Crear(a models.Alimentacion) (models.Alimentacion, error) {
	if err := validacionAlimentacion(a); err != nil {
		return models.Alimentacion{}, err
	}
	return s.repo.CrearAlimentacion(a), nil
}

func (s *AlimentacionService) Actualizar(id int, a models.Alimentacion) (models.Alimentacion, error) {
	if err := validacionAlimentacion(a); err != nil {
		return models.Alimentacion{}, err
	}

	actualizado, ok := s.repo.ActualizarAlimentacion(id, a)
	if !ok {
		return models.Alimentacion{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *AlimentacionService) Borrar(id int) error {
	if !s.repo.BorrarAlimentacion(id) {
		return ErrNoEncontrado
	}
	return nil
}


func validacionAlimentacion(a models.Alimentacion) error {

	if strings.TrimSpace(a.NombreLocal) == "" {
		return ErrNombreVacio
	}

	if strings.TrimSpace(a.Descripcion) == "" {
		return ErrDescripcionVacio
	}

	if strings.TrimSpace(a.Ubicacion) == "" {
		return ErrUbicacionVacia
	}

	if strings.TrimSpace(a.Direccion) == "" {
		return ErrDireccionVacia
	}

	if strings.TrimSpace(a.Telefono) == "" {
		return ErrTelefonoVacio
	}

	if strings.TrimSpace(a.TipoComida) == "" {
		return ErrTipoComidaVacia
	}

	if a.PrecioPromedio <= 0 {
		return ErrPrecioVacio
	}

	if a.ProviderID == 0 {
		return ErrNoEncontrado
	}

	return nil
}