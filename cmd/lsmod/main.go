package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	lsmod "github.com/Djarvur/go-lsmod"
)

func main() {
	mods, err := lsmod.LsMod()
	if err != nil {
		panic(err)
	}

	var (
		longestName = 0
		longestSize = 0
		longestInst = 0
		names       = make([]string, 0, len(mods))
	)

	for name, info := range mods {
		names = append(names, name)

		if len(name) > longestName {
			longestName = len(name)
		}

		numStr := strconv.Itoa(int(info.Mem))
		if len(numStr) > longestSize {
			longestSize = len(numStr)
		}

		numStr = strconv.Itoa(int(info.Instances))
		if len(numStr) > longestInst {
			longestInst = len(numStr)
		}
	}

	sort.Strings(names)

	fmtStr := fmt.Sprintf(
		"%%-%ds %%%dd %%%dd %%s\n",
		longestName,
		longestSize,
		longestInst,
	)

	for _, name := range names {
		info := mods[name]
		fmt.Printf(fmtStr, name, info.Mem, info.Instances, strings.Join(info.Depends, ","))
	}
}
