package LeetCode_Go

func findMedianSortedArrays(nums1 []int, nums2 [] int) float64 {
	nums := make([]int, len(nums1) + len(nums2))

	idx1, idx2 := 0, 0
	i := 0
	for ; idx1 < len(nums1) && idx2 < len(nums2); i++ {
		if nums1[idx1] < nums2[idx2] {
			nums[i] = nums1[idx1]
			idx1++
		} else {
			nums[i] = nums2[idx2]
			idx2++
		}
	}

	
}