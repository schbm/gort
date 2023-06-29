package gort

func CreateIntSlice(n int) []int {
	sl := make([]int, n)
	num := n - 1
	for i := 0; i < n; i++ {
		sl[i] = num
		num--
	}
	return sl
}
