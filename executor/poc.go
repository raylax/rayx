package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/raylax/rayx/cmd"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
	"github.com/raylax/rayx/transport"
)

type State string

const (
	StateError State = "E"
	StateHit   State = "H"
	StateMISS  State = "M"
)

type PocExecutor struct {
}

func (e *PocExecutor) Execute(config *cmd.Config, dsl dsl.Poc) (State, error) {
	if dsl.Transport != "http" {
		return StateError, errors.New("unsupported transport " + dsl.Transport)
	}

	set := NewSet(*dsl.Set)
	env := expression.NewEnvironment(context.Background(), set)
	t := transport.NewHttpTransport(env, config.Url, config.Headers, config.Cookies)

	var rules = Rules{}
	for name, rule := range *dsl.Rules {
		f := &httpRequestFunc{
			expression: rule.Expression,
			d:          rule.Request,
			env:        env,
			t:          t,
		}
		if rule.Request.Cache {
			rules[name] = &cachedFunction{
				f: f,
			}
		} else {
			rules[name] = f
		}
	}

	value, err := env.EvalWithVars(dsl.Expression, rules)
	if err != nil {
		return StateError, err
	}

	switch value.(type) {
	case expression.EBool:
		{
			if value.(expression.EBool) {
				return StateHit, nil
			} else {
				return StateMISS, nil
			}
		}
	default:
		return StateError, fmt.Errorf("expect TBool, got %v", value)
	}
}

type httpRequestFunc struct {
	d          *dsl.Request
	expression string
	env        *expression.Environment
	t          *transport.HttpTransport
}

func (h *httpRequestFunc) ToString() expression.EString {
	return "httpRequestFunc"
}

func (h *httpRequestFunc) Call(args []expression.EValue) (expression.EValue, error) {
	response, err := h.t.Do(h.env.Context, h.d)
	if err != nil {
		return nil, err
	}

	value, err := h.env.EvalWithVars(h.expression, expression.MapVars{
		"response": response,
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}
