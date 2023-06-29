package gort

/*
PROPERTIES
Stable
O(1) extra space
O(n2) comparisons and swaps
Adaptive: O(n) time when nearly sorted
Very low overhead

for i = 2:n,
    for (k = i; k > 1 and a[k] < a[k-1]; k--)
        swap a[k,k-1]
    → invariant: a[1..i] is sorted
end
*/

func Insertion[C Ordered](arr []C) {
	var currVal C
	var currentIndex int
	var currentIndexLeft int
	var size int = len(arr)

	for currentIndex = 1; currentIndex < size; currentIndex++ { //start on second item and iterate to the right
		currVal = arr[currentIndex]         //cache curr item
		currentIndexLeft = currentIndex - 1 //compare with the left (repr. sorted list)

		for ; currentIndexLeft >= 0 && currVal < arr[currentIndexLeft]; currentIndexLeft-- { //iterate down and compare
			arr[currentIndexLeft+1] = arr[currentIndexLeft] //shift to the right
		}
		arr[currentIndexLeft+1] = currVal
	}
}
