package main

import "fmt"

func main() {
	var a [10]int

	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 3: 4, 5}
	d := [10]int{1, 3: 4, 5, 8: 9}

	a[0] = 1
	a[1] = 2

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var dp [5][5]int

	dp[0][0] = 1
	dp[0][1] = 2
	dp[4][4] = 3

	s := b[2:5]
	f := d[3:]

	defer fmt.Println(s)
	defer fmt.Println(f)

	fmt.Println(dp)
}
