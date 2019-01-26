package util

//s中如果有重复， 就移除前面的。
func RemoveSliceItem(s []int, item int) []int {

	var (
		le    = len(s)
		index int
	)

	//if le == 0 {
	//	return s
	//}

	for index1, value := range s {

		if value == item {
			index = index1
			break
		}
	}

	switch {

	case index == le-1 && index == 0:
		return []int{}
	case index == le-1:
		return s[:le-1]

	case index == 0:
		return s[1:]

	default:
		return append(s[:index], s[index+1:]...)

	}

}

func RemoveSliceItem2(s []int, item int) []int {

	var (
		le    = len(s)
		index int
		s2    []int
	)

	s2 = make([]int, le)
	copy(s2, s)

	//if le == 0 {
	//	return s
	//}

	for index1, value := range s {

		if value == item {
			index = index1
			break
		}
	}

	switch {

	case index == le-1 && index == 0:
		return []int{}
	case index == le-1:
		return s2[:le-1]

	case index == 0:
		return s2[1:]

	default:
		return append(s2[:index], s2[index+1:]...)

	}

}

func IsSliceHaveItem(s []int, item int) bool {
	for _, value := range s {
		if value == item {
			return true
		}
	}
	return false
}
