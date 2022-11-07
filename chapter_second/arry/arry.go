package arry

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrySum(x string) {
	ar := strings.Split(x, " ")
	t := 0

	for i  := 0; i < len(ar); i++ {
		n,err := strconv.Atoi(ar[i])
		if err != nil {
			fmt.Println("Error")
		}
		t += n
	}
	fmt.Println("total:", t)
}

func ArryRange(x string) {
	ar := strings.Split(x, " ")
	t := 0

	for _ , v :=  range ar {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error")
		}
		t += n
	}
}
// 最後に追加 変数=push(スライス,値)
func Push(a []int, v int) []int {
	return append(a,v)
}

// 最後を削除 変数=push(スライス)
func Pop(a []int) []int {
	return a[:len(a)-1]
}

//最初に追加 変数=unshift(スライス,値)
func Unshift(a []int, v int) []int{
	return append([]int{v},a...)
}

//最初を削除 変数=shift(スライス)
func Shift(a []int, v int) []int {
	return a[1:]
}

//指定の位置に追加 変数=insert(スライス,値,インデックス)
func Insert(a []int, v int, p int) []int {
	a = append(a,0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

//指定の位置を削除 変数=remove(スライス,インデックス)
func Remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
