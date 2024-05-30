package main

import (
	"fmt"
	"regexp"
)

func main() {

	rx := regexp.MustCompile(`Andr.{2,}w`)

	var Name string

	fmt.Println("What is your name? ")
	fmt.Scan(&Name)

	if rx.MatchString(Name) {

		fmt.Printf("Fuck you, %s! ", Name)

	} else {
		fmt.Printf("Hello, %v!", Name)
	}

}
