package util

func HaveError(err ...error) bool {
	for _, value := range err {
		if value != nil {
			return true
		}
	}
	return false
}
