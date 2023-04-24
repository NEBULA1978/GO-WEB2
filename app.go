package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Estructura datos
type Users struct {
	Name   string
	Skills string
	Age    int
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	// Con _ Para no capturar el error
	template, err := template.ParseFiles("template/index.html")

	user := Users{Name: "Ramon Pascual", Skills: "desarrolando webs", Age: 27}

	if err != nil {
		panic(err)
	} else {
		// Mostramos mensaje con respuesta
		template.Execute(rw, user)
	}

	fmt.Fprintln(rw, "Hola Mundo")

}

// Ruta
func main() {
	http.HandleFunc("/", Index)

	// Crea el servidor
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}

// Para poder retornar algo ,vamos a tomar una ruta y un handler que es una
// función que va Lo que va a hacer es
// devolver con un mensaje a la petición
// del cliente

// DESCRIPCION

// Este código crea un servidor web simple en Go utilizando el paquete net/http y el paquete text/template. Define una estructura de datos llamada Users que contiene tres campos: Name, Skills y Age.

// La función Index es un controlador HTTP que se encarga de procesar las solicitudes HTTP entrantes. Dentro de esta función, primero se intenta analizar un archivo de plantilla HTML llamado template/index.html. Luego, crea una instancia de Users con datos específicos y, si no hay errores al analizar la plantilla, ejecuta la plantilla con el objeto Users como entrada y escribe la salida en la respuesta HTTP. Por último, también escribe "Hola Mundo" en la respuesta HTTP.

// La función main configura la función Index como el controlador para la ruta principal ("/") y crea un servidor HTTP en el puerto 3000, imprimiendo en la consola la dirección para acceder al servidor.

// En resumen, este código crea un servidor web simple que responde a las solicitudes entrantes en la ruta principal, procesa una plantilla HTML y escribe la salida en la respuesta HTTP junto con el mensaje "Hola Mundo".
