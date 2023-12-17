package icecream

import "fmt"

// An IceCreqmRequest is a simple request for ice cream
// In this case, the number is akin to the jobid, and the Recipe is the allocaeed
type IceCreamRequest struct {
	Number uint64
	Recipe string
}

func (i *IceCreamRequest) Satisfied() bool {
	return i.Recipe != ""
}

// Show the customer their final request
func (i *IceCreamRequest) Show() {
	if i.Satisfied() {
		fmt.Printf("\n😍️ Your Ice Cream Order is Ready!\n")
		fmt.Printf("Order Number: %d\n", i.Number)
		fmt.Printf("Recipe:\n%s", i.Recipe)
	} else {
		fmt.Printf("\n😭️ Oh no, we could not satisfy your order!\n")
	}
}
