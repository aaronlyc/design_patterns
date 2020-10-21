package hospital

// 部门处理者接口
type department interface {
	execute(*patient)
	setNext(department)
}

// 病人
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
