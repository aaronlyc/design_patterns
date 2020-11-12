package hospital_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"design_patterns/behavioral/13chain_of_responsibility/hospital"
)

var _ = Describe("Doctor", func() {
	cashier := &hospital.Cashier{}
	medical := &hospital.Medical{}
	doctor := &hospital.Doctor{}
	reception := &hospital.Reception{}


	BeforeEach(func() {
		medical.SetNext(cashier)
		doctor.SetNext(medical)
		reception.SetNext(doctor)
	})

	Describe("测试病人不同阶段的流程", func() {
		Context("病人刚开始来的流程", func() {
			It("这将从前台开始", func() {
				patient := &hospital.Patient{Name: "start man"}
				reception.Execute(patient)
			})
		})

		Context("病人已完成医生流程", func() {
			It("这将显示", func() {
				patient := &hospital.Patient{Name: "docker looked", DoctorCheckUpDone: true, RegistrationDone: true}
				reception.Execute(patient)
			})
		})
	})
})
