package hospital

import "fmt"

// 定义前台接待员的行为
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
	} else {
		fmt.Println("Reception registering patient")
		fmt.Println("waiting some time...")
		p.registrationDone = true
	}
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}
