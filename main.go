package main

import "fmt"

func main() {

	var num int

	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		result := num * i
		fmt.Println(num, "x", i, "=", result)
	}

}
