package service

import (
"for-neos-api/internal/alimentacion/models"
"for-neos-api/internal/storage"
"strings"
)

type MenuDiarioService struct {
repo storage.MenuDiarioRepository
}

func NuevoMenuDiarioService(repo storage.MenuDiarioRepository) *MenuDiarioService {
return &MenuDiarioService{repo: repo}
}

func (s *MenuDiarioService) Listar() []models.MenuDiario {
return s.repo.ListarMenuDiarios()
}

func (s *MenuDiarioService) Obtener(id int) (models.MenuDiario, error) {
m, ok := s.repo.BuscarMenuDiarioPorID(id)
if !ok {
return models.MenuDiario{}, ErrNoEncontrado
}
return m, nil
}

func (s *MenuDiarioService) Crear(m models.MenuDiario) (models.MenuDiario, error) {
if err := validarMenuDiario(m); err != nil {
return models.MenuDiario{}, err
}
return s.repo.CrearMenuDiario(m), nil
}

func (s *MenuDiarioService) Actualizar(id int, m models.MenuDiario) (models.MenuDiario, error) {
if err := validarMenuDiario(m); err != nil {
return models.MenuDiario{}, err
}


actualizado, ok := s.repo.ActualizarMenuDiario(id, m)
if !ok {
	return models.MenuDiario{}, ErrNoEncontrado
}

return actualizado, nil


}

func (s *MenuDiarioService) Borrar(id int) error {
if !s.repo.BorrarMenuDiario(id) {
return ErrNoEncontrado
}
return nil
}

func validarMenuDiario(m models.MenuDiario) error {

if strings.TrimSpace(m.Fecha) == "" {
	return ErrFechaVacia
}

if m.AlimentacionID == 0 {
	return ErrNoEncontrado
}

return nil

}