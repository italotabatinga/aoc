package aoc

import (
	"hash/fnv"
	"math"

	"golang.org/x/exp/constraints"
)

func Sort[V any](a []V, f func(V, V) int) {
	quickSort(a, 0, len(a)-1, f)
}

func quickSort[V any](a []V, low int, high int, f func(V, V) int) {
	if low < high {
		pi := partition(a, low, high, f)
		quickSort(a, low, pi-1, f)
		quickSort(a, pi+1, high, f)
	}
}

func partition[V any](a []V, low int, high int, f func(V, V) int) int {
	// pivot (Element to be placed at right position)
	pivot := a[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if f(a[j], pivot) > 0 {
			i++ // increment index of smaller element
			tmp := a[i]
			a[i] = a[j]
			a[j] = tmp
		}
	}

	tmp := a[i+1]
	a[i+1] = a[high]
	a[high] = tmp
	return i + 1
}

func Abs[T constraints.Integer](x T) T {
	if x > 0 {
		return x
	}

	return -x
}

func Min(elems ...int) int {
	min := math.MaxInt
	for _, elem := range elems {
		if elem < min {
			min = elem
		}
	}

	return min
}

func StringToInt(str string) int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

func Compare[C constraints.Ordered](a, b C) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}

	return 0
}
