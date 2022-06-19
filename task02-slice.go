package homework

func reverse(input []int64) (result []int64) {
	n := len(input)
	for i := n-1; i>=0; i--{
		result = append(result, input[i])
	}
	return
}