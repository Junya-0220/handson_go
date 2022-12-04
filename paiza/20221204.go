package paiza

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func paiza() {
	sc := bufio.NewScanner(os.Stdin)
    var t []string
    var m []string
    
	for sc.Scan(){
	    str := sc.Text()
        stringSplited := strings.Split(str, " ")
        t = append(t,stringSplited[0])
        m = append(m,stringSplited[1])
	}
	tf1 := fmt.Sprintf("2020-11-01T%s:%s:00+09:00",t[0],m[0])
	tf2 := fmt.Sprintf("2020-11-01T%s:%s:00+09:00",t[1],m[1])
	t1, _ := time.Parse(time.RFC3339, tf1)
	t2, _ := time.Parse(time.RFC3339, tf2)
    if((t1.Equal(t2)) || (t1.After(t2))){
        fmt.Println("Yes")
    }else if(t1.Before(t2)){
         fmt.Println("No")
    }

}