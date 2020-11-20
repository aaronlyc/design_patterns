package gun

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

func (a *ak47) getName() string {
	return a.name + " rewrite"
}