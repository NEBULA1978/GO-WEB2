package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Definimos una estructura de datos para los usuarios
type Users struct {
	Name   string
	Skills string
	Age    int
}

// Handler para la ruta principal
func Index(rw http.ResponseWriter, r *http.Request) {
	// Parseamos el archivo de plantilla HTML
	template, err := template.ParseFiles("template/index.html")

	// Creamos un usuario de ejemplo
	user := Users{Name: "Ramon Pascual", Skills: "desarrollando webs", Age: 27}

	// Comprobamos si ha habido algún error al parsear la plantilla
	if err != nil {
		// Si hay un error, lo capturamos y mostramos un mensaje de error
		panic(err)
	} else {
		// Si no hay errores, ejecutamos la plantilla con los datos del usuario y la escribimos en la respuesta HTTP
		template.Execute(rw, user)
	}

	// Escribimos un mensaje adicional en la respuesta HTTP
	// Sale debajo como si fuera un Footer
	// fmt.Fprintln(rw, "Hola Mundo")
}

// Función principal que establece las rutas y crea el servidor HTTP
func main() {
	// Asociamos la ruta principal ("/") con el handler Index
	http.HandleFunc("/", Index)

	// Creamos y configuramos el servidor HTTP en el puerto 3000
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
