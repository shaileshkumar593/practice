package main

import "fmt"

type vechicle interface {
	accelerate()
}

func foo(v vechicle) {
	fmt.Println(v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelrating")
}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelarated fast")

}

func main() {
	c1 := car{"suzuki", "blue"}
	t1 := toyota{"toyota", "Red", 100}
	c1.accelerate()
	t1.accelerate()

	foo(c1)
	foo(t1)
}
