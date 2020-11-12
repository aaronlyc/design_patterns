package hospital

// 部门处理者接口
type department interface {
	Execute(*Patient)
	SetNext(department)
}

// 病人
type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}
