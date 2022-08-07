package common

func ConvType2Any[T any](srcSlice []T) []any {
	result := make([]any, len(srcSlice))
	for i, value := range srcSlice {
		result[i] = value
	}
	return result
}

func ConvAny2Type[T any](srcSlice []any) []T {
	result := make([]T, len(srcSlice))
	for i, value := range srcSlice {
		result[i] = value.(T)
	}
	return result
}
