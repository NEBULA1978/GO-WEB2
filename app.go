package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola Mundo")
	})
	http.ListenAndServe("localhost:3000", nil)
}

// Para poder retornar algo ,vamos a tomar una ruta y un handler que es una
// función que va Lo que va a hacer es
// devolver con un mensaje a la petición
// del cliente
