package homework

func reverse(input []int64) (result []int64) {

	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i] < input[j] {
				el1 := input[i]
				el2 := input[j]

				input[i] = el2
				input[j] = el1

			}
		}
	}
	return input

}
