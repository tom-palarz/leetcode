package submissions

// more efficient solution doing the actual work in place
func removeDuplicates(nums []int) int {

	// the first element will always be unique, so count that one
	// and start at the next one
	k := 1

	for i := 1; i < len(nums); i++ {

		// if the current one we are looking at is unique, copy it
		// into the unique placeholder position k and move the unique
		// pointer up
		if nums[i] != nums[k] {
			nums[k] = nums[i]
			k++
		}
	}

	// return the unique placeholder count
	return k
}

/* func removeDuplicates(nums []int) int {

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
} */
