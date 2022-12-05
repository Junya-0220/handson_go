package paiza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func f20221205() {
        sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := (sc.Text())
        stringSplited := strings.Split(s, " ")
        fmt.Println(stringSplited)
        con , _:= strconv.Atoi(stringSplited[1])
        var r int


	for sc.Scan(){
                str := sc.Text()
                i, _ := strconv.Atoi(str)
                if (i <= con) {
                        r += i
                }
	}
        fmt.Println(r)
}