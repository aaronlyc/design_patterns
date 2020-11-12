package hospital

import "fmt"

// 定义药房的行为
type Medical struct {
	next department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
	} else {
		fmt.Println("Medical giving medicine to patient")
		fmt.Println("waiting some time...")
		p.MedicineDone = true
	}
	m.next.Execute(p)
}

func (m *Medical) SetNext(next department) {
	m.next = next
}
