// Boiler plate: Un pezzo di codice che viene molto spesso usato
package main

//Pronuncia fumt
import (
	"fmt"
	"strconv"
)

func main() {
	// Come craer una variabile
	// var [nome] <tipo>

	// PREMERE F2 PER MINOMINARE TUTTE LE VARIABILI CON L'HO STESSO NOME
	var car string
	car = "Ferrari"
	//%q Verb: Quote, stampa la stringa e mette le virgolette "stringa"
	fmt.Printf("Car: %q\n", car)

	// Creare una varibile più velocmente, dando il type in automatico
	cheapCar := "FIAT"
	fmt.Printf("Cheaper car: %q\n", cheapCar)

	// Creare più varibabli insieme
	// Si puà fare purchè siano dello stesso tipo
	// var [nome], [nome2] <tipo per entrambi>
	var myAge, majorAge int
	myAge = 16
	majorAge = 18
	isAdult := myAge > majorAge

	//Pronuncia fumt
	// %v Verb: Semplice formattazione quello che entra viene scritto senza formattazione
	fmt.Printf("Are you adult? %v\n", isAdult)

	// Cast da ubter ia stringa
	myBalance := 1000
	// Converte un Itero in una stringa e proviene dal packge strconv che implementa delle conversioni da stringa e a stringa.
	// https://pkg.go.dev/strconv
	// Itoa converte un intero in una strainga (int to string)
	myBalanceStr := strconv.Itoa(myBalance)
	fmt.Printf("My balance is %q", myBalanceStr)
	// Truffare il compilatore, se ad esempio hai una variabile che non hai usato: BlankIdentifyer
	_ = myBalanceStr
}
