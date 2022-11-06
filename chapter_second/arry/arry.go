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