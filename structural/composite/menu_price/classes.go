package menu_price

import "fmt"

type classes struct {
	dishes []price
	name   string
}

func (c *classes) addPrice(i *int) {
	fmt.Printf("add the classes %s for dishes\n", c.name)
	for _, dish := range c.dishes {
		dish.addPrice(i)
	}
}

func (c *classes) addDish(d price) {
	c.dishes = append(c.dishes, d)
}
