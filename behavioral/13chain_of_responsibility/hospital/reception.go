package hospital

import "fmt"

// 定义前台接待员的行为
type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
	} else {
		fmt.Println("Reception registering patient")
		fmt.Println("waiting some time...")
		p.RegistrationDone = true
	}
	r.next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}
