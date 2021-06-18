package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n",&t)
	for i := 0; i < t; i++ {
		var s string
		result := int64(0)
		fmt.Scanf("%s\n",&s)
		var tail [200002]int64
		for j := 0; j < len(s); j++ {
			inc := int64(0)
			q := -1
			qcount := 0

			inc =

			for h := j; h >= 0; h-- {
				if h == j {
					if s[h] == '?' {
						inc++
						q = h
						qcount++
					} else {
						inc++
					}
					continue
				}

				if q >= 0 {
					if s[h] == '?' {
						qcount++
						inc++
					} else {
						if q == j {
							qcount++
							inc++
						} else {
							if s[q+1] == s[h] && qcount % 2 == 1 {
								inc++
							} else if s[q+1] != s[h] && qcount % 2 ==0 {
								inc++
							} else {
								break
							}
						}
						qcount = 0
						q = -1
					}
				} else {
					if s[h] == '?' {
						inc++
						q = h
						qcount++
					} else if s[h] != s[h+1] {
						inc++
					} else {
						break
					}
				}
			}
			result += inc
		}
		fmt.Printf("%d\n",result)
	}
}
