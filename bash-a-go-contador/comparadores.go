package main

import "fmt"

func main() {
	// Comparar dos números enteros
	a := 5
	b := 10
	if a < b {
		fmt.Println("a es menor que b")
	}

	// Comparar dos cadenas de texto
	str1 := "hola"
	str2 := "mundo"
	if str1 == str2 {
		fmt.Println("Las cadenas son iguales")
	}

	// Comparar dos booleanos
	bool1 := true
	bool2 := false
	if bool1 && !bool2 {
		fmt.Println("bool1 es verdadero y bool2 es falso")
	}
}

// En el primer ejemplo, se declaran e inicializan las variables a y b usando el operador :=, y luego se comparan para verificar si a es menor que b.

// En el segundo ejemplo, se declaran e inicializan las variables str1 y str2 usando el operador :=, y luego se comparan para verificar si son iguales.

// En el tercer ejemplo, se declaran e inicializan las variables bool1 y bool2 usando el operador :=, y luego se comparan para verificar si bool1 es verdadero y bool2 es falso usando el operador && y el operador de negación !.

// En este código, se importa el paquete fmt para poder imprimir en la consola usando fmt.Println(). Luego, dentro de la función main(), se declaran e inicializan variables usando el operador :=, y se comparan usando operadores como <, ==, &&, y !.
