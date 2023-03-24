package expression

import "fmt"

type Vars interface {
	GetValue(env *Environment, key string) (EValue, error)
}

type MapVars map[string]EValue

func (m MapVars) GetValue(env *Environment, key string) (EValue, error) {
	if value, ok := m[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("'%s' undefined", key)
}

type combinedVars []Vars

func (cv combinedVars) GetValue(env *Environment, key string) (EValue, error) {
	max := len(cv) - 1
	for i := max; i >= 0; i-- {
		value, err := cv[i].GetValue(env, key)
		if err == nil {
			return value, nil
		}
		if i == max {
			return nil, err
		}
	}
	panic("unreached code")
}
