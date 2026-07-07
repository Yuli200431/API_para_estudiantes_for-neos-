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
	if strings.TrimSpace(v.TipoVivienda) == "" {
		return ErrTipoViviendaVacio
	}
	if v.Precio == 0 {
		return ErrPrecioVacio
	}
	if v.Garantia == nil {
		return ErrGarantiaVacia
	}
	if v.PrecioGarantia == 0 {
		return ErrPrecioGarantiaVacio
	}
	if v.Luz == nil {
		return ErrLuzVacia
	}
	if v.Agua == nil {
		return ErrAguaVacia
	}
	if v.Amueblado == nil {
		return ErrAmuebladoVacio
	}
	if v.Internet == nil {
		return ErrInternetVacio
	}
	if v.BañoPrivado == nil {
		return ErrBañoPrivadoVacio
	}
	if int(v.NumeroHabitaciones) == 0 {
		return ErrNumeroHabitacionesVacio
	}
	if v.Mascotas == nil {
		return ErrMascotasVacia
	}
	if strings.TrimSpace(v.GeneroPreferido) == "" {
		return ErrGeneroPreferidoVacio
	}
	if strings.TrimSpace(v.ReglasConvivencia) == "" {
		return ErrReglasConvivenciaVacio
	}
	if strings.TrimSpace(v.Telefono) == "" {
		return ErrTelefonoVacio
	}
	if strings.TrimSpace(v.Email) == "" {
		return ErrEmailVacio
	}
	if strings.TrimSpace(v.Estado) == "" {
		return ErrEstadoVacio
	}
	if strings.TrimSpace(v.Comentario) == "" {
		return ErrComentarioVacio
	}
	if v.SectorID == 0 {
		return ErrNoEncontrado
	}
	if v.ProveedorID == 0 {
		return ErrNoEncontrado
	}
	if v.Fotos == nil {
		return ErrNoEncontrado
	}
	for _, f := range v.Fotos {
		if f.FotoID == 0 {
			return ErrNoEncontrado
		}
	}
	return nil
}
