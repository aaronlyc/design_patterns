package hospital

import "fmt"

// 定义收银的行为
type Cashier struct {
	next department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	} else {
		fmt.Println("Cashier getting money from patient")
		p.PaymentDone = true
	}
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}
