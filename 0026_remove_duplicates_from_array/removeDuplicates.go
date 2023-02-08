/*
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
*/
// time: 2018-12-20
package removeduplicatesfromarray

func removeDuplicates(nums []int) int {

	ln := len(nums)
	// Return, if array is empty
	// or contains a single element

	if ln <= 1 {
		return ln
	}
	j := 0 // points to  the index of last filled position
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}

	}
	return j + 1

}
