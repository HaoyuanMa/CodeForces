package main

import "fmt"

func main()  {
	const HELLO = "hello"
	var s string
	fmt.Scanf("%s",&s)
	i := 0
	j := 0
	for  {
		if i > 4 || j >= len(s){
			break
		}
		if s[j] == HELLO[i]{
			i++
			j++
		}else {
			j++
		}
	}
	if i > 4 {
		fmt.Printf("YES\n")
	}else {
		fmt.Printf("NO\n")
	}
}
