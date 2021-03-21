package app

import (
	"fmt"
	"errors"
	"example.com/practise/fundamental"
)

type App struct {
	name string
}

type AppFactory struct {
}

func (factory AppFactory) Create(dep interface{}) (interface{}, error) {
	fmt.Println("AppFactory Create")

	name, ok := dep.(string)
	if !ok {
		return nil, errors.New("dep is not string")
	}
	
	return App { name }, nil
}

func init() {
	fundamental.RegisterFactory((*App)(nil), AppFactory {})
}

