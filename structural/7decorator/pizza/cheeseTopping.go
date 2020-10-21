package pizza

import "fmt"

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	fmt.Println("this is cheese price")
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
