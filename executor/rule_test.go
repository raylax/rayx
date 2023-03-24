package executor

import (
	"errors"
	"github.com/raylax/rayx/expression"
	"math/rand"
	"testing"
)

type randomFunction struct {
}

func (r randomFunction) Call(args []expression.EValue) (expression.EValue, error) {
	return expression.EInt(rand.Intn(10000000)), nil
}

func (r randomFunction) ToString() expression.EString {
	return "randomFunction"
}

func TestCachedFunction_Call(t *testing.T) {
	f := cachedFunction{
		f: &randomFunction{},
	}
	value1, err := f.Call(nil)
	if err != nil {
		t.Error(err)
	}
	value2, err := f.Call(nil)
	if err != nil {
		t.Error(err)
	}
	if value1 != value2 {
		t.Errorf("value1 != value2")
	}
}

type errFunction struct {
}

func (r errFunction) Call(args []expression.EValue) (expression.EValue, error) {
	return nil, errors.New("")
}

func (r errFunction) ToString() expression.EString {
	return "errFunction"
}

func TestCachedFunction_CallError(t *testing.T) {
	f := cachedFunction{
		f: &errFunction{},
	}
	_, err1 := f.Call(nil)
	if err1 == nil {
		t.Error(err1)
	}
	_, err2 := f.Call(nil)
	if err2 == nil {
		t.Error(err2)
	}
	if err1 != err2 {
		t.Errorf("err1 != err2")
	}
}
