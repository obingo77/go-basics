//switch expresion toto check whether a number is in a specific range

package main

import (
	"fmt"
	"time"
)

func main() {

	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Born on weekend")

	default:
		fmt.Println("Born on another day")
	}

}
