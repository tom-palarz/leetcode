package submissions

func twoSum(nums []int, target int) []int {
	indices := map[int]int{}

	for i, n := range nums {

		need := target - n

		if _, ok := indices[need]; ok {
			return []int{indices[need], i}
		}

		indices[n] = i

	}

	panic("A valid answer exists, but was not found")
}
