package executor

import (
	"fmt"
	"github.com/raylax/rayx/expression"
)

type Rules map[string]expression.EFunction

func (r Rules) GetValue(env *expression.Environment, key string) (expression.EValue, error) {
	if rule, ok := r[key]; ok {
		return rule, nil
	}
	return nil, fmt.Errorf("'%s' undefined", key)
}

type cachedFunction struct {
	f        expression.EFunction
	computed bool
	value    expression.EValue
	err      error
}

func (c *cachedFunction) ToString() expression.EString {
	return "cachedFunction"
}

func (c *cachedFunction) Call(args []expression.EValue) (expression.EValue, error) {
	if c.computed {
		return c.value, c.err
	}
	value, err := c.f.Call(args)
	c.computed = true
	if err != nil {
		c.err = err
		return nil, err
	} else {
		c.value = value
		return value, nil
	}
}
