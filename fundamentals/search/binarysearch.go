package search

// The {@code BinarySearch} class provides a static method for binary
//  *  searching for an integer in a sorted array of integers.
//  *  <p>
//  *  The <em>indexOf</em> operations takes logarithmic time in the worst case.
//  *  <p>
//  *  For additional documentation, see <a href="https://algs4.cs.princeton.edu/11model">Section 1.1</a> of
//  *  <i>Algorithms, 4th Edition</i> by Robert Sedgewick and Kevin Wayne.

// BinarySearch must be given a sorted array
func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for mid := (high-low)/2 + (low); low <= high; mid = (high-low)/2 + (low) {
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
