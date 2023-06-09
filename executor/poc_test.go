package executor

import (
	"fmt"
	"github.com/raylax/rayx/dsl"
	"testing"
)

func TestPocExecutor_ExecuteUnsupportedTransport(t *testing.T) {
	executor := &PocExecutor{}
	config := &Config{}

	d := &dsl.Poc{
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

	tests := []struct {
		config *Config
		d      dsl.Poc
		state  State
		err    bool
	}{
		{
			config: &Config{
				Timeout: 30,
				Url:     "https://httpbin.org",
				Headers: nil,
				Cookies: nil,
			},
			d: dsl.Poc{
				Transport: "http",
				Set: &dsl.Set{
					"r1": "base64('aaa/getaaaa')",
				},
				Rules: &dsl.Rules{
					"r1": dsl.Rule{
						Request: &dsl.Request{
							Path: "/base64/{{r1}}",
						},
						Expression: `response.status == 200 && response.body.bcontains(bytes('aaa/getaaaa'))
&& response.headers['content-type'] == 'text/html; charset=utf-8'
&& response.headers['content_type'] == 'text/html; charset=utf-8'
&& response.headers['Content-Type'] == 'text/html; charset=utf-8'
&& response.content_type == 'text/html; charset=utf-8'
`,
						Output: &dsl.Output{
							"search": `'aaa(?P<method>.*?)aaaa'.bsubmatch(response.body)`,
							"method": `search['method']`,
						},
					},
					"r2": dsl.Rule{
						Request: &dsl.Request{
							Path: "{{r1}}",
						},
						Expression: `response.status == 200 && response.body.bcontains(bytes('"url": "https://httpbin.org/get"'))`,
					},
				},
				Expression: "r1()",
			},
			state: StateHit,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			executor := &PocExecutor{}
			state, err := executor.Execute(tt.config, &tt.d)
			if tt.err && err == nil {
				t.Errorf("expect error")
			}
			if !tt.err && tt.state != state {
				t.Errorf("expect %s, got %s - %v", tt.state, state, err)
			}
		})
	}

}
