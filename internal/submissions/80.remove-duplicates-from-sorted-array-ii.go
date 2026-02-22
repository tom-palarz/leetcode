package submissions

// more efficient solution
func removeDuplicates2(nums []int) int {

	// since we can have up to 2 duplicates, first 2 will always be
	// correct, even if they are different, so we start at the 3rd
	// element of the array
	k := 2

	// unless it is less than 2 of course, in which case we just return that length
	// and leave the array alone
	if len(nums) < 2 {
		return len(nums)
	}

	for i := 2; i < len(nums); i++ {

		// if current element is different, we have moved onto the
		// next number to see if it is a dup or not. we use a window
		// size of 2 to stay ahead (we'll stay 2 ahead if it is a dup)
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// inefficient solution
/* func removeDuplicates2(nums []int) int {

	count := make(map[int]int, len(nums))

	for _, num := range nums {
		count[num]++
	}

	k := 0

	for _, num := range slices.Sorted(maps.Keys(count)) {
		for i := 1; i <= 2; i++ {
			if count[num] >= i {
				nums[k] = num
				k++
			}
		}
	}

	return k
} */
