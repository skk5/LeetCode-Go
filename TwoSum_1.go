package LeetCode_Go

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i := 0; i <len(nums); i++ {
		if j, ok := indexMap[target - nums[i]]; ok {
			return []int{i, j}
		} else {
			indexMap[nums[i]] = i
		}
	}

	return []int{-1, -1}
}