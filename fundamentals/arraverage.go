package main

// *  The {@code Average} class provides a client for reading in a sequence
// *  of real numbers and printing out their average.
// *  <p>
// *  For additional documentation, see <a href="https://algs4.cs.princeton.edu/11model">Section 1.1</a> of
// *  <i>Algorithms, 4th Edition</i> by Robert Sedgewick and Kevin Wayne.
// *
// *  @author Robert Sedgewick
// *  @author Kevin Wayne

// AverageMeh gives aveerage of array
func AverageMeh(arr []int) float64 {
	avg := 0.0
	x := 0.0

	for i := 0; i < len(arr); i++ {
		avg = ((avg * x) / (x + 1)) + (float64(arr[i]) / (x + 1))
		x++
	}

	return avg
}

// Average
func Average(arr []int) float64 {
	sum := 0.0

	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}

	return sum / float64(len(arr))
}
