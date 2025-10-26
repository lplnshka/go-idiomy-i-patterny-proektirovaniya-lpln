/*Задание: написать программу с тремя переменными: b типа byte, smallI типа int32 и bigI типа uint64. Присвоить им максимальные допустимые значения соответственно типам, затем добавить к каждой 1 и вывести значения.*/

package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println(b) //выводит 0
	fmt.Println(smallI) //выводит -2147483648
	fmt.Println(bigI) //выводит 0
}
