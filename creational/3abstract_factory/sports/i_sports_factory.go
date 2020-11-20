package sports

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSprotsFactory(brand string) (iSportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed.")
}
