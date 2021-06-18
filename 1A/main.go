package main

import "fmt"

func main()  {
	var (
		m uint64
		n uint64
		a uint64
	)
	fmt.Scanf("%d %d %d",&m,&n,&a)
	x := m / a
	if m % a != 0{
		x = x + 1
	}
	y := n / a
	if n % a != 0 {
		y = y + 1
	}
	result := x*y
	fmt.Printf("%d\n",result)
}
