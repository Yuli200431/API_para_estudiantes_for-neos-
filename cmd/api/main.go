package main

import (
	"fmt"
	"net/http"

	//"github.com/go-chi/chi/v5"

	//"for-neos-api/internal/transporte/handlers"
	//"for-neos-api/internal/transporte/storage"
	
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API para estudiantes foráneos funcionando")
	})

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
