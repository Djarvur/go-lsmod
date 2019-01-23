// +build !linux

package lsmod

// LsMod is a function reading and parsing /proc/modules pseudo-file
func LsMod() (map[string]ModInfo, error) {
	panic("not implemented")
}
