package main

import (
	"errors"
	"fmt"
	// modulo/<nome_package>
)

// Non si può fare
// Non si può dichiarare una sruct in un package, e i suoi metodi un un altro package
// func (c repo.Car) TrunOn() {}

// Function with return values
// Best practice i valori restituiti devono essere possibilmente sempre i dati utili e per ultimo l'errore
func divide(a, b int) (int, error) {
	if b == 0 {
		// Bisogna ritornare sempre entrambi i valori che abbiamo specificato
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}

// Lezione sui metodi
// Struct definition
type player struct {
	// Sintassi
	// nome proprietà, datatype <type>
	name    string
	age     int
	country string
}

// Function
// Signature: Firma della funzione
func printPlayer(input player) {
	// Utilizziamo Println
	fmt.Println("name:", input.name)
	fmt.Printf("age:\t%d\n", input.age)
	fmt.Printf("contry:\t%q\n", input.country)
}

// Method
// Descrizione:
// la prima parte è l'iniziale della struct a cui si attacca questo metodo
// Il reciver type sarebbe una variabile che prende i valori della stuct.
// Quando si va ad instanziare la struct il metodo viene instanziato nell struct(oggetto).

// Nel primo caso si crea una copia della struct, quindi se viene modificata una properità nell metodo non verrà modificata la struct originale
// value reciver type     => func (p player) print() {}

// Qui invece passiamo un puntatore alla struct player quindi non viene create una copia della struct
// reference reciver type => func (p *player) print() {}

// value receiver type
func (p /* <- Reciver type*/ player) print() {
	fmt.Printf("name:\t%q\n", p.name)
	fmt.Printf("age:\t%d\n", p.age)
	fmt.Printf("country:\t%q\n", p.country)
}

// Reference receiver type
func (p *player) printPtr() {
	fmt.Printf("name:\t%q\n", p.name)
	fmt.Printf("age:\t%d\n", p.age)
	fmt.Printf("country:\t%q\n", p.country)
}

func main() {
	// ford := repo.Car{
	// 	Name: "Ford",
	// }

	num1, num2 := 5, 0

	// Se una funzione può restituire un errore, sempre controllarlo
	res, err := divide(num1, num2)

	// Best practice per loggare l'errore
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("res", res)

	// Initialization
	player1 := player{
		name:    "Jhon",
		age:     24,
		country: "USA",
	}

	// Function invocation
	fmt.Println("Print with a function")
	printPlayer(player1)

	// Method invocation - value receiver type
	fmt.Println("Print with a method - value receiver type")
	player1.print()

	// Method invocation - reference receiver type
	fmt.Println("Print with a method - reference receiver type")
	player1.print()
}
