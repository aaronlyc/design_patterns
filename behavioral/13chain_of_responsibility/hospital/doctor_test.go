package hospital_test

import (
	. "design_patterns/behavioral/13chain_of_responsibility/hospital"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Doctor", func() {
	cashier := &Cashier{}
	medical := &Medical{}
	doctor := &Doctor{}
	reception := &Reception{}


	BeforeEach(func() {
		medical.SetNext(cashier)
		doctor.SetNext(medical)
		reception.SetNext(doctor)
	})


	Describe("测试病人不同阶段的流程", func() {
		Context("病人刚开始来的流程", func() {
			It("这将从前台开始", func() {
				patient := &Patient{Name: "start man"}
				reception.Execute(patient)
			})
		})

		Context("病人已完成医生流程", func() {
			It("这将显示", func() {
				patient := &Patient{Name: "docker looked", DoctorCheckUpDone: true, RegistrationDone: true}
				reception.Execute(patient)
			})
		})
	})
})
