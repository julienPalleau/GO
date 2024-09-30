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

func whereIsBadGuy(guy badGuy) {
	x := guy.pos.x
	y := guy.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {
	// First solution to initialise your struct
	// var p Position
	// p.x = 4
	// p.y = 2

	// Second solution to initialize your struct
	p := position{4, 2}
	fmt.Println(p.x)
	fmt.Println(p.y)

	badguy := badGuy{"Jabba The Hut", 100, p}
	fmt.Println(badguy) // print will put {} around a struct
	whereIsBadGuy(badguy)
}
