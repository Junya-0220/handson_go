package paiza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func paiza() {
	sc := bufio.NewScanner(os.Stdin)
    var t []int
    var m []int
    
	for sc.Scan(){
	    str := sc.Text()
        stringSplited := strings.Split(str, " ")
        i1, _ := strconv.Atoi(stringSplited[0])
        i2, _ := strconv.Atoi(stringSplited[1])
        t = append(t,i1)
        m = append(m,i2)
	}
    var t1 = time.Date(2022, 4, 1, t[0], m[0], 0, 0, time.Local)
    var t2 = time.Date(2022, 4, 1, t[1], m[1], 0, 0, time.Local)
    if ((t1.Equal(t2)) || (t1.After(t2))){
        fmt.Println("Yes")
    }else if((t1.Before(t2))){
        fmt.Println("No")
    }

}