package calculator

import (
	"fmt"
	"strconv"
)

func WichOddOEven(x string) {
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
		t +=5
		fallthrough
		
	case 4:
		t +=4
		fallthrough
	case 3:
		t +=3
		fallthrough
	case 2:
		t +=2
		fallthrough
	case 1:
		t ++
	default:
		fmt.Println("範囲外です")
		return
	}
	fmt.Println(t,"です")
}