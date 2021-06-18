package main

import (
	"fmt"
)

func main()  {
	var t int
	fmt.Scanf("%d\n",&t)
	for i := 0; i < t; i++ {
		var (
			a int
			b int
			c int
			n int
		)
		fmt.Scanf("%d %d %d %d\n",&a,&b,&c,&n)
		all := a + b + c + n
		if all % 3 != 0{
			fmt.Println("NO")
			continue
		}
		average := all / 3
		if a > average || b > average || c > average{
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
	}
}
