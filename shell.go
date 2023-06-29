package gort

/*
h = 1
while h < n, h = 3*h + 1
while h > 0,
    h = h / 3
    for k = 1:h, insertion sort a[k:h:n]
    → invariant: each h-sub-array is sorted
end

PROPERTIES
Not stable
O(1) extra space
O(n3/2) time as shown (see below)
Adaptive: O(n·lg(n)) time when nearly sorted
*/
func Shell[C Ordered](arr []C) {
	var size int = len(arr)
	//3x+1 incr seq : 1,4,13,40
	var h int = 1    //distance between values
	for h < size/3 { //calc sequence
		h = 3*h + 1
	}
	//h-sort
	for ; h >= 1; h = h / 3 { //iterate until h dist = 1
		//h-sort
		for i := h; i < size; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j = j - h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
	}
}
