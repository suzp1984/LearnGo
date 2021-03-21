package app

import (
	"fmt"
	"errors"
	"example.com/practise/fundamental"
)

type Md struct {
	num int
}

func createMd(dep interface{}) (interface{}, error) {
	fmt.Println("create md")
	
	if num, ok := dep.(int); ok {
		return Md { num }, nil
	}

	return nil, errors.New("dep type error, need int.")
}

func init() {
	if err := fundamental.RegisterFactory((*Md)(nil), createMd); err != nil {
		fmt.Println("err ", err)
	}
	
}
