package redis_lock

import (
	"fmt"
	"sort"
	"testing"
)

func SortMap(m map[int]interface{}) {
	fmt.Println("len(m):", len(m))
	keys := make([]int, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", m[key])
	}
}

func TestT3(t *testing.T) {

}
