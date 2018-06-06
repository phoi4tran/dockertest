package main

import "fmt" 


func add(x, y int) int{
	return x + y
}

func sub(x, y int) int{
	return x - y
}

func zero(x int) int{
	return 0
}
func main() {
	fmt.Println(add(1,2))
	fmt.Println(sub(10,5))
	fmt.Println(zero(4))
}
