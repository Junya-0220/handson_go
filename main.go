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
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
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
	ob := [2]Data{}
	ob[0] = new(MyData)
	ob[0].SetValue(map[string]string{
		"name": "sachiko",
		"data": "22, 33, 44",
	})
	ob[1] = new(YourData)
	ob[1].SetValue(map[string]string {
		"name": "mami",
		"mail": "mami@gmail.com",
		"age": "45",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}