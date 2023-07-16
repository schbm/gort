package gort

/*
		PROPERTIES
		Not stable
		O(1) extra space
		Θ(n2) comparisons
		Θ(n) swaps
		Not adaptive

		for i = 1:n,
	    k = i
	    for j = i+1:n, if a[j] < a[k], k = j
	    → invariant: a[k] smallest of a[i..n]
	    swap a[i,k]
	    → invariant: a[1..i] in final position
		end

		start with min variation
*/
func Selection[C Ordered](arr []C) {
	var (
		size         int = len(arr)
		currentIndex int // current index of iteration -> left to right
		minIndex     int // index of num smaller than current index
	)

	for currentIndex = 0; currentIndex < size-1; currentIndex++ {
		minIndex = currentIndex
		for i := currentIndex + 1; i < size; i++ { //set min index to value which is smaller
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
		}
		arr[currentIndex], arr[minIndex] = arr[minIndex], arr[currentIndex]
	}
}
