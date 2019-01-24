package lsmod

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var testData = map[string]ModInfo{
	"mod1": ModInfo{
		Depends: []string{
			"mod2",
			"mod3",
		},
		Mem:       1,
		Instances: 1,
		Offset:    1,
		State:     StateLive,
	},
	"mod2": ModInfo{
		Depends: []string{
			"mod3",
		},
		Mem:       2,
		Instances: 1,
		Offset:    2,
		State:     StateLive,
	},
	"mod3": ModInfo{
		Mem:       3,
		Instances: 1,
		Offset:    3,
		Tained:    TainedF,
		State:     StateLive,
	},
	"mod4": ModInfo{
		Mem:    4,
		Offset: 4,
		State:  StateLive,
	},
	"mod5": ModInfo{
		Mem:       5,
		Instances: 1,
		Offset:    5,
		State:     StateUnloading,
	},
}

func TestParse(t *testing.T) {
	mods, err := parse("testdata.txt")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testData, mods) {
		t.Error(fmt.Errorf("hardcoded test data not equal to parsed\nhardcoded:\n%s\nparsed:\n%s\n", serializeMap(testData), serializeMap(mods)))
	}
}

func serializeMap(o map[string]ModInfo) string {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	return string(b)
}
