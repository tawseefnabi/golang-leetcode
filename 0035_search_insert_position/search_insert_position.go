/*
35. Search Insert Position
https://leetcode.com/problems/search-insert-position/
Given a sorted array and a target value,
return the index if the target is found.
If not, return the index where it would be if it were inserted in order.
You may assume no duplicates in the array.
*/
package searchinsertposition

func SearchInsertPosition(nums []int, target int) int {
	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end + 1) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
