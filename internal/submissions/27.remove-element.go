package submissions

func removeElement(nums []int, val int) int {

	removed := 0

	for i, n := range nums {
		if n == val {
			// if we've hit our placeholders, we're done
			if n == -100 {
				break
			}

			for nums[i] == val {
				// swap with the end and mark it as removed with -100
				nums[i] = nums[len(nums)-removed-1]
				nums[len(nums)-removed-1] = -100

				removed++
			}
		}
	}

	return len(nums) - removed
}
