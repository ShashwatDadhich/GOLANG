package main

import (
	"fmt"
)

func add2(x, y int) int{
	return x+y;
}

func main(){
	c := add2(2,3)
	fmt.Println(c);
}