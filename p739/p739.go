package p739

func dailyTemperatures(temperatures []int) []int {
	var result = make([]int, len(temperatures))
	for i := len(result) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] == temperatures[j] {
				if result[j] == 0 {
					result[i] = 0
				} else {
					result[i] = result[j] + j - i
				}
				break
			}
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}
	return result
}
