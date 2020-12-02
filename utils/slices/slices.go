package slices

func StringFilter(slice []string, filter func(string) bool) []string {
	resSlice := make([]string, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}

func IntFilter(slice []int, filter func(int) bool) []int {
	resSlice := make([]int, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}
