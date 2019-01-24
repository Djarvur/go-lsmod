// +build !linux

package lsmod

import "runtime"

// LsMod is a function reading and parsing /proc/modules pseudo-file
func LsMod() (map[string]ModInfo, error) {
	if runtime.GOOS != "linux" {
		panic("not implemented")
	}

	return parse()
}
