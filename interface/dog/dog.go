package dog

import (
	"fmt"
	"strings"
)

type Dog struct {
	Name string `json:"name"`
}

func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, fmt.Errorf("Please type the name")
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("The name is too long")
	}

	return &Dog{
		Name: name,
	}, nil
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Speak() string {
	return d.Name + " noi rang: Se co nhung con ca phai tra gio!!"
}

func (d *Dog) Eat() string {
	return d.Name + " dang an hakyfood"
}
