package main

import "fmt"

func main() {
	var i = 20
	var f = float64(i) //происходит преобразование типа int в float64
	fmt.Println(i)    //20
	fmt.Println(f)    //20
}
