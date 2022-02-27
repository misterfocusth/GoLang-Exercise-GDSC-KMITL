package exercise

import (
	"fmt"
)

var a, b, c = "BNK", 48, true

func Exercise1() {
	a, b, c := "BNK", 8, true

	word := fmt.Sprintf("%v.%v.%v", a, b, c)
	fmt.Println(word)

	type misterfocus string
	var focus misterfocus
	fmt.Printf("%T", focus)
}
