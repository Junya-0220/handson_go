package main

import "fmt"

type Data interface {
	Initial(name string, data[]int)
	PrintData()
}

type MyData struct {
	Name string
	Data []int
}

func(md *MyData) Initial(name string, data[]int) {
	md.Name = name
	md.Data = data
}

func(md *MyData) PrintData() {
	fmt.Println("Name: " , md.Name)
	fmt.Println("Data: " , md.Data)
}

func main() {
	var ob MyData = MyData{}
	ob.Initial("Sachiko",[]int{10,29,59})
	ob.PrintData()
}