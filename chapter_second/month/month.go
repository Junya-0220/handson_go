package month

import (
	"fmt"
	"strconv"
)

func WhatMonth(x string){
	switch n, err := strconv.Atoi(x); n{
	case 0:
		fmt.Println("整数値が得られませんでした")
		fmt.Println(err)
	case 1,2,12:
		fmt.Println("冬です")
	case 3,4,5:
		fmt.Println("春です")
	case 6,7,8:
		fmt.Println("夏です")
	case 9,10,11:
		fmt.Println("秋です")
	default:
		fmt.Println("月の値じゃありませんよ？")
	}

}

