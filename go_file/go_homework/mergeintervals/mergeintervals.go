package mergeintervals

func MergeIntervals(intervals [][]int) [][]int {
	temp := []int{}
	for i := 0; i < len(intervals)-1; i++ {
		if (intervals[i][0] <= intervals[i+1][1] && intervals[i][0] >= intervals[i+1][0]) ||
			(intervals[i][1] <= intervals[i+1][1] && intervals[i][1] >= intervals[i+1][0]) {
			temp = intervals[i]
			// 前一个区间的最小值是否包含在下一个区间中
			if intervals[i][0] <= intervals[i+1][1] && intervals[i][0] >= intervals[i+1][0] {
				temp[0] = intervals[i+1][0]
			} else {
				temp[0] = intervals[i][0]
			}
			// 前一个区间的最大值是否包含在下一个区间中
			if intervals[i][1] <= intervals[i+1][1] && intervals[i][1] >= intervals[i+1][0] {
				temp[1] = intervals[i+1][1]
			} else {
				temp[1] = intervals[i][1]
			}
			intervals = append(intervals[:i], intervals[i+1:]...)
			intervals[i] = temp
			i--
		}
	}
	return intervals
}
