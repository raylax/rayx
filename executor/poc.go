package executor

import (
	"errors"
	"github.com/raylax/rayx/cmd"
	"github.com/raylax/rayx/dsl"
)

type State string

const (
	StateError State = "E"
	StateHit   State = "H"
	StateMISS  State = "M"
)

type PocExecutor struct {
}

func (e *PocExecutor) Execute(config *cmd.Config, dsl dsl.Poc) (State, error) {
	if dsl.Transport != "http" {
		return StateError, errors.New("unsupported transport " + dsl.Transport)
	}

	return StateMISS, nil
}
