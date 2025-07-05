package main

import "fmt"

func main() {
	s := ""
	fmt.Scanf("%s", s)
	cnt_0 := 0
	cnt_1 := 0
	ans := 1
	for _, c := range []byte(s) {
		if c == '0' {
			ans += cnt_1
			cnt_0 += 1
		} else {
			ans += cnt_0
			cnt_1 += 1
		}
	}
	fmt.Println(ans)
}


