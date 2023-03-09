package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definiamo una funzone ma utilizziamo una funzionalità di go per definire i le variabili dei valori di ritorno direttamente dalla signature della funzione

// Legenda: ?: Opzione !: Obligatorio
// func <nome>(<paramatri [nome]<tipo>>) (Valori di return: [?nomevariabile]<!tipo>)
func decimalToRoman(decimalValue int) (romanNumber string, err error) {
	// var romanNumber string
	// var err error

	// 1. Check if the "decimalValue" is inside the range 1, 3999
	if decimalValue < 1 || decimalValue > 3999 {
		//Errorf: Formatta l'errore con i tamplate (verb)
		return "", fmt.Errorf("%d in not in the valid range (1, 3999)", decimalValue)
	}

	// 2. Define a map with decimalValue as key and roman representation as value
	romanNums := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// 3. Based on the previous map, build a slice whit only the keys
	romanKeys := make([]int, 0)
	for key := range romanNums { // Prendiamo solo la key, aloposto di key value
		romanKeys = append(romanKeys, key)
	}

	// 4. Sort the lisce in ascending order
	//
	sort.Ints(romanKeys)

	// 5. The algorithm...
	for decimalValue > 0 {
		for i := 0; i < len(romanKeys); i++ {
			if i == len(romanKeys)-1 || (decimalValue >= romanKeys[i] && decimalValue < romanKeys[i+1]) {
				// Decrement operator -=
				decimalValue -= romanKeys[i]
				romanNumber += romanNums[romanKeys[i]]
				break // end the looop.Esce dal loop
				// continue => goes to the next iteration, continue blocca elesecuzione del codice a quel puto e va alla prossima iterazione, praticaemtne salta di un giro
			}
		}
	}

	return romanNumber, nil
	// Si possono anche omettere le variabili perché le abbiamo definite come valori di ritorno
	// return <-
}

func main() {
	// 18 => XVIII

	// Creiamo una slice di valori per testare l'algroitmo
	values := []int{0, 4, 18, 34, 50, 68, 999, 1415, 4000}

	//  K, V: ma la key non ci interessa quindi utilizzaimo il blank idetifyer
	for _, v := range values {
		res, err := decimalToRoman(v)
		if err != nil {
			fmt.Printf("Value %d reised the err %q\n", v, err.Error())
			continue
		}

		fmt.Printf("Value %d converted to %q\n", v, res)
	}
	// break vs continue

	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i) // when i == 3 this line won't be executed
	}

	// Stampa una stringa da 30 caratteri di cancelleti ("#")
	fmt.Println(strings.Repeat("#", 30))

	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i) // when i == 3 the lop ends and the "3", "4" won't be printed
	}
}
