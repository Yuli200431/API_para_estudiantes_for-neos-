package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/usuario/models"

	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretoPorDefecto  = "cafeteria-uleam-secreto-solo-dev"
	duracionPorDefecto = 24 * time.Hour
)

// Claims: Todo lo que quiero que este Token contenga, la información que yo quiero que este en este Token
type Claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo     storage.UserRepository
	secreto  []byte
	duracion time.Duration
}

// AuthOption configura un AuthService en su construccion (patron Options).
type AuthOption func(*AuthService)

// WithSecreto inyecta la clave de firma del JWT (desde config/.env en produccion).
// Si recibe un secreto vacio, conserva el valor por defecto.
func WithSecreto(secreto []byte) AuthOption {
	return func(a *AuthService) {
		if len(secreto) > 0 {
			a.secreto = secreto
		}
	}
}

// WithDuracionToken inyecta la validez del token. Si recibe <= 0, conserva el default.
func WithDuracionToken(d time.Duration) AuthOption {
	return func(a *AuthService) {
		if d > 0 {
			a.duracion = d
		}
	}
}

func NuevoAuthService(repo storage.UserRepository, opts ...AuthOption) *AuthService {
	s := &AuthService{
		repo:     repo,
		secreto:  []byte(secretoPorDefecto),
		duracion: duracionPorDefecto,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func (s *AuthService) Registrar(nombre, email, password string) (models.Usuario, error) {
	nombre = strings.TrimSpace(nombre)
	email = strings.TrimSpace(strings.ToLower(email))
	if nombre == "" || strings.TrimSpace(password) == "" {
		return models.Usuario{}, ErrNombreVacio
	}
	if email == "" || strings.TrimSpace(password) == "" {
		return models.Usuario{}, ErrEmailVacio
	}
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}
	//Si todo funciona bien va a retornar el usuario Creado
	return s.repo.CrearUsuario(
		models.Usuario{
			Nombre:       nombre,
			Email:        email,
			PasswordHash: string(hash),
		})
}

func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	//Primero busco el usuario
	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	//Luego la condición de que el usuario exista
	if !existe {
		//Nada que hacer, retorno el error
		return "", ErrEmailEnUso
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}
	// Si todo funciona, retorno el token. El token su funcion es ser el que se le pasa a los endpoints
	return s.generarToken(u)
}

// Este Token es el que se le pasa a los endpoints. Generar el token
func (s *AuthService) generarToken(u models.Usuario) (string, error) {
	//Es lo que se va ocupar en el Token
	claims := Claims{
		//Datos del usuario, los que queramos.
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			//Hasta cuando el token va hacer valido
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.duracion)),
			//Que se va a usar para identificar al token
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	// Creamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secreto)
}

func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return s.secreto, nil
	})
	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	return claims.UsuarioID, nil
}
