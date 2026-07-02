package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Puerto       string
	RutaDB       string
	Backend      string
	JWTSecreto   []byte
	JWTDuracion  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Cargar() Config {
	_ = godotenv.Load()

	return Config{
		Puerto:       conTexto("PUERTO", ":8080"),
		RutaDB:       conTexto("RUTA_DB", "for-neos-api.db"),
		Backend:      conTexto("STORAGE", "gorm"),
		JWTSecreto:   []byte(conTexto("JWT_SECRETO", "for-neos-secreto-solo-dev")),
		JWTDuracion:  conDuracion("JWT_DURACION", 24*time.Hour),
		ReadTimeout:  conDuracion("HTTP_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: conDuracion("HTTP_WRITE_TIMEOUT", 10*time.Second),
	}
}

func conTexto(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func conDuracion(clave string, porDefecto time.Duration) time.Duration {
	v := os.Getenv(clave)
	if v == "" {
		return porDefecto
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return porDefecto
	}
	return d
}
