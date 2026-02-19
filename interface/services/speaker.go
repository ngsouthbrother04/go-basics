package services

type Speaker interface {
	Speak() string
	GetName() string
}

func Greet(s Speaker) string {
	return s.Speak()
}

func GetName(s Speaker) string {
	return s.GetName()
}
