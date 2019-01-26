package util

import "sort"

func IsSliceSame(s1 []int, s2 []int) bool {

	var (
		le1 = len(s1)
		le2 = len(s2)
	)

	if le1 != le2 {
		return false
	}

	sort.Ints(s1)
	sort.Ints(s2)

	for i := 0; i < le1; i++ {

		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

//
func IsSliceHaveSameItem(s []int) bool {

	var (
		m = make(map[int]int)
	)

	for _, value := range s {
		m[value]++

		value := m[value]
		if value >= 2 {
			return true
		}

	}

	return false
}
