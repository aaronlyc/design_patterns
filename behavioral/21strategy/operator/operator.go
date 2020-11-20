package operator

type Operator interface {
	Apply(int, int) int
}

type operation struct {
	operator Operator
}

func (o *operation) Operate(i1, i2 int) int {
	return o.operator.Apply(i1, i2)
}

// addition operator
type addition struct{}

func (addition) Apply(i1, i2 int) int {
	return i1 + i2
}

// addition operator
type multiplication struct{}

func (multiplication) Apply(i1, i2 int) int {
	return i1 * i2
}
