package gun

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
