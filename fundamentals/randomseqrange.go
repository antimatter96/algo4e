package main

import (
	"math/rand"
	"time"
)

//  *  The {@code RandomSeq} class is a client that prints out a pseudorandom
//  *  sequence of real numbers in a given range.
//  *  <p>
//  *  For additional documentation, see <a href="https://algs4.cs.princeton.edu/11model">Section 1.1</a> of
//  *  <i>Algorithms, 4th Edition</i> by Robert Sedgewick and Kevin Wayne.
//  *
//  *  @author Robert Sedgewick
//  *  @author Kevin Wayne

// RandomSeq fills `dest` with random numbers betweel [low, high)
func RandomSeq(low, high, n int, dest []int) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	if n > len(dest) {
		panic("Wanted more than given")
	}

	limit := high - low

	for i := 0; i < n; i++ {
		dest[i] = low + random.Intn(limit)
	}
}
