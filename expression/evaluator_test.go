package expression

import (
	"reflect"
	"testing"
)

func Test_evaluator_Visit(t *testing.T) {
	tests := []struct {
		expression string
		want       any
		err        bool
	}{
		{
			expression: "1 + 2",
			want:       3,
		},
		{
			expression: "'1' + '2'",
			want:       "12",
		},
		{
			expression: `'\n\n\\'`,
			want:       "\n\n\\",
		},
		{
			expression: `b'A\x41'`,
			want:       []byte{byte(0x41), byte(0x41)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			env := NewEnvironment()
			val, err := env.Eval(tt.expression)
			if tt.err && err == nil {
				t.Errorf("want error - %s", tt.expression)
			}
			if !reflect.DeepEqual(val, tt.want) {
				t.Errorf("want %v, got %v", tt.want, val)
			}
		})
	}
}
