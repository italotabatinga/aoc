package collections

type Tuple[T, U any] struct {
	First  T
	Second U
}

type Pair[T any] Tuple[T, T]
