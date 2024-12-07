package helper

func IteratePermutations[T any](src []T, size int) func(yield func([]T) bool) {
	var res = make([]T, size)
	var generate func(int) bool

	return func(yield func([]T) bool) {
		generate = func(idx int) bool {
			if idx == size {
				if !yield(res) {
					return false
				}

				return true
			}

			for _, v := range src {
				res[idx] = v
				if !generate(idx + 1) {
					return false
				}
			}
			return true
		}

		generate(0)
	}
}
