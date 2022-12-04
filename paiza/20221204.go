package paiza

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func paiza() {
	sc := bufio.NewScanner(os.Stdin)
    var t []string
    var m []string
    var b bool

	for sc.Scan(){
	    str := sc.Text()
        stringSplited := strings.Split(str, " ")
        t = append(t,stringSplited[0])
        m = append(m,stringSplited[1])
	}
	
	if(t[0] == t[1] && m[0] < m[1]) {
        b = false
    }
    if(t[0] < t[1] ){
        b = false
    }
    if(t[0] == t[1] && m[0] == m[1]){
        b = true
    }
    if(t[0] == t[1] && m[0] > m[1]) {
        b = true
    }
    if(t[0] > t[1] ){
        b = true
    }

    if(b){
        fmt.Println("Yes")
    }else{
        fmt.Println("No")
    }
}