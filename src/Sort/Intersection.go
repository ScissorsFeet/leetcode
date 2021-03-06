package Sort

func intersection(nums1 []int, nums2 []int) []int {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	if nums1Len == 0 || nums2Len == 0 {
		return []int{}
	}
	var ret []int
	var baseArr []int
	var compArr []int
	if nums2Len > nums1Len {
		baseArr = nums1
		compArr = nums2
	} else {
		baseArr = nums2
		compArr = nums1
	}
	m := make(map[int]rune)
	for i := range baseArr {
		if _, ok := m[baseArr[i]]; !ok {
			m[baseArr[i]] = 1
		}
	}
	for i := range compArr {
		if _, ok := m[compArr[i]]; ok {
			ret = append(ret, compArr[i])
			delete(m, compArr[i])
		}
	}
	return ret
}
