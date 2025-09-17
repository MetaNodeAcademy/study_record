package thenumberthatappearsonlyonce

func SingleNumber(nums []int) int {
	nummap := map[int]int{}
	for _, v := range nums {
		nummap[v] += 1
	}
	for i, v := range nummap {
		if v == 1 {
			return i
		}
	}
	return 0
}
