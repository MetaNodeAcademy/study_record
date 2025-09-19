package twonumberssum

func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		sum := target - v
		for j, v2 := range nums {
			if j != i && v2 == sum {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
