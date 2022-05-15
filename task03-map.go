package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	//m := map[string]int{"A": 21, "C": 3, "B": 46}

	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	slice := []string{}
	for _, j := range keys {
		slice = append(slice, input[j])

	}

	return slice
}
