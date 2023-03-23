package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

type mapValue struct {
	value   *map[string]string
	changed bool
}

func newMapValue() mapValue {
	value := map[string]string{}
	return mapValue{value: &value}
}

func (s *mapValue) Set(val string) error {
	var ss []string
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as key=value", val)
	case 1:
		ss = append(ss, strings.Trim(val, `"`))
	default:
		r := csv.NewReader(strings.NewReader(val))
		var err error
		ss, err = r.Read()
		if err != nil {
			return err
		}
	}

	out := make(map[string]string, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}
		out[kv[0]] = kv[1]
	}
	if !s.changed {
		*s.value = out
	} else {
		for k, v := range out {
			(*s.value)[k] = v
		}
	}
	s.changed = true
	return nil
}

func (s *mapValue) Type() string {
	return ""
}

func (s *mapValue) String() string {
	if len(*s.value) == 0 {
		return ""
	}
	records := make([]string, 0, len(*s.value)>>1)
	for k, v := range *s.value {
		records = append(records, k+"="+v)
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return "[" + strings.TrimSpace(buf.String()) + "]"
}
