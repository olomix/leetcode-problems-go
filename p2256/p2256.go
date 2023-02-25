package p2256

func minimumAverageDifference(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	if total == 0 || len(nums) == 0 {
		return 0
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	rightAvg := func(leftAcc int, idx int) int {
		right := 0
		rightLn := len(nums) - idx - 1
		if len(nums)-idx-1 > 0 {
			right = (total - leftAcc) / rightLn
		}
		return right
	}

	acc := nums[0]
	minIdx := 0
	right := rightAvg(acc, 0)
	minVal := abs(acc - right)

	for i := 1; i < len(nums); i++ {
		acc += nums[i]
		right = rightAvg(acc, i)
		val := abs(acc/(i+1) - rightAvg(acc, i))
		if val < minVal {
			minVal = val
			minIdx = i
		}
	}

	return minIdx
}

func minimumAverageDifference2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = nums[i] + arr[i-1]
	}
	//log.Println(nums)
	//log.Println(arr)
	for i := 0; i < len(nums)-1; i++ {
		leftN := i + 1
		rightN := len(nums) - i - 1
		leftSum := arr[i]
		rightSum := arr[len(nums)-1] - leftSum
		arr[i] = leftSum/leftN - rightSum/rightN
		if arr[i] < 0 {
			arr[i] = -arr[i]
		}
	}
	arr[len(nums)-1] = arr[len(nums)-1] / len(nums)

	min := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}

	return min
}
