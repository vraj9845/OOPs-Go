package main

// In this program we can see that AvailableFunds() method has two implementations.
// 1 with struct CreditAccount and 2nd time with struct CheckingAccount.
// This will cause confusion and conficts to the func that calls AvailableFunds()
// as 2 methods exist with same name.
// Solution: We create a 3rd method with same name but with receiver as  HybridAccount{}
// Inside the 3rd imlpementation we call our desired implementation 
// i.e. CheckingAccount.AvailableFunds() or CreditAccount.AvailableFunds()
import "fmt"

type CreditAccount struct{}

func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

type CheckingAccount struct{}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 125
}

type HybridAccount struct {
	creditAccount   CreditAccount
	checkingAccount CheckingAccount
}

func (h *HybridAccount) AvailableFunds() float32 {
	return h.creditAccount.AvailableFunds() // or replace with return h.checkingAccount.AvailableFunds()
}

func main() {
	ha := &HybridAccount{}
	fmt.Println(ha.AvailableFunds())
}