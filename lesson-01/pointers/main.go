package main

import (
	"fmt"
	"strings"
)

func main() {
	//Int value <int>
	var intVal int
	intVal = 34

	fmt.Printf("intVal value: %d\n", intVal)

	// AddressOf => "&"
	fmt.Printf("intVal memory address: %v\n", &intVal)

	// Pointers
	// * during variable definition means pointer variable
	// Pointer è un type che si mette prima del tipe originale
	var intPtr *int
	_ = intPtr // NOOOO!
	intPtr = &intVal
	fmt.Printf("intPtr memory address: %v\n", intPtr)
	// * => Dereferece operator
	fmt.Printf("intPtr value: %d\n", *intPtr)

	// Se utilizzo l'asterisco qunado definisco la variabile deifinisco il pointer, se utilizzo l'asterisco su un pointer, vado a predere il valore dell indirizzo di memoria al quale sta puntato il pointer

	//Let's change the value
	intVal = 3
	fmt.Println(strings.Repeat("#", 50))
	fmt.Printf("intPtr memory addres: %v\n", intPtr)
	fmt.Printf("intVal value: %d\n", intVal)
	fmt.Printf("intPtr value: %d\n", *intPtr)

	stringPointer()
}

func stringPointer() {
	fmt.Println(strings.Repeat("#", 50))
	stringValue := "Wow Pointer"
	fmt.Printf("stringValue value: %q\n", stringValue)
	fmt.Printf("stringValue memory address: %v\n", &stringValue)

	var stringPtr *string
	stringPtr = &stringValue
	fmt.Printf("stringPtr memory address: %v\n", stringPtr)
	fmt.Printf("stringPtr value: %q\n", *stringPtr)
}
