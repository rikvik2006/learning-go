package main

import "fmt"

// Lezione sugli struct (credo siano gli oggeti)

// Struct definition
// Se il nome ha una minuscola allora è locale al package
// Mentre se inizia con la lettera maiuscola vuol dire che può essere vista fuori dal package

// "player" scoped to this package
// "Player" can be refereced from others packages (e.g. repo. Player)
type player struct {
	// Sintassi
	// nome proprietà, datatype <type>
	name    string
	age     int
	country string
}

func main() {
	// Initialization
	player1 := player{
		name:    "Jhon",
		age:     24,
		country: "USA",
	}

	// Verb %v nessuna formattazione
	fmt.Printf("Player1: %v\n", player1)
}
