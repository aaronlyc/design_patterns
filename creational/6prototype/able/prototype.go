package able

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManger struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManger() *PrototypeManger {
	return &PrototypeManger{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
