// +build !linux

package lsmod

const (
	ProcModules = ""
)

func LsMod() (map[string]ModInfo, error) {
	panic("not implemented")
}
