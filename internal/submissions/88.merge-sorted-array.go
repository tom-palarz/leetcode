package submissions

/* O(m+n) time complexity and O(1) space complexity solution */

func merge(nums1 []int, m int, nums2 []int, n int) {

	// Since the end of nums1 will be zeroes (for as long as nums2 is)
	// and the arrays are already sorted, let's start from the back and move up
	i := m - 1     // end of nums1 slice before the extra zeroes
	j := n - 1     // end of nums2 slice
	k := m + n - 1 // end of full nums1 array that has the zeroes

	// Merge from the end while there are numbers to insert
	for j >= 0 {

		// if nums1 has elements remaining and the current element
		// is bigger than end of nums2
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--

			// else, the current element from the end of nums2
			// is bigger and should be inserted in the tail
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k-- // Move back in the merged array
	}

}
