package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cont := 0
	intentos := 0
	numToCompare := 0

	for cont < 6 {
		num1 := rand.Intn(10) + 1
		num2 := rand.Intn(10) + 1
		num3 := rand.Intn(10) + 1
		num4 := rand.Intn(10) + 1
		numToTest1 := rand.Intn(10) + 1
		numToTest2 := rand.Intn(10) + 1
		numToTest3 := rand.Intn(10) + 1
		numToTest4 := rand.Intn(10) + 1

		if cont == numToCompare || num1 == numToTest1 || num2 == numToTest2 || num3 == numToTest3 || num4 == numToTest4 {
			fmt.Printf("Intento %d: Número correcto (%d)\n", intentos, cont)
			intentos++
		} else {
			fmt.Printf("Intento %d: El número %d no es correcto\n", intentos, cont)
			intentos++
		}

		cont++
		time.Sleep(time.Second)
	}
}
