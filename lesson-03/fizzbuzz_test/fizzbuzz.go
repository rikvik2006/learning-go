package fizzbuzztest

import "strconv"

func Fizzbuzz(input int) string {
	// Operatore Modulo: calcola il resto
	if (input%3) == 0 && (input%5) == 0 {
		return "FizzBuzz"
	} else if (input % 5) == 0 {
		return "Buzz"
	} else if (input % 3) == 0 {
		return "Fizz"
	} else {
		// Converte un intero in una stringa
		// Non restituisce errori perché non c'è bisogno di restituirli
		return strconv.Itoa(input)
	}
}
