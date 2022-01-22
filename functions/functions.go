package main

import (
	"fmt"
)

func sayHey(name string) {
	fmt.Println("Hello", name)

}

//some edit idk    L + ratio

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSociable(name string) {
	sayHey(name)
	fmt.Println("Hows the weather?")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABuch() {
	fmt.Println("Hello")
	sayHelloABuch()
}

func main() {
	beSociable("Bob")
	beSociable("Alice")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))
	fmt.Println(x)

	sayHelloABuch()
}
