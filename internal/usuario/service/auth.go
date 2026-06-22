package service

import (
	"for-neos-api/internal/storage"
	"for-neos-api/internal/usuario/models"

	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretoJWT = []byte("palabra_bastante_secreta")

// El tiempo que queremos que duren el Token
var duracionToken = time.Hour * 24

// Claims: Todo lo que quiero que este Token contenga, la información que yo quiero que este en este Token
type Claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo storage.UsuarioStorage
}

func NuevoAuthService(repo storage.UsuarioStorage) *AuthService {
	return &AuthService{repo: repo}

}
func (s *AuthService) Registrar(email, password string) (models.Usuario, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || strings.TrimSpace(password) == "" {
		return models.Usuario{}, ErrNombreVacio
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
			//Que se va a usar para identificar al token
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	// Creamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretoJWT)

}

func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return secretoJWT, nil
	})
	if err != nil || token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	return claims.UsuarioID, nil
}
