package main

import "fmt"

type item struct {
	productID int
	quantity int
	price float64
}

type order struct {
	userID int
	//Nested struct here. Slice of struct item.
	items []item
}

//Func returns the total value of an order.
func(o order) totalValue() float64{
	total := 0.0
	for _, item := range o.items{
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		//Populating a nested struct.
		items: []item{
			{1,2,12.10} ,
			{2,3,4.0},
		},
	}
	fmt.Printf("The valus is: R$%.2f", order.totalValue())
}
