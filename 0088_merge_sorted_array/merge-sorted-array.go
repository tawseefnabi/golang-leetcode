/*
88. Merge Sorted Array
https://leetcode.com/problems/merge-sorted-array/
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
Note:
 1. The number of elements initialized in nums1 and nums2 are m and n respectively.
 2. You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/
package msa

func msa(nums1 []int, m int, nums2 []int, n int) {
	for i := n + m - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}

	var (
		i = n // pointer for nums1 [n, n+m)
		j = 0 // pointer for nums2 [0, n)
		k = 0 // pointer for merged nums1 [0, n+m)
	)

	for k < n+m {
		if i >= n+m {
			nums1[k] = nums2[j]
			k++
			j++
		} else if j >= n {
			break
			// nums1[k] = nums1[i]
			// k++
			// i++
		} else if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			i++
			k++
		} else { //nums1[i] >= nums2[j]
			nums1[k] = nums2[j]
			k++
			j++
		}
	}

}
