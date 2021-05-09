package messagePassingOne

// Note: WIP
import (
	"fmt"
)

type ingredients struct {
	bread  string
	gravy  string
	cheese string
}

func New(ingredientsChannel chan string) {
	s := make([]string, 3)
	var y ingredients
	go func(ingredientsChannel chan string) {
		one := <-ingredientsChannel
		two := <-ingredientsChannel
		three := <-ingredientsChannel
		y = ingredients{one,two,three}
	}(ingredientsChannel)

fmt.Println(y)
	i := ingredients{bread: s[0], gravy: s[1], cheese: s[2]}
	fmt.Println(i)
	i.makePizza()
}

func (i ingredients) makePizza() {
	fmt.Printf("We will make pizza using %v ,%v and %v.\n", i.bread, i.cheese, i.gravy)
}
