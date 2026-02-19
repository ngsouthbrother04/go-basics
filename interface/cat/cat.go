package cat

import (
	"fmt"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, fmt.Errorf("Please type the name")
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("The name is too long")
	}

	return &Cat{
		Name: name,
	}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (d *Cat) Speak() string {
	return "Meow!"
}
