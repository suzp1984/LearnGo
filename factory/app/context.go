package app

import (
	"context"
	"example.com/practise/fundamental"	
)

type ContextFactory struct {
}

func (f ContextFactory) Create(dep interface{}) (interface{}, error) {
	return context.Background(), nil
}

func init() {
	fundamental.RegisterFactory((*context.Context)(nil), ContextFactory {})
}
