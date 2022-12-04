package paiza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 空白や空文字だけの値を除去したSplit()
func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
        stringSplited := strings.Split(stringTargeted, delim)
        for _, str := range stringSplited {
                if str != "" {
                        stringReturned = append(stringReturned, str)
                }
        }
        return
}

// 空白や空文字だけの値を除去したSplitN()
func SplitWithoutEmptyN(stringTargeted string, delim string, n int) (stringReturned []string) {
        stringSplited := strings.SplitN(stringTargeted, delim, n)
        for _, str := range stringSplited {
                if str != "" {
                        stringReturned = append(stringReturned, str)
                }
        }
        return
}

// デリミタで分割して整数値スライスを取得
func SplitAndConvertToInt(stringTargeted, delim string) (intSlices []int, err error) {
        // 分割
        stringSplited := SplitWithoutEmpty(stringTargeted, delim)
        // 整数スライスに保存
        for i := range stringSplited {
                var iparam int
                iparam, err = strconv.Atoi(stringSplited[i])
                if err != nil {
                        return
                }
                intSlices = append(intSlices, iparam)
        }
        return
}

/**
 * paiza専用: 行指定なしデータ列読み込み
 *
 */
func PaizaSequenceGetsAsInt() (line [][]int) {
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
                str := scanner.Text()
                data, err := SplitAndConvertToInt(str, " ")
                if err != nil || data == nil {
                        break
                }
                line = append(line, data)
        }
        return
}

/**
 * paiza専用: 行指定なし文字列読み込み
 *
 */
func PaizaSequenceGets() (lines [][]string) {
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
                str := scanner.Text()
                data := SplitWithoutEmpty(str, " ")
                if data == nil {
                        break
                }
                lines = append(lines, data)
        }
        return
}

/**
 * paiza専用: 「文字 数値 数値 数値...」の取得
 *
 */
func PaizaStrFirstIntAfter(data string, delim string) (head string, body []int) {
        list := SplitWithoutEmptyN(data, delim, 2)
        head = list[0]
        body, _ = SplitAndConvertToInt(list[1], delim)
        return
}

/**
 * paiza専用: 文字列一括取得・最初X,Y、以下Y行文字列
 *
 */
func PaizaGetsXY() (x, y int, line []string) {
scanner := bufio.NewScanner(os.Stdin)
	first := true
	for scanner.Scan() {
			str := scanner.Text()
			if first {
					xy, _ := SplitAndConvertToInt(str, " ")
					x = xy[0]
					y = xy[1]
					first = false
			} else {
					line = append(line, strings.TrimSpace(str))
			}
	}
	return
}

/**
 * paiza専用: 文字列一括取得・最初行数、以下文字列
 *
 */
func PaizaGets() (count int, line []string) {
        scanner := bufio.NewScanner(os.Stdin)

        n := -1
        for scanner.Scan() {
                str := scanner.Text()
                if n == -1 {
                        count, _ = strconv.Atoi(strings.TrimSpace(str))
                } else {
                        line = append(line, strings.TrimSpace(str))
                }
                n += 1

                if n >= count {
                        break
                }
        }
        return
}

/**
 * paiza専用: 数字列1行取得・1整数以上対応
 * TestPaizaGetNum
 */
func PaizaGetNums() (intReturned []int) {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        str := scanner.Text()
        intReturned, _ = SplitAndConvertToInt(str, " ")
        return
}

/**
 * paiza専用: 文字列取得・複数単語
 * 
 */
func PaizaGetWordList() (stringReturned []string) {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        str := scanner.Text()
        stringReturned = SplitWithoutEmpty(str, " ")
        return
}

/**
 * paiza専用: 文字列取得・1単語
 * 
 */
func PaizaGetWord() (stringReturned string) {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        stringReturned = scanner.Text()
        return
}

///////////////////////////////////
// 使い方サンプル
///////////////////////////////////

// 1行・文字・1単語
/**
	* 1行・文字・1単語
	*
	* 入力例)
	* foo
*/
func TestPaizaGetWord() {
        line := PaizaGetWord()
        fmt.Println(line)
}

// 1行・文字・複数単語
/**
	* 1行・文字・複数単語
	*
	* 入力例)
	* foo bar baz
*/
func TestPaizaGetWordList() {
        list := PaizaGetWordList()
        for _, str := range list {
                fmt.Println(str)
        }
}

// 1行・数値・空白区切り対応
/**
	* 1行・数値・複数整数(空白区切りのない整数値も可)
	*
	* 入力例)
	* 12 34 56
*/
func TestPaizaGetNum() {
        nums := PaizaGetNums()
        fmt.Println(nums)
        // 数値であることを示すため
        if len(nums) >= 2 {
                fmt.Printf("%d+%d=%d\n", nums[0], nums[1], nums[0]+nums[1])
        } else {
                fmt.Printf("%d*2=%d\n", nums[0], nums[0]*2)
        }
}

// 1行目.......継続行数
// 2行目以降...文字列(空白区切り)
/**
	* 1行目.......継続行数
	* 2行目以降...文字列(空白区切り)
	*
	* 入力例)
	* 3
	* AA BB CC
	* DD EE FF
	* GG HH II
*/
func TestPaizaGets() {
        count, line := PaizaGets()
        fmt.Printf("count=%d\n", count)
        for _, d := range line {
                strs := SplitWithoutEmpty(d, " ")
                fmt.Println(strs)
        }
}

// 1行目.......継続行数
// 2行目以降...文字、整数値列(空白区切り)
/**
	* 1行目.......継続行数
	* 2行目以降...文字、整数値列(空白区切り)
	*
	* 入力例)
	* 3
	* A 10 20 30
	* B 11 22 33
	* C 99 88 77
*/
func TestPaizaStrFirstIntAfter() {
	count, line := PaizaGets()
	fmt.Printf("count=%d\n", count)
	for _, d := range line {
		head, body := PaizaStrFirstIntAfter(d, " ")
		fmt.Println(head)
		fmt.Println(body)
		fmt.Printf("%d+%d=%d\n", body[0], body[1], body[0]+body[1]) // 数値であることを示すため
	}
}



// 改行のみ続くまで数値列入力
/**
	* 改行のみ続くまで数値列入力
	*
	* 入力例)
	* 2 3 4
	* 11 12 34
	* 9 87 654
	*    :
	*    :
	* (空行)
*/
func TestPaizaSequenceGetsAsInt() {
        line := PaizaSequenceGetsAsInt()
        for _, d := range line {
                fmt.Println(d)
                fmt.Printf("%d+%d+%d=%d\n", d[0], d[1], d[2], d[0]+d[1]+d[2])
        }
}

// 改行のみ続くまで文字列入力
/**
	* 改行のみ続くまで文字列入力
	*
	* 入力例)
	* foo bar baz
	* hoge fuga
	* piyo
	* AA BB
	*    :
	*    :
	* (空行)
*/
func TestPaizaSequenceGets() {
        line := PaizaSequenceGets()
        for i, d := range line {
                fmt.Printf("-- %d --\n", i)
                for _, s := range d {
                        fmt.Printf("%s\n", s)
                }
        }
}

// main()
func main() {
	TestPaizaSequenceGets()
}