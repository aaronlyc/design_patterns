package operator

import "testing"

func TestOperation(t *testing.T) {
	//addition 3 + 5
	a := operation{addition{}}
	t.Logf("the addition operation result is:%d\n", a.Operate(3, 5))

	//multiplication 3 * 5
	a = operation{multiplication{}}
	t.Logf("the multiplication operation result is:%d\n", a.Operate(3, 5))
}
