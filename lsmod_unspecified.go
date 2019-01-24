// +build !linux

package lsmod

import "runtime"

// LsMod is a function reading and parsing /proc/modules pseudo-file
func LsMod() (map[string]ModInfo, error) {
	if runtime.GOOS == "linux" {
		return parse(ProcModules) // just to stop linter comlaining about unused code on !linux
	}

	panic("not implemented")
}
