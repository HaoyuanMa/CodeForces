package main

import (
	"fmt"
	"sort"
)


type IntSlice []int64

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}


func main() {
	var n int
	var l int64
	var a IntSlice
	fmt.Scanf("%d %d\n",&n,&l)
	a = make([]int64,n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d",&a[i])
	}
	sort.Sort(a)
	d := float64(a[0])
	last := 0
	next := 1
	for next < n {
		if float64(a[next] - a[last]) > 2 * d {
			d = float64(a[next] - a[last]) / 2
		}
		next++
		last++
	}
	if float64(l - a[n-1]) > d {
		d = float64(l - a[n-1])
	}
	fmt.Printf("%f",d)
}

