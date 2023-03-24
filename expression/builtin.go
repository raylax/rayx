package expression

import (
	"fmt"
	"math/rand"
	"reflect"
)

type eFunctionToString struct {
}

func (e eFunctionToString) ToString() EString {
	return "ToString"
}

func (e eFunctionToString) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expect one argument, got %d", len(args))
	}
	arg := args[0]

	switch arg.(type) {
	case EInt, EFloat, EBool, EString, EBytes:
		return arg.ToString(), nil
	default:
		return nil, fmt.Errorf("expect EValue, got %s", reflect.TypeOf(arg))
	}

}

type eFunctionToBytes struct {
}

func (e eFunctionToBytes) ToString() EString {
	return "ToBytes"
}

func (e eFunctionToBytes) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expect one argument, got %d", len(args))
	}
	arg := args[0]

	switch arg.(type) {
	case EString:
		return EBytes(arg.(EString)), nil
	default:
		return nil, fmt.Errorf("expect EString, got %s", reflect.TypeOf(arg))
	}
}

type eFunctionRandomInt struct {
}

func (e eFunctionRandomInt) ToString() EString {
	return "RandomInt"
}

func (e eFunctionRandomInt) Call(args []EValue) (EValue, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expect two argument, got %d", len(args))
	}

	min, ok := args[0].(EInt)
	if !ok {
		return nil, fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0]))
	}
	max, ok := args[1].(EInt)
	if !ok {
		return nil, fmt.Errorf("argument 2 expect EInt, got %s", reflect.TypeOf(args[1]))
	}

	return EInt(rand.Intn(int(max-min)+1) + 1), nil
}
