package hospital

import "fmt"

// 定义收银的行为
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	} else {
		fmt.Println("Cashier getting money from patient")
		p.paymentDone = true
	}
}

func (c *cashier) setNext(next department) {
	c.next = next
}
