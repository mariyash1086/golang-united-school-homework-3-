package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	sum_elements := float32(0)
	n := len(input)
	j := 0
	for i := 0; i < n; i++ {
		sum_elements += input[i]
		if input[i] > 0 {
			j++

		}
	}
	return float32(sum_elements) / float32(j)
}
