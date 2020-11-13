package main

import "fmt"

var arr1 = []int{1, 2, 3, 0, 0, 0}
var arr2 = []int{2, 5, 6}

func merge(nums1 []int, m int, nums2 []int, n int) {
	ind, ind1, ind2 := m+n-1, m-1, n-1
	for ind1 >= 0 && ind2 >= 0 {

		if nums1[ind1] > nums2[ind2] {
			nums1[ind] = nums1[ind1]
			ind1--
		} else {
			nums1[ind] = nums2[ind2]
			ind2--
		}
		ind--
	}
	for ; ind2 >= 0; ind2-- {
		nums1[ind2] = nums2[ind2]
	}

}

func main() {
	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1)
}
