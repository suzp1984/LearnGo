package main

import (
	"fmt"
	"context"
	"example.com/practise/fundamental"
	"example.com/practise/app"
)

func whatSoEver(dep interface{}) (interface{}, error) {
	fmt.Println("what so ever function call")

	return nil, nil
}

type FunType func(dep interface{}) (interface{}, error)

func main() {
	ap, err := fundamental.CreateObject((*app.App)(nil), "new")

	if err == nil {
		fmt.Println("created App is ", ap)

		app, ok := ap.(app.App)

		fmt.Println("created App = ", app, "; ok = ", ok)
	}

	ctx, err := fundamental.CreateObject((*context.Context)(nil), nil)

	if err == nil {
		if ctx, ok := ctx.(context.Context); ok {
			fmt.Println("created Context is ", ctx)
		}
		
	}

	if md, err := fundamental.CreateObject((*app.Md)(nil), 8); err == nil {
		fmt.Println("Created Md is ", md)
	}
	

	var t interface{}

	t = FunType(whatSoEver)

	// f, _ := t.(func(dep interface{}) (interface{}, error))
	f, _ := t.(FunType)

	if f == nil {
		fmt.Println("f is nil")
	}
	
	f(nil)
}
