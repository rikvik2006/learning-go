package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var <nome> [grandezza]<tipo>
	// Value type
	var grades [3]int
	grades[0] = 8
	grades[1] = 5
	grades[2] = 9
	// Verb default, nessuna formattazione
	// Quanto si è indecisi del verb da usare
	// %v ci salva, perché non applica nessuna formattazione
	fmt.Printf("Grades: %v\n", grades)

	// Slices
	// <nome_var> := []<type>{valore1, valore2, ...}
	// Reference type
	friends := []string{"John", "Mary"}
	friends = append(friends, "Bob")
	fmt.Printf("Friends: %v\n", friends)

	for index, value := range friends {
		fmt.Printf("%d\t%q\n", index, value)
	}
	// Se non mi interessa l'index
	for _, value := range friends {
		fmt.Printf("%q\n", value)
	}

	// Maps
	// "john": "3A",
	// "alice": "1B"

	//Keyword: map[key: type]value: type
	//es: map[<type>]<type>
	students := make(map[string]string)
	// Settato per il valore "john"
	students["john"] = "3A"
	students["alice"] = "1B"

	for key, value := range students {
		fmt.Printf("Student %q belongs to the %q class\n", key, value)
	}
}
