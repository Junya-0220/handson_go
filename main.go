package main

import (
	"fmt"
	calculator "handson_go/chapter_second/calc"
	"handson_go/hello"
)

func main(){
	x := hello.Input("type 1~5")
	calculator.FallThrough(x)
}

func f(n int) int{
	fmt.Println("No,",n)
	return n
}