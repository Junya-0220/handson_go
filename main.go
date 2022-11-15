package main

import (
	"fmt"
	"strconv"
	"strings"
)

 type Data interface {
	SetValue(vals map[string]string)
	PrintData()
 }

type MyData struct {
	Name string
	Data []int
}

func (md *MyData) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali,n)
	}
	md.Data = vali
}

func (md *MyData) PrintData() {
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	}else {
		fmt.Println("** This is Nil value. **")
	}
}

type YourData struct {
	Name string
	Mail string
	Age int
}

func (md *YourData) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

func (md *YourData) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", md.Name,md.Age)
	fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {
 var ob *MyData
 ob.PrintData()
 ob = &MyData{}
 ob.SetValue(map[string]string{
	"name": "Jiro",
	"data": "123 456 678",
 })
}