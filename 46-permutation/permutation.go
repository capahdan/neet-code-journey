package permutation

func permute(nums []int) [][]int {
	allPermutations := [][]int{}

	var buildPermutation func(pickedNums []int, unpickedNums []int)

	buildPermutation = func(pickedNums []int, unpickedNums []int) {
		// Base case: no more numbers to pick → we have a complete permutation
		if len(unpickedNums) == 0 {
			// Save a copy of the completed permutation
			completed := make([]int, len(pickedNums))
			copy(completed, pickedNums)
			allPermutations = append(allPermutations, completed)
			return
		}

		// Try picking each unpicked number one at a time
		for i, candidate := range unpickedNums {
			// Remove candidate from unpicked list
			remaining := append(append([]int{}, unpickedNums[:i]...), unpickedNums[i+1:]...)

			// Add candidate to our current permutation and go deeper
			buildPermutation(append(pickedNums, candidate), remaining)
		}
	}

	buildPermutation([]int{}, nums)
	return allPermutations
}

/*
nums = [1, 2, 3]

Pick 1 → remaining [2,3]
    Pick 2 → remaining [3]
        Pick 3 → remaining [] → DONE ✅ [1,2,3]
    Pick 3 → remaining [2]
        Pick 2 → remaining [] → DONE ✅ [1,3,2]

Pick 2 → remaining [1,3]
    Pick 1 → remaining [3]
        Pick 3 → remaining [] → DONE ✅ [2,1,3]
    ...and so on

*/
