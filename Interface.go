// Golang program illustrates how
// to implement an interface
package main

import "fmt"

// Creating an interface
type tank interface {
	// Methods
	tankArea() float64
	tankVolume() float64
}

type cylinderTank struct {
	radius float64
	height float64
}

type cubicTank struct {
	side float64
}

// Implementing methods of
// the tank interface
func (m cylinderTank) tankArea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m cylinderTank) tankVolume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// the tank interface
func (c cubicTank) tankArea() float64 {

	return c.side * c.side * 6
}

func (c cubicTank) tankVolume() float64 {

	return c.side * c.side * c.side
}

func calculateAreaAndVoulme(t tank) {
	fmt.Println("=====================")
	fmt.Println("Object=", t)
	fmt.Println("Area=", t.tankArea())
	fmt.Println("Volume=", t.tankVolume())
	fmt.Println("=====================")
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	var t1, t2 tank
	t1 = cylinderTank{10, 14}
	t2 = cubicTank{10}
	calculateAreaAndVoulme(t1)
	calculateAreaAndVoulme(t2)
}
