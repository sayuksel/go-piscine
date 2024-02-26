package main

import (
	"fmt"
	// "strings"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println("------------")
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println("------------")

	fmt.Println(piscine.Index("Ola!", "Ol"))
	fmt.Println("------------")

	// strings.Index()
}
