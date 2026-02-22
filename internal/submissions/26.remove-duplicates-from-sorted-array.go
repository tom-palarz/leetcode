package submissions

func removeDuplicates(nums []int) int {

	unique := make([]int, len(nums))

	// first element will always be unique
	unique[0] = nums[0]

	// keep track of how many unique we copy
	ret := 1

	for i, n := range nums {
		// we already copied the first, could alternatively for i :=1....
		if i == 0 {
			continue
		}

		// skip if not unique
		if nums[i] == nums[i-1] {
			continue
		}

		// track unique
		unique[ret] = n
		ret++
	}

	// update input array
	for i := range unique {
		nums[i] = unique[i]
	}

	return ret
}
