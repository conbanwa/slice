package slice

// Slice
//
// start Optional
// Zero-based index at which to start extraction, converted to an integer.
//
// Negative index counts back from the end of the array — if start < 0, start + array.length is used.
// If start < -array.length or start is omitted, 0 is used.
// If start >= array.length, nothing is extracted.
//
// end Optional
// Zero-based index at which to end extraction, converted to an integer. slice() extracts up to but not including end.
//
// Negative index counts back from the end of the array — if end < 0, end + array.length is used.
// If end < -array.length, 0 is used.
// If end >= array.length or end is omitted, array.length is used, causing all elements until the end to be extracted.
// If end is positioned before or at start after normalization, nothing is extracted.
func Slice[T comparable](arr []T, limit ...int) []T {
	start := 0
	end := len(arr)
	if len(limit) > 0 {
		start = limit[0]
		if len(limit) > 1 {
			end = limit[1]
		}
	}
	switch true {
	case start <= -len(arr):
		start = 0
	case start < 0:
		start = len(arr) + start
	case start >= len(arr):
		return arr[0:0]
	}
	switch true {
	case end <= -len(arr):
		end = 0
	case end < 0:
		end = len(arr) + end
	case end > len(arr):
		end = len(arr)
	}
	if start > end {
		return arr[0:0]
	}
	return arr[start:end]
}

func Cut[T comparable](arr []T, max ...int) []T {
	limit := 99
	if len(max) != 0 {
		limit = max[0]
	}
	return Slice(arr, 0, limit)
}

func MapKeys[T comparable, V any](m map[T]V) []T {
	sl := make([]T, 0, len(m))
	for k := range m {
		sl = append(sl, k)
	}
	return sl
}

func ToMap[T comparable](sl []T) map[T]struct{} {
	m := make(map[T]struct{}, len(sl))
	for _, k := range sl {
		m[k] = struct{}{}
	}
	return m
}

func UniqSlice[T comparable](list []T) (x []T) {
	for _, i := range list {
		hasSeen := false
		for _, v := range x {
			if i == v {
				hasSeen = true
				break
			}
		}
		if !hasSeen {
			x = append(x, i)
		}
	}
	return
}
