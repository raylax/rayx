package executor

import (
	"fmt"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
)

type Set struct {
	dsl   dsl.Set
	cache map[string]setItem
}

func NewSet(dsl dsl.Set) *Set {
	return &Set{
		dsl:   dsl,
		cache: map[string]setItem{},
	}
}

func (s *Set) GetValue(env *expression.Environment, key string) (expression.EValue, error) {
	if c, ok := s.cache[key]; ok {
		if c.err == nil {
			return c.value, nil
		} else {
			return nil, c.err
		}
	}
	if expr, ok := s.dsl[key]; ok {
		value, err := env.Eval(expr)
		if err == nil {
			s.cache[key] = setItem{
				value: value,
				err:   nil,
			}
			return value, nil
		} else {
			s.cache[key] = setItem{
				value: nil,
				err:   err,
			}
			return nil, err
		}
	}

	return nil, fmt.Errorf("'%s' undefined", key)
}

type setItem struct {
	value expression.EValue
	err   error
}
