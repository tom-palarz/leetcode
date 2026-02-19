package submissions

func twoSum(nums []int, target int) []int {
	needs := map[int]int{}
	indices := map[int]int{}

	for i, n := range nums {
		// target positive
		//if target >= 0 {
		// can't be in sum if bigger
		//if n <= target {
		// check if we've already found the pair
		//fmt.Printf("\n--------------------------------\n")
		if value, ok := needs[n]; ok {
			//fmt.Printf("Found needs[%d] = %d\n", n, value)
			if n+value == target {
				if _, ok := indices[value]; ok {
					return []int{indices[value], i}
				}
			}
		} /* else {
		    fmt.Printf("%d not found in %+v\n", n, needs)
		}*/

		// else process in maps
		partner := target - n
		needs[partner] = n
		//needs[n] = partner
		indices[n] = i

		//fmt.Printf("After processing - needs: %+v\n", needs)
		//fmt.Printf("After processing - indices: %+v\n", indices)
		//}
		// else skip
		/*} else {
		    // can't be in sum if more negative than negative target
		    //if n >= target {
		        // check if we've already found the pair
		        if value, ok := needs[n]; ok {
		            return []int{indices[value], i}
		        }

		        // else process maps
		        partner := target - n
		        needs[partner] = n
		        needs[n] = partner
		        indices[n] = i
		    //}
		}*/
	} // end for

	panic("A valid answer exists, but was not found")
}
