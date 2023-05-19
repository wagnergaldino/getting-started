package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello:", quote.Hello())
	fmt.Println("Go:", quote.Go())
	fmt.Println("Glass:", quote.Glass())
	fmt.Println("Opt:", quote.Opt())
}
