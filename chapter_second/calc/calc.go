package calculator

import (
	"fmt"
	"strconv"
)

func WichOddOrEven(x string) {
	n, err := strconv.Atoi(x)
	if err != nil{
		fmt.Println("Error")
		return
	}
	switch {
	case n % 2 == 0:
		fmt.Println("偶数です")
	case n % 2 == 1:
		fmt.Println("奇数です")
	default:
		fmt.Println("不正な値です")
	}
}
// 入力された1から5までの合計値をコンソールに出す関数
func FallThrough(x string) {
	n,err := strconv.Atoi(x)
	if err != nil{
		fmt.Println("Error")
		return
	}
	fmt.Print(x + "までの合計は")
	t := 0
	switch n{
	case 5:
		t += 5
		fallthrough
		
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t ++
	default:
		fmt.Println("範囲外です")
		return
	}
	fmt.Println(t,"です")
}

// 入力された値の合計値をコンソールに出す関数
func ForStatement(x string) {
	n,err := strconv.Atoi(x)
	if err != nil{
		fmt.Println("Error")
		return
	}
	fmt.Print("1から" + x + "までの合計は、")
	t := 0
	c := 1
	for c <= n {
		t += c
		c ++
		fmt.Println(t)
	}
	fmt.Println(t,"です。")
}
// 入力された値の合計値をコンソールに出す関数
func ForStatement2(x string) {
	n,err := strconv.Atoi(x)
	if err != nil{
		fmt.Println("Error")
		return
	}
	fmt.Print("1から" + x + "までの合計は、")
	t := 0
	for i := 1; i <= n; i++ {
		t += i;
		fmt.Println(t)
	}
	fmt.Println(t,"です。")
}

//forループの中でbreak,continueを使ってみる
func ForStatement3(x string){
	n,err := strconv.Atoi(x)
	if err != nil{
		fmt.Println("Error")
		return
	}
	fmt.Print("1から" + x + "までの偶数の合計は、")
	t := 0
	c := 0
	for i := 1; i <= n; i++ {
		c ++
		if c % 2 == 1{
			continue
		}
		if c > n {
			break
		}
		t += c
	}
	fmt.Println(t,"です。")
}