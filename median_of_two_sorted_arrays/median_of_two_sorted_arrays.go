package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)
	return median(nums)
}

func merge(nums1 []int, nums2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}
	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}
	return result
}

func median(nums []int) float64 {
	n := len(nums)
	if n%2 == 0 {
		return float64(nums[n/2-1]+nums[n/2]) / 2
	}
	return float64(nums[n/2])
}
