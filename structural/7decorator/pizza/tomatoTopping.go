package pizza

import "fmt"

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	fmt.Println("this is tomato price")
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}
