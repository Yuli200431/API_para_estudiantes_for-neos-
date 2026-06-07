package main

import "for-neos-api/internal/vivienda/storage"

func main() {
	memoria := storage.NuevaMemoria()

	memoria.SeedViviendas()

	println(len(memoria.ListarViviendas()))
}
