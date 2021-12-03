package iterable

type IterableSlice struct {
	i int
	s []string
}

func (is *IterableSlice) Next() (value string, done bool) {
	is.i++
	if is.i >= len(is.s) {
		return "", true
	}
	return is.s[is.i], false
}

func NewIterableSlice(s []string) *IterableSlice {
	return &IterableSlice{-1, s}
}
