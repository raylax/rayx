package dsl

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const testdataPath = "../_testdata/pocs"

func TestParsePoc(t *testing.T) {
	entries, err := os.ReadDir(testdataPath)
	if err != nil {
		t.Error(err)
	}
	for _, f := range entries {
		t.Run(f.Name(), func(t *testing.T) {
			poc, err := ParsePoc(filepath.Join(testdataPath, f.Name()))
			if err != nil {
				t.Error(err)
			}
			out, err := yaml.Marshal(poc)
			if err != nil {
				t.Error(err)
			}
			newPoc := &Poc{}
			err = yaml.Unmarshal(out, newPoc)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(poc, newPoc) {
				file, _ := ioutil.ReadFile(filepath.Join(testdataPath, f.Name()))
				t.Errorf("=== got \n%s \n===want \n%s", string(out), string(file))
			}
		})
	}
}
