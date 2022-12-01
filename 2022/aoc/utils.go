package aoc

import "golang.org/x/exp/constraints"

func Sort[V any, C constraints.Ordered](a []V, f func(V) C) {
	quickSort(a, 0, len(a)-1, f)
}

func quickSort[V any, C constraints.Ordered](a []V, low int, high int, f func(V) C) {
	if low < high {
		pi := partition(a, low, high, f)
		quickSort(a, low, pi-1, f)
		quickSort(a, pi+1, high, f)
	}
}

func partition[V any, C constraints.Ordered](a []V, low int, high int, f func(V) C) int {
	// pivot (Element to be placed at right position)
	pivot := a[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if f(a[j]) < f(pivot) {
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
