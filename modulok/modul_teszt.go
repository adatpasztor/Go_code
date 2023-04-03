package main

import (
	"fmt"

	"modulok/sajat"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println("modul_teszt->ok")

	fmt.Println("sajat.s2 S2_var->", sajat.S2_var)
	sajat.S1_info()
}
