package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World...")
}

func HandlePayments(partySelection string, amount string, userinfo string) {
	// depending on which payment method is selected, forward the payment details to that party
	// for now, just print the details
	fmt.Println("Party Selection: ", partySelection)
	fmt.Println("Amount: ", amount)
	fmt.Println("User Info: ", userinfo)
}
