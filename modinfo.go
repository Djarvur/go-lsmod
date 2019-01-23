package lsmod

type ModInfo struct {
	Depends   []string
	Mem       uint64
	Instances uint64
	Offset    uint64
	Tained    modTained
	State     modState
}
