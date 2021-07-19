// password must have => lowercase + uppercase letters,a number,a symbol and must be 8/+ xters long

package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	if len(pwR) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	//Loop over the multi-byte xters one at a time
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		//result is true if either of these functions returns true
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	// checks all four variables

	return hasUpper && hasLower && hasNumber && hasSymbol
} // function

func main() {
	if passwordChecker("") {
		fmt.Println("password is good!")
	} else {
		fmt.Println("bad password")

	}

	//call the function with a valid password

	if passwordChecker("This!15A") {
		fmt.Println("good password")
	} else {
		fmt.Println("bad password")
	}

}
