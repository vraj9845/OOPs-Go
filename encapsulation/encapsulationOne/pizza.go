package encapsulationOne

// This is program shows how we can achieve encapsulation in Go
// Note: Only methods are accessible as they exported.
//		 Data/Variables are not accessible to the external packages.
import "fmt"

type ingredients struct {
	bread string
	gravy string
	cheese string
}

// method New() is exported
func New(ingredient1,ingredient2,ingredient3 string) ingredients{
return ingredients{
	bread: ingredient1,
	gravy: ingredient2,
	cheese: ingredient3,
}
}
 
// method MakePizza() is exported
func (i ingredients)MakePizza() {
	fmt.Printf("We will make pizza using %v ,%v and %v.\n",i.bread,i.cheese,i.gravy)
}