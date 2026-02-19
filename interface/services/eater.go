package services

type Eater interface {
	Eat() string
}

func Eating(e Eater) string {
	return e.Eat()
}
