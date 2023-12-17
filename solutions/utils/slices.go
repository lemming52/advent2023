package utils

func Contains(a int, s []int) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}
	return false
}
