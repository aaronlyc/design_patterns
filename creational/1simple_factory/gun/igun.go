package gun

type iGun interface {
	setName(string)
	setPower(int)
	getName() string
	getPower() int
}