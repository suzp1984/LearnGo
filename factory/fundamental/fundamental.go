package fundamental

import (
	//	"context"
	"errors"
	"reflect"
)

type factoryMethod func(dep interface{}) (interface{}, error)

type Factory interface {
	Create(dep interface{}) (interface{}, error)
}

var (
	typeFactoryRegistry = make(map[reflect.Type]interface{})
)

func RegisterFactory(nilValue interface{}, factory interface{}) error {
	_, ok := factory.(Factory)
	if !ok {
		_, ok := factory.(factoryMethod)
		if !ok {
			_, ok := factory.(func(dep interface{}) (interface{}, error))
			if !ok {
				return errors.New("factory must be one of Factory and typeFactoryRegistry")
			}
		}
	}

	typeType := reflect.TypeOf(nilValue)

	if _, found := typeFactoryRegistry[typeType]; found {
		return errors.New("factory " + typeType.Name() + " already registed.")
	}

	if f, ok := factory.(func(dep interface{}) (interface{}, error)); ok {
		typeFactoryRegistry[typeType] = factoryMethod(f)
	} else {
		typeFactoryRegistry[typeType] = factory
	}
	
	return nil
}

func RegisterFactoryMethod(nilValue interface{}, method factoryMethod) error {
	typeType := reflect.TypeOf(nilValue)

	if _, found := typeFactoryRegistry[typeType]; found {
		return errors.New("factory " + typeType.Name() + " already registed.")
	}

	typeFactoryRegistry[typeType] = method

	return nil 
}

func CreateObject(nilValue interface{}, dep interface{}) (interface{}, error) {
	typeType := reflect.TypeOf(nilValue)
	ins, found := typeFactoryRegistry[typeType]

	if !found {
		return nil, errors.New("factory " + typeType.Name() + " not found.")
	}

	if factory, ok := ins.(Factory); ok {
		return factory.Create(dep)
	}

	if md, ok := ins.(factoryMethod); ok {
		return md(dep)
	}

	if f, ok := ins.(func(dep interface{}) (interface{}, error)); ok {
		return f(dep)
	}

	return nil, errors.New("unrecognized factory type")
	
}
