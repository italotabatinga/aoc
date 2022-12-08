package structures

type Tuple[F, S any] struct {
	First  F
	Second S
}

type Pair[T any] Tuple[T, T]

func (t Tuple[F, S]) Values() (F, S) { return t.First, t.Second }
