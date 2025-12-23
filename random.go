package wfc

func randomIndex(arr []float64, r float64) int {
	sum := 0.0

	for _, v := range arr {
		sum += v
	}
	
	if sum == 0.0 {
		for i := range arr {
			arr[i] = 1.0
		}
		sum = float64(len(arr))
	}

	for i := range arr {
		arr[i] /= sum
	}

	sum = 0.0

	for i := range arr {
		sum += arr[i]
		if sum >= r {
			return i
		}
	}
	return 0
}
