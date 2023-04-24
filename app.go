package main

import (
	"fmt"
	"net/http"
)

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hola Mundo")

}

// Ruta
func main() {
	http.HandleFunc("/", Index)

	// Crea el servidor
	fmt.Println("El servidor esta corriendo enel puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}

// Para poder retornar algo ,vamos a tomar una ruta y un handler que es una
// función que va Lo que va a hacer es
// devolver con un mensaje a la petición
// del cliente
