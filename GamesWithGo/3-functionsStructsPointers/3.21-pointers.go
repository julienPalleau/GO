package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func addOne(num int) {
	num = num + 1
}

func addOneUsingPtr(num *int) {
	*num = *num + 1 // deferencement d'un pointeur: give me the thing that the pointer point to
}

func whereIsBadGuy(guy *badGuy) {
	x := guy.pos.x // In go when using a struct it can automatically derefence the pointer 
	y := guy.pos.y // this is why here we don't need to deference
	fmt.Println("badGuy est a la location suivante: (", x, ",", y, ")")
}

func main() {
	x := 5
	fmt.Println("print valeur de x:", x)

	// xPtr := &x //it means give me the address of this value
	// fmt.Println(xPtr)

	println()
	var xPtr *int = &x
	fmt.Println("print adresse de x:", xPtr)

	println()
	addOne(x)
	fmt.Println("print valeur de x:", x)

	println()
	addOneUsingPtr(xPtr)
	fmt.Println("print valeur de x apres que l'on est passe l'adresse de x a addOneUsingPtr: ", x) // here we print the value of x which is equal to 5.

	println()
	p := position{4, 2}
	badguy := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(&badguy) // On appel whereIsBadGuy en passant l'adresse de badguy avec &badguy au lieu de passer la struct
	// c'est beaucoup plus leger on passe qu'une adresse au lieu de une string un int une struct position
}
