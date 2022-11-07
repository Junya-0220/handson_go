package main

import (
	"fmt"
	"handson_go/chapter_second/arry"
)

func main(){
	a:= []int{10,20,30}
	fmt.Println(a)
	a = arry.Push(a,1000)
	fmt.Println(a)
	a = arry.Pop(a)
	fmt.Println(a)
	a = arry.Unshift(a,90)
	fmt.Println(a)
	a = arry.Shift(a,500)
	fmt.Println(a)
	a = arry.Insert(a,400,2)
	fmt.Println(a)
	a = arry.Remove(a,2)
	fmt.Println(a)
}


