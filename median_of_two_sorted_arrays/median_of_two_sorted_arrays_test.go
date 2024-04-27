package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := 2.0
	result := findMedianSortedArrays(nums1, nums2)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
