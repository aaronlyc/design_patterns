package hospital

import "fmt"

// 定义医生的行为
type Doctor struct {
	next department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Doctor checking patient")
		fmt.Println("waiting some time...")
		p.DoctorCheckUpDone = true
	}
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next department) {
	d.next = next
}
