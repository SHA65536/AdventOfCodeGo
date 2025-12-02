package helper

func AllEqual[T comparable](input ...T) bool {
	for _, e := range input {
		if e != input[0] {
			return false
		}
	}
	return true
}
