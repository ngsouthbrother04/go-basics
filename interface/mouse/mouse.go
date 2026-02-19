package mouse

import "nnama.com/interface/services"

type Mouse struct {
	name string
}

func New(name string) (services.Animal, error) {
	return &Mouse{
		name: name,
	}, nil
}

func (m *Mouse) GetName() string {
	return m.name
}

func (m *Mouse) Speak() string {
	return "Chít chít"
}

func (m *Mouse) Eat() string {
	return "Ăn phô mai"
}

func (m *Mouse) Sleep() bool {
	return true
}
