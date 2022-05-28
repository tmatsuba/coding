package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
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

	ret := 0
	for _, s := range str {
		ret = ret + atoi(string(s))
	}
	fmt.Println(ret)
}

func atoi(str string) int {
	i, e :=  strconv.Atoi(string(str))
	if e != nil {
		panic(e)
	}
	return i
}
