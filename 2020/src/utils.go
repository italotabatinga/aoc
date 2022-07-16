package src

type Part int

const (
	Part1 Part = iota
	Part2
)


func Contains[V comparable](s []V, v V) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}
