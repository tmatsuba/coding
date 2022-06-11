package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
)

func main() {
	sc.Split(bufio.ScanWords)
	
	R := nextInt()
	C := nextInt()
	
	var a1 []int
	var a2 []int

	a1 = append(a1, nextInt())
	a1 = append(a1, nextInt())
	a2 = append(a2, nextInt())
	a2 = append(a2, nextInt())

	if R == 1 {
		fmt.Println(a1[C-1])
	} else if R == 2 {
		fmt.Println(a2[C-1])
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
