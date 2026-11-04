// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)

// 	// var result [][]int
// 	var resultMap = make(map[[3]int]struct{})

// 	left := 0;
// 	right := len(nums) - 1

// 	for left < right {
// 		sum := nums[left] + nums[right]

// 		for i := left + 1; i < right; i++ {
// 			if sum + nums[i] == 0 {
// 				// resultRow := []int{nums[left], nums[i], nums[right]}
// 				// result = append(result, resultRow)

// 				resultMap[[3]int{nums[left], nums[i], nums[right]}] = struct{}{}
// 				break
// 			}
// 		}

// 		if sum < 0 {
// 			left ++
// 		}

// 		if sum > 0 {
// 			right --
// 		}

// 		if sum == 0 {
// 			left ++
// 			right --
// 		}
// 	}

// 	var result [][]int
// 	for key, _ := range resultMap {
// 		result = append(result, key[:])
// 	}

// 	return result
// }


func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        a := nums[i]
        if a > 0 {
            break
        }
        if i > 0 && a == nums[i-1] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            threeSum := a + nums[l] + nums[r]
            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                res = append(res, []int{a, nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            }
        }
    }

    return res
}