package romannumeralstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// table driven test
func TestDecimalToRoman(t *testing.T) {
	// Definizione test suite
	// Ovvero una slice di delle struct anonime
	testSuite := []struct {
		// Definizione struct
		name    string
		decimal int
		res     string
	}{
		// Wrappare tutti gli ementi che fatto parte della slice di struct
		{
			// Test che rispetano la sintassi della struct, perché quello è il tipo da seguire, e come se fose un array di classi
			"When4_ThenIV",
			4,
			"IV",
		},
		// Cremo un nuovo test molto semplicemente
		{
			"When8_ThenVIII",
			8,
			"VIII",
		},
		{
			"When4000_ThenStringEmpty",
			4000,
			"",
		},
		{
			"When8000_ThenStringEmpty",
			8000,
			"",
		},
	}

	// Tests execution
	// tt indica ogni struct anonima della sclice, tt sta per testSuite, e si usa per ciclare i test table driven test
	for _, tt := range testSuite {
		// Execution
		t.Run(tt.name, func(t *testing.T) {
			res, err := DecimalToRoman(tt.decimal)

			assert.Equal(t, tt.res, res)

			// Check if expected != "" there is an error
			if tt.res == "" {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

// func TestDecimalToRoman(t *testing.T) {
// 	res, err := DecimalToRoman(4)

// 	assert.Equal(t, "IV", res)
// 	// if res != "IV" {
// 	// 	t.Errorf("expected %q got %q", "IV", res)
// 	// }

// 	assert.Nil(t, err)
// 	// if err != nil {
// 	// 	t.Errorf("expected nil got %v", err)
// 	// }

// }

// func TestDecimalToRoman_Err(t *testing.T) {
// 	res, err := DecimalToRoman(4000)

// 	assert.Equal(t, "", res)
// 	assert.NotNil(t, err)
// }
