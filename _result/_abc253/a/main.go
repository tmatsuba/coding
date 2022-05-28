package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()
	ret := "No"

	if a <= b {
		if b <= c {
			ret = "Yes"
		}
	} 
	if a >= b {
		if b >= c {
			ret = "Yes"
		}
	}
	fmt.Println(ret)
}

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

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
