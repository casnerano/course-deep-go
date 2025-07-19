package functions

func Map[T1 any, T2 any](data []T1, action func(T1) T2) []T2 {
	if data == nil {
		return nil
	}

	result := make([]T2, len(data))
	for i := range data {
		result[i] = action(data[i])
	}

	return result
}

func Filter[T any](data []T, action func(T) bool) []T {
	if data == nil {
		return nil
	}

	result := make([]T, 0)
	for _, v := range data {
		if action(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[T any](data []T, initial T, action func(T, T) T) T {
	result := initial
	for _, v := range data {
		result = action(result, v)
	}
	return result
}
