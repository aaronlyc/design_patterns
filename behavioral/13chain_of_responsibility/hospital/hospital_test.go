package hospital

import "testing"

func TestHospital(t *testing.T) {
	// 设计调用流程: 从最后的过程到最开始的过程走
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "test", doctorCheckUpDone: true}

	reception.execute(patient)
}
