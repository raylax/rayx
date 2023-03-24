package executor

import (
	"context"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/expression"
	"testing"
)

func TestSet_GetValue(t *testing.T) {
	set, env := newSetEnv(dsl.Set{
		"r1": "1",
		"r2": "2",
		"r3": "randomInt(10000, 20000)",
	})

	value1, err := set.GetValue(env, "r3")
	if err != nil {
		t.Error(err)
	}
	value2, err := set.GetValue(env, "r3")
	if err != nil {
		t.Error(err)
	}
	if value1 != value2 {
		t.Error("value1 != value2")
	}

}

func TestSet_GetValueError(t *testing.T) {
	set, env := newSetEnv(dsl.Set{
		"r1": "balabala",
		"r2": "string(1,2)",
	})

	_, err1 := set.GetValue(env, "r1")
	if err1 == nil {
		t.Error(err1)
	}
	_, err2 := set.GetValue(env, "r1")
	if err2 == nil {
		t.Error(err2)
	}
	if err1 != err2 {
		t.Errorf("err1:%v != err2:%v", err1, err2)
	}

}

func TestSet_GetValueUndefined(t *testing.T) {
	set, env := newSetEnv(dsl.Set{
		"r1": "balabala",
		"r2": "string(1,2)",
	})

	_, err1 := set.GetValue(env, "r1111")
	if err1 == nil {
		t.Error(err1)
	}

}

func newSetEnv(dsl dsl.Set) (*Set, *expression.Environment) {
	return NewSet(dsl), expression.NewEnvironment(context.Background(), nil)
}
