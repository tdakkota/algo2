package cache

import (
	"encoding/json"
	"strings"
)

type testCall struct {
	t      string
	args   [2]int
	result int
}

func unmarshalString(s string, i interface{}) {
	err := json.Unmarshal([]byte(s), i)
	if err != nil {
		panic(err)
	}
}

func parseTest(test string) []testCall {
	split := strings.Split(test, "\n")
	var names []string
	var args [][2]int
	var result []int
	unmarshalString(split[0], &names)
	unmarshalString(split[1], &args)
	unmarshalString(split[2], &result)

	n := make([]testCall, len(names))
	for i := range n {
		n[i] = testCall{names[i], args[i], result[i]}
	}
	return n
}
