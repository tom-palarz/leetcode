package submissions

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	ptr := &nums1
	if m == 0 {
		(*ptr) = append([]int(nil), nums2...)
		fmt.Printf("m = 0; nums1: %v\n", nums1)
		return
	}

	j := 0
	for i := 0; i < m+n; i++ {

		fmt.Printf("\n--------------\n")
		fmt.Printf("i: %d, j: %d\n", i, j)

		if nums2[j] <= nums1[i] || (nums1[i] == 0 && i > m) {
			fmt.Printf("======\ninsert\n")
			fmt.Printf("i: %d, j: %d\n", i, j)
			fmt.Printf("BEFORE nums1: %v ; nums2: %v\n", nums1, nums2)
			//nums1 = slices.Concat(nums1[0:i], []int{nums2[j]}, nums1[i:m+n-1])
			//tail := nums1[i+1:m+n-1]
			var tail = make([]int, m+n-i-1)
			copy(tail, nums1[i:m+n-1])
			nums1 = append(nums1[0:i], nums2[j])
			nums1 = append(nums1, tail...)
			fmt.Printf("AFTER nums1: %v ; nums2: %v\n", nums1, nums2)
			//i++
			j++
			fmt.Printf("NEW i: %d, j: %d\n", i, j)
			if j == n {
				break
			}
		}

	}
	fmt.Printf("End For Loop\n")
	fmt.Printf("nums1: %v\n", nums1)
}
