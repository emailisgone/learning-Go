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

func whereisBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	p := position{4, 2}
	fmt.Println(p.x)

	badguy := badGuy{"Bad guy", 100, p}
	whereisBadGuy(badguy)
}
