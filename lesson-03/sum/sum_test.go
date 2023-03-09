package sum

import "testing"

// 1. file called sum_test.go
// 2. package "sum" (White-Box Testing)
// 3. The name of the function must follow TestXxx
// 4. "t" must be a param of type *testing.T

func TestSum(t *testing.T) {
	// Arrange
	var num1, num2 int
	num1 = 3
	num2 = 2

	// act
	res := Sum(num1, num2)

	// assert
	if res != 5 {
		t.Errorf("Expected %d got %d", 5, res)
	}
}

// Eseguire un test
// go test <- Mi dice solo se il test è passato o e falitto
// go test -v <- Verbose, mi di ce cosa è stato eseguito
