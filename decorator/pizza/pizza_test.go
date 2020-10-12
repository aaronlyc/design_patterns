package pizza

import (
	"fmt"
	"testing"
)

func TestPizza(t *testing.T) {
	pizza := &veggeMania{}
	fmt.Printf("the pizza base price is: %d\n", pizza.getPrice())
	fmt.Println()

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}
	fmt.Printf("the pizza with cheese price is: %d\n", pizzaWithCheese.getPrice())
	fmt.Println()

	pizzaWithTomato := &tomatoTopping{
		pizza: pizza,
	}
	fmt.Printf("the pizza with tomato price is: %d\n", pizzaWithTomato.getPrice())
	fmt.Println()

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("the pizza with tomato and cheese price is: %d\n", pizzaWithCheeseAndTomato.getPrice())
}
