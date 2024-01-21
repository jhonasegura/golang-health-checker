package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configura el manejador de salud
	http.HandleFunc("/health", healthHandler)

	// Configura el servidor y escucha en el puerto 8080
	port := 8080
	fmt.Printf("Servidor de salud escuchando en el puerto %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}

// Manejador de salud
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica aquí cualquier condición de salud que desees
	// Por ejemplo, si la base de datos está disponible, etc.

	// Si todo está bien, responde con un código de estado 200 (OK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
