package lsmod

// ModInfo contains a module info provided by /proc/modules
type ModInfo struct {
	Depends   []string
	Mem       uint64
	Instances uint64
	Offset    uint64
	Tained    modTained
	State     modState
}
