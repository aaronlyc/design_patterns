package hospital

import "fmt"

// 定义医生的行为
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Doctor checking patient")
		fmt.Println("waiting some time...")
		p.doctorCheckUpDone = true
	}
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}
