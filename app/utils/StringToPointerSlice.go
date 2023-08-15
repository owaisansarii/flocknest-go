package utils

func StringSliceToPointerSlice(slice []string) []*string {
	pointerSlice := make([]*string, len(slice))
	for i, val := range slice {
		pointerSlice[i] = &val
	}
	return pointerSlice
}
