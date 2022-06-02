package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
)

func main() {
	T := nextInt()
	
	for t := 0 ; t < T; t++ {
		s := nextLine()
		l := len(s)
		ret := pow(10, l - 1) - 1

		for j := 1; j <= l/2; j++ {
			if l % j == 0 {

				init := atoi(string(s[:j]))
				lower := false
				for k := 0; k < l / j; k++ {
					v := atoi(string(s[k * j: (k + 1)* j]))
					if v > init {
						break
					}
					if v < init {
						lower = true
						break
					}
				}

				if lower {
					init--
				}

				tmp := init
				for q := 1; q < l / j; q++ {
					tmp += init * pow(10, q * j)
				}
				ret = Max(tmp, ret)
			}
		}
		fmt.Println(ret)
	}
}


var sc = bufio.NewScanner(os.Stdin)


func atoi(s string) int {
	ret, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return ret
}

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

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
