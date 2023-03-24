package expression

import (
	"context"
	"github.com/raylax/rayx/dsl"
	"os"
	"path/filepath"
	"testing"
)

const testdataPath = "../_testdata/pocs"

func TestEnvironment_Verify(t *testing.T) {
	entries, err := os.ReadDir(testdataPath)
	if err != nil {
		t.Error(err)
	}
	for _, f := range entries {
		t.Run(f.Name(), func(t *testing.T) {
			env := NewEnvironment(context.Background(), nil)
			poc, err := dsl.ParsePoc(filepath.Join(testdataPath, f.Name()))
			if err != nil {
				t.Error(err)
			}
			err = env.Verify(poc.Expression)
			if err != nil {
				t.Error(poc.Expression, err)
			}

			if poc.Set != nil {
				for _, expr := range *poc.Set {
					err = env.Verify(expr)
					if err != nil {
						t.Error(expr, err)
					}
				}
			}

		})
	}
}
