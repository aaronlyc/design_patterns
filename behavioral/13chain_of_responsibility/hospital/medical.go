package hospital

import "fmt"

// 定义药房的行为
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
	} else {
		fmt.Println("Medical giving medicine to patient")
		fmt.Println("waiting some time...")
		p.medicineDone = true
	}
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}
