package main

import (
	"fmt"
	"sort"

	"github.com/antimatter96/algo4e/fundamentals/search"
)

var arr = []int{
	23, 43, 144, 171, 175, 191, 249, 343, 381, 425, 429, 449, 512, 519, 544, 547, 654, 690, 706, 707, 714, 715, 810, 848, 910, 911, 994, 1003, 1090, 1098, 1109, 1147, 1212, 1249, 1251, 1262, 1298, 1312, 1370, 1399, 1449, 1472, 1720, 1731, 1792, 1850, 1896, 1956, 1977, 1992,
}

func main() {
	sort.Ints(arr)
	fmt.Println(search.BinarySearch(arr, 191))

	a := make([]int, 100)

	lo := 20
	high := 56
	n := 99

	RandomSeq(lo, high, n, a)
	fmt.Println(a)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[a[i]]++
	}
	fmt.Println(len(mp))

	for i := lo; i < high; i++ {
		fmt.Printf("%d -> %d\n", i, mp[i])
	}
}
