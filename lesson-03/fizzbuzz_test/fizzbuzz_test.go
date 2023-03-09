package fizzbuzztest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Cleen Test
func TestFizzBuzz(t *testing.T) {
	// Sub test si utlizza il metodo Run che proviene dalla variabile t, all interno si mette il nome del subTest (nome qualunque) e una calback function (funzione anonima)
	t.Run("When1_Then1", func(t *testing.T) {
		// Dentro la call back function in questo caso un sub test,
		// Nome subtest: "When1_Then1", Subtest: é il codice dentro la funzione anonima
		res := Fizzbuzz(1)
		assert.Equal(t, "1", res)
	})

	t.Run("When3_ThenFizz", func(t *testing.T) {
		res := Fizzbuzz(3)
		// La variabile t serve in caso falissce qualcosa, e quindi server per dire quale test è fallito
		assert.Equal(t, "Fizz", res)
	})

	t.Run("When5_ThenBuzz", func(t *testing.T) {
		res := Fizzbuzz(5)
		assert.Equal(t, "Buzz", res)
	})

	t.Run("When30_ThenFizzBuzz", func(t *testing.T) {
		res := Fizzbuzz(30)
		assert.Equal(t, "FizzBuzz", res)
	})
}

// func TestFizzBuzz_1(t *testing.T) {
// 	// Arrange && Assert
// 	res := Fizzbuzz(1)

// 	// Assert

// 	// Dobbiamo installare il plugin con il comando go get <url> in questo caso `github.com/stretchr/testify/assert`
// 	// Se modifichiamo le dipendenze nel file go.mod dobbiamo runnare il comando `go mod tidy`

// 	// Sia "1", che res sono di tipo stringa, quindi si può testare
// 	assert.Equal(t, "1", res)
// }

// func TestFizzBuzz_3(t *testing.T) {
// 	res := Fizzbuzz(3)

// 	assert.Equal(t, "Fizz", res)
// }

// func TestFizzBuzzd(t *testing.T) {
// 	// arrange
// 	var input int
// 	input = 1

// 	// act
// 	// Short annotation
// 	res := Fizzbuzz(input)

// 	// assert
// 	if res != "1" {
// 		t.Errorf("expected %q got %q", "1", res)
// 	}
// }
