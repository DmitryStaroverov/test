package main

import "fmt"

func main() {

	var num int

	var sum int = 0

	fmt.Print("Введите число: ")

	fmt.Scan(&num)

	for i := 1; i <= num; i++ {

		sum += i

	}

	fmt.Print("Сумма: ", sum)

}
