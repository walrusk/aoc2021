package iterable

type StringIterator interface {
	Next() (value string, done bool)
}
