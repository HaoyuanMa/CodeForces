package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var s string
		result := int64(0)
		tail := int64(0)
		q := -1
		qcount := 0
		fmt.Scanf("%s\n", &s)
		for j := 0; j < len(s); j++ {
			if s[j] == '?' {
				if q < 0 {
					q = j
					qcount++
					tail = tail + 1
				} else {
					qcount++
					tail = tail + 1
				}
			} else {
				if q < 0 {
					if j == 0 {
						tail = 1
					} else if s[j] == s[j-1] {
						tail = 1
					} else {
						tail = tail + 1
					}
				} else {
					if q == 0 {
						tail = tail + 1
					} else if s[q-1] == s[j] && qcount%2 == 1 {
						tail = tail + 1
					} else if s[q-1] != s[j] && qcount%2 == 0 {
						tail = tail + 1
					} else {
						tail = int64(qcount) + 1
					}
					q = -1
					qcount = 0
				}
			}
			result += tail
		}
		fmt.Printf("%d\n", result)
	}
}
