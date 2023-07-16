package gort

/*
for i = 1:n,

	swapped = false
	for j = n:i+1,
	    if a[j] < a[j-1],
	        swap a[j,j-1]
	        swapped = true
	→ invariant: a[1..i] in final position
	break if not swapped

end

PROPERTIES
Stable
O(1) extra space
O(n2) comparisons and swaps
Adaptive: O(n) when nearly sorted
*/
func Bubble[C Ordered](arr []C) {
	var (
		size      int = len(arr)
		currIndex int
		isSwapped bool = false
	)

	for currIndex = 0; currIndex < size-1; currIndex++ {
		isSwapped = false
		for i := 0; i < size-currIndex-1; i++ {
			if arr[i] > arr[i+1] { //this moves the biggest el to the end of the list
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
