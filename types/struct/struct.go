package main

import "fmt"

type product struct{
	name string
	price float64
	discount float64
}

//Methods: functions that deal with data struct. functions that have receivers.
//Product here is the owner (receiver) of this func, once the func can only receive products.
func (p product) priceWithDiscount() float64{
	return p.price * (1 - p.discount)
}

func main()  {
	var product1 product
	product1 = product{
		name: "pencil",
		price: 1.50,
		discount: 0.25,
	}
	fmt.Println(product1, product1.priceWithDiscount())

	//Another way to declare is...
	product2 := product{"notebook", 10.2, 1}
	fmt.Println(product2, product2.priceWithDiscount())
}