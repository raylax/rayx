package expression

import (
	"context"
	"fmt"
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
			expression: "true && false",
			want:       EBool(false),
		},
		{
			expression: "true && true",
			want:       EBool(true),
		},
		{
			expression: "false && false",
			want:       EBool(false),
		},
		{
			expression: "true || true",
			want:       EBool(true),
		},
		{
			expression: "true || false",
			want:       EBool(true),
		},
		{
			expression: "false || true",
			want:       EBool(true),
		},
		{
			expression: "false || false",
			want:       EBool(false),
		},
		{
			expression: "1 + 2",
			want:       EInt(3),
		},
		{
			expression: "1 * 2",
			want:       EInt(2),
		},
		{
			expression: "4 / 2",
			want:       EInt(2),
		},
		{
			expression: "'1' + '2'",
			want:       EString("12"),
		},
		{
			expression: "!true",
			want:       EBool(false),
		},
		{
			expression: "1.0 + 2.0",
			want:       EFloat(3.0),
		},
		{
			expression: `'\n\n\\'`,
			want:       EString("\n\n\\"),
		},
		{
			expression: `b'A\x41'`,
			want:       EBytes([]byte{byte(0x41), byte(0x41)}),
		},
		{
			expression: `bytes('A')`,
			want:       EBytes([]byte{byte(0x41)}),
		},
		{
			expression: `bytes('A') + bytes('B')`,
			want:       EBytes([]byte{byte(0x41), byte(0x42)}),
		},
		{
			expression: `string(1)`,
			want:       EString('1'),
		},
		{
			expression: `'1' + 1`,
			err:        true,
		},
		{
			expression: "string(true)",
			want:       EString("true"),
		},
		{
			expression: "string(bytes('A'))",
			want:       EString("A"),
		},
		{
			expression: "string(1.2)",
			want:       EString("1.2"),
		},
		{
			expression: `string(1, 2)`,
			err:        true,
		},
		{
			expression: `bytes(1, 2)`,
			err:        true,
		},
		{
			expression: `bytes(1)`,
			err:        true,
		},
		{
			expression: `aaaaaaaa(1)`,
			err:        true,
		},
		{
			expression: `bbb(1)`,
			err:        true,
		},
		{
			expression: `'1' + bbb`,
			want:       EString("1bv"),
		},
		{
			expression: `string(string(bbb))`,
			want:       EString("bv"),
		},
		{
			expression: `randomInt(1, 1)`,
			want:       EInt(1),
		},
		{
			expression: `randomInt(999, 999)`,
			want:       EInt(1),
		},
		{
			expression: `randomInt(999)`,
			err:        true,
		},
		{
			expression: `randomInt(999, 'aaa')`,
			err:        true,
		},
		{
			expression: `randomInt('999', 999)`,
			err:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			env := NewEnvironment(context.Background(), vars{
				"bbb": EString("bv"),
			})
			val, err := env.Eval(tt.expression)
			if tt.err && err != nil {
				return
			}
			if tt.err && err == nil {
				t.Errorf("want error - %s", tt.expression)
			}
			if !reflect.DeepEqual(val, tt.want) {
				t.Errorf("want %v, got %v", tt.want, val)
			}
		})
	}
}

type vars map[string]EValue

func (v vars) GetValue(env *Environment, key string) (EValue, error) {
	if value, ok := v[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("'%s' undefined", key)
}
