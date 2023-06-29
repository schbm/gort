package gort

import (
	"math/rand"
)

/*
PROPERTIES
Not stable
O(lg(n)) extra space (see discussion)
O(n2) time, but typically O(n·lg(n)) time
Not adaptive

_# choose pivot_
swap a[1,rand(1,n)]

_# 2-way partition_
k = 1
for i = 2:n, if a[i] < a[1], swap a[++k,i]
swap a[1,k]
_→ invariant: a[1..k-1] < a[k] <= a[k+1..n]_

_# recursive sorts_
sort a[1..k-1]
sort a[k+1,n]
*/

func Quick[C Ordered](arr []C) {
	var size int = len(arr)
	if size <= 1 {
		return // Base case: already sorted
	}

	//_# choose pivot_
	pivotIndex := rand.Intn(size)
	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0] // Move pivot to the first position

	//_# 2-way partition_
	pivot := arr[0]
	i := 1
	j := size - 1

	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i] // Swap elements that are out of place
		}
	}

	arr[0], arr[j] = arr[j], arr[0] // Move pivot to its final sorted position

	Quick(arr[:j])   // Recursively sort the left partition
	Quick(arr[j+1:]) // Recursively sort the right partition
}

func QuickOpt[C Ordered](arr []C) {
	const threshold = 10
	quickSort(arr, 0, len(arr)-1, threshold)
}

func quickSort[C Ordered](arr []C, lo, hi, threshold int) {
	if hi-lo <= threshold {
		// Use Insertion Sort for small subarrays
		Insertion(arr[lo : hi+1])
		return
	}

	// Median-of-Three pivot selection
	mid := (lo + hi) / 2
	if arr[lo] > arr[mid] {
		arr[lo], arr[mid] = arr[mid], arr[lo]
	}
	if arr[mid] > arr[hi] {
		arr[mid], arr[hi] = arr[hi], arr[mid]
		if arr[lo] > arr[mid] {
			arr[lo], arr[mid] = arr[mid], arr[lo]
		}
	}

	pivot := arr[mid]
	i, j := lo, hi

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		if arr[i] == arr[j] {
			i++
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	quickSort(arr, lo, j, threshold)   // Sort left partition
	quickSort(arr, j+1, hi, threshold) // Sort right partition
}
