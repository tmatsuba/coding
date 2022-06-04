package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"strings"
)

func main() {
	X := nextInt()
	var ret [][]string
	for i := 0; i < X; i++ {
		var a  []string
		for j := 0; j <= i; j++ {
			if i == j || j == 0 {
				a = append(a, "1")
			} else {
				a = append(a, strconv.Itoa((atoi(ret[i-1][j - 1]) + atoi(ret[i-1][j]))))
			}
		}
		ret = append(ret, a)
	}
	
	for _, b := range ret {
		fmt.Println(strings.Join(b, " "))
	}

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

func atoi(s string) int {
	ret, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return ret
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
