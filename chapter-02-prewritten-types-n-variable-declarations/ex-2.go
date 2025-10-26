/*Задание: написать программу, с константой value, которая в последующем будет присвоена целочисленной переменной i и переменной с плавающей запятой f. Вывести значения i и f.*/

package main

import "fmt"

func main() {
	const value = 10
	i := value
	var f float64 = value
	fmt.Println(i)
	fmt.Println(f)
}
