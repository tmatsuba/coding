package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
	str := nextLine()
	ret := ""
	for i,  s := range str {
		if i == 0 {
			ret = ret + strings.ToUpper(string(s))
		} else {
			ret = ret + strings.ToLower(string(s))
		}
	}
	fmt.Println(ret)
}
