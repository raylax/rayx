package dsl

import (
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Raw = map[string]any

type Poc struct {
	Name       string  `yaml:"name"`
	Manual     bool    `yaml:"manual"`
	Transport  string  `yaml:"transport"`
	Set        *Set    `yaml:"set"`
	Rules      *Rules  `yaml:"rules"`
	Expression string  `yaml:"expression"`
	Detail     *Detail `yaml:"detail"`
}

type Detail struct {
	Author      string   `yaml:"author,omitempty"`
	Links       []string `yaml:"links,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Info        string   `yaml:"info,omitempty"`
	raw         *Raw
}

func (d *Detail) MarshalYAML() (interface{}, error) {
	return d.raw, nil
}

func (d *Detail) UnmarshalYAML(value *yaml.Node) error {
	err := value.Decode(&d.raw)
	if err != nil {
		return err
	}
	return mapstructure.Decode(d.raw, d)
}

type Set map[string]string

type Rules map[string]Rule

type Header map[string]string

type Rule struct {
	Request    *Request `yaml:"request"`
	Expression string   `yaml:"expression"`
	Output     *Output  `yaml:"output,omitempty"`
}

type Request struct {
	Cache           bool   `yaml:"cache"`
	Method          string `yaml:"method"`
	Path            string `yaml:"path"`
	FollowRedirects *bool  `yaml:"follow_redirects"`
	Headers         Header `yaml:"headers,omitempty"`
	Body            string `yaml:"body"`
}

type Output map[string]string

func ParsePoc(path string) (*Poc, error) {
	poc := &Poc{}
	err := unmarshal(path, poc)
	if err != nil {
		return nil, err
	}
	return poc, nil
}

func unmarshal[T any](path string, out *T) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, out)
}
