package main

import "fmt"

func main()  {
	var (
		a int
		b int
	)
	fmt.Scanf("%d %d",&a,&b)
	m := 0
	result := 0
	for{
		if a <= 0 {
			break
		}
		result += a
		new := (a + m) / b
		m = (a + m) % b
		a = new

	}
	fmt.Printf("%d\n",result)
}

