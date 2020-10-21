package menu_price

import "fmt"

type dish struct {
	number   int
	perPrice int
	name     string
}

func (d *dish) addPrice(i *int) {
	fmt.Printf("++dish %s unit price is: %d, number is: %d, total price is %d\n", d.name, d.perPrice, d.number, d.getPrice())
	*i = *i + d.getPrice()
}

func (d *dish) getPrice() int {
	return d.number * d.perPrice
}
