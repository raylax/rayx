package executor

import (
	"github.com/raylax/rayx/cmd"
	"github.com/raylax/rayx/dsl"
	"testing"
)

func TestPocExecutor_ExecuteUnsupportedTransport(t *testing.T) {
	executor := &PocExecutor{}
	config := &cmd.Config{}

	d := dsl.Poc{
		Transport:  "ssh",
		Set:        nil,
		Rules:      nil,
		Expression: "",
		Detail:     nil,
	}
	_, err := executor.Execute(config, d)
	if err == nil {
		t.Error("expect error")
	}
}

func TestPocExecutor_Execute(t *testing.T) {
	executor := &PocExecutor{}
	config := &cmd.Config{
		Url:     "https://httpbin.org/get",
		Headers: nil,
		Cookies: nil,
	}

	d := dsl.Poc{
		Transport: "http",
		Set: &dsl.Set{
			"r1": "'/get'",
		},
		Rules: &dsl.Rules{
			"r1": dsl.Rule{
				Request: &dsl.Request{
					Path: "{{r1}}",
				},
				Expression: "response.status == 200",
			},
		},
		Expression: "r1()",
	}
	state, err := executor.Execute(config, d)
	if err != nil {
		t.Error(err)
	}
	println(state)
}
