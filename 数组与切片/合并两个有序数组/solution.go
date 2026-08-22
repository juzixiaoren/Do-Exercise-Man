package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	left := m - 1
	right := n - 1
	for i := n + m - 1; i >= 0; i-- {
		if left == -1 {
			nums1[i] = nums2[right]
			right--
		} else if right == -1 {
			nums1[i] = nums1[left]
			left--
		} else if nums1[left] >= nums2[right] {
			nums1[i] = nums1[left]
			left--
		} else {
			nums1[i] = nums2[right]
			right--
		}
	}
}
