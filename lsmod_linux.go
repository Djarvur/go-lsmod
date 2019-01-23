package lsmod

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	ProcModules = "/proc/modules"

	minFieldsPerLine = 6
	maxFieldsPerLine = 6
	noDeps           = "-"
	delimDeps        = ","
)

func LsMod() (map[string]ModInfo, error) {
	file, err := os.Open(ProcModules)
	if err != nil {
		return nil, errors.Wrapf(err, "Error opening %q", ProcModules)
	}
	defer file.Close()

	mods := make(map[string]ModInfo)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) < minFieldsPerLine || len(fields) > maxFieldsPerLine {
			return nil, fmt.Errorf("invalid input line %q", line)
		}

		info, err := parseInfo(fields)
		if err != nil {
			return nil, errors.Wrapf(err, "error parsing %q", ProcModules)
		}

		mods[fields[0]] = info
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading %q", ProcModules)
	}

	return mods, nil
}

func parseInfo(fields []string) (info ModInfo, err error) {
	info.Mem, err = strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return info, errors.Wrap(err, "invalid mem (field 2)")
	}
	instances, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return info, errors.Wrap(err, "invalid instances (field 3)")
	}
	info.Instances = instances
	info.Depends = splitDeps(fields[3])
	info.State, err = parseState(fields[4])
	if err != nil {
		return info, errors.Wrap(err, "unknown state (field 5)")
	}
	info.Offset, err = strconv.ParseUint(fields[5], 16, 64)
	if err != nil {
		return info, errors.Wrap(err, "invalid offset (field 6)")
	}
	if len(fields) == maxFieldsPerLine {
		info.Tained, err = parseTained(fields[6])
		if err != nil {
			return info, errors.Wrap(err, "unknown tained (field 7)")
		}
	}

	return info, nil
}

func splitDeps(line string) []string {
	if line == noDeps {
		return nil
	}

	return strings.Split(strings.TrimRight(line, delimDeps), delimDeps)
}
