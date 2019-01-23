package lsmod

import (
	"fmt"
	"strings"
)

type modState uint8

const (
	StateLive      modState = 1
	StateLoading   modState = 2
	StateUnloading modState = 3
)

func parseState(s string) (modState, error) {
	switch strings.ToLower(s) {
	case "live":
		return StateLive, nil
	case "loading":
		return StateLoading, nil
	case "unloading":
		return StateUnloading, nil
	}

	return 0, fmt.Errorf("unknown state %q", s)
}

func (s modState) String() string {
	switch s {
	case StateLive:
		return "Live"
	case StateLoading:
		return "Loading"
	case StateUnloading:
		return "Unloading"
	}

	panic(fmt.Errorf("unknoiwn state %d", s))
}
