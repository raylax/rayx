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

	ctx := context.Background()

	set := NewSet(*dsl.Set)

	outputs := &expression.MapVars{}
	env := expression.NewEnvironment(ctx, set, outputs)
	t := transport.NewHttpTransport(env, config.Url, config.Headers, config.Cookies)

	var rules = Rules{}
	for name, rule := range *dsl.Rules {
		f := &httpRequestFunc{
			expression: rule.Expression,
			output:     rule.Output,
			outputs:    outputs,
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

	env = expression.NewEnvironment(ctx, set)
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
	output     *dsl.Output
	outputs    *expression.MapVars
	expression string
	env        *expression.Environment
	t          *transport.HttpTransport
}

func (h *httpRequestFunc) ToString() expression.EString {
	return "httpRequestFunc"
}

func (h *httpRequestFunc) Call(args []expression.EValue) (expression.EValue, error) {
	// 发起request
	response, err := h.t.Do(h.env.Context, h.d)
	if err != nil {
		return nil, err
	}

	vars := expression.MapVars{"response": response}

	value, err := h.env.EvalWithVars(h.expression, vars)
	if err != nil {
		return nil, err
	}

	// 计算output
	if h.output != nil {
		for key, expr := range *h.output {
			value, err := h.env.EvalWithVars(expr, vars)
			if err != nil {
				return nil, err
			}
			(*h.outputs)[key] = value
		}
	}

	return value, nil
}
