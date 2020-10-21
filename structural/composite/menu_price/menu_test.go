package menu_price

import (
	"fmt"
	"testing"
)

func TestMenu(t *testing.T) {
	dish1 := &dish{
		number:   10,
		perPrice: 1,
		name:     "鸡蛋",
	}
	dish2 := &dish{
		number:   6,
		perPrice: 2,
		name:     "萝卜",
	}
	dish3 := &dish{
		number:   5,
		perPrice: 3,
		name:     "白菜",
	}
	dish4 := &dish{
		number:   1,
		perPrice: 40,
		name:     "牛肉",
	}

	classes1 := &classes{
		name:   "肉类",
	}

	classes2 := &classes{
		name:   "蛋类",
	}

	classes3 := &classes{
		name:   "蔬菜",
	}

	classes1.addDish(dish4)
	classes2.addDish(dish1)
	classes3.addDish(dish2)
	classes3.addDish(dish3)

	classes0 := &classes{
		name: "总价",
	}
	classes0.addDish(classes1)
	classes0.addDish(classes2)
	classes0.addDish(classes3)
	var total int
	classes0.addPrice(&total)
	fmt.Printf("the total is %d\n", total)
}
