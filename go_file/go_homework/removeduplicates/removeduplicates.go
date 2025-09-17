package removeduplicates

func RemoveDuplicates(nums []int) int {
	nummap := map[int]int{}
	num := 1
	for _, v := range nums {
		if nummap[v] == 0 {
			nummap[v] = num
			num++
		}
	}
	for i, v := range nummap {
		nums[v-1] = i
	}
	return len(nummap)
}
