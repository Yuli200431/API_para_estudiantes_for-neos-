package storage

import (
	"errors"

	"for-neos-api/internal/usuario/models"
)

type UsuarioStorage struct {
	usuarios []models.Usuario

	nextID int
}

func NewUsuarioStorage() *UsuarioStorage {

	return &UsuarioStorage{
		usuarios: []models.Usuario{},

		nextID: 1,
	}
}

func (s *UsuarioStorage) CrearUsuario(u models.Usuario) (models.Usuario, error) {

	u.ID =
		s.nextID

	s.nextID++

	s.usuarios =
		append(
			s.usuarios,
			u,
		)

	return u, nil
}

func (s *UsuarioStorage) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {

	for _, u := range s.usuarios {

		if u.Email ==
			email {

			return u,
				true
		}
	}

	return models.Usuario{},
		false
}

var _ = errors.New
