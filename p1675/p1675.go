package p1675

import "sort"

func minimumDeviation(nums []int) int {
	for i := range nums {
		if nums[i]%2 != 0 {
			nums[i] *= 2
		}
	}

	nums = unique(nums)
	dev := nums[len(nums)-1] - nums[0]

	for nums[len(nums)-1]%2 == 0 {
		el := nums[len(nums)-1] / 2
		idx := sort.Search(len(nums), func(i int) bool {
			return nums[i] >= el
		})
		copy(nums[idx+1:], nums[idx:len(nums)-1])
		nums[idx] = el

		newDev := nums[len(nums)-1] - nums[0]
		if newDev < dev {
			dev = newDev
		}
	}

	return dev
}

func unique(in []int) []int {
	sort.Ints(in)
	dstIdx := 0
	for i := 1; i < len(in); i++ {
		if in[i] == in[dstIdx] {
			continue
		}
		dstIdx++
		if dstIdx != i {
			in[dstIdx] = in[i]
		}
	}
	return in[:dstIdx+1]
}
