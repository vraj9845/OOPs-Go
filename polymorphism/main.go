package main

import "fmt"

// Interfaces are implemented implicitly in Go
// Go uses structural typing to determine if a type implements an interface
type pizzaMaker interface {
	makePizza()
}

type ingredients struct {
	bread  string
	gravy  string
	cheese string
}

// type ingredients implements makePizza()
func (i ingredients) makePizza() {
	fmt.Printf("We will make pizza using %v ,%v and %v.\n", i.bread, i.cheese, i.gravy)
}

type oven struct{}

// type oven implements makePizza()
func (o oven) makePizza() {
	fmt.Println("Pizza is baked for 30 minutes or until pizza becomes crispy and golden brown.")
}

func main() {
	var pm pizzaMaker
	pm = ingredients{"Focacia bread", "Tomatoe gravy", "Mozarella Cheese"}
	pm.makePizza()
	pm = oven{}
	pm.makePizza()
}
