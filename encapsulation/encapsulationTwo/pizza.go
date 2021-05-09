package encapsulationTwo

// This is program shows how we can achieve encapsulation in Go
// Note: Only methods are accessible as they exported.
//		 Data/Variables are not accessible to the external packages.
import "fmt"

type Interface interface {
	BakePizza()
}

type oven struct{}

// NewOven is exported
func NewOven() oven {
	return oven{}
}

// BakePizza is exported
func (o oven)BakePizza() {
	fmt.Println("Pizza is baked for 30 minutes or until pizza becomes crispy and golden brown.")
}