package main

import "fmt"

func main()  {
	var w int
	fmt.Scanf("%d",&w)
	if w < 4 || w % 2 == 1{
		fmt.Printf("NO")
	}else {
		fmt.Printf("YES")
	}
}
