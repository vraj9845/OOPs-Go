package main

import (
	"OOPs-Go/encapsulation/encapsulationOne"
	"OOPs-Go/encapsulation/encapsulationTwo"
)

func main() {
// enacpsulation implemented with way 1
pizza := encapsulationOne.New("Focaccia bread","Tomato sauce","Mozzarella cheese")
pizza.MakePizza()

// enacpsulation implemented with way 2
oven := encapsulationTwo.NewOven()
oven.BakePizza()
}