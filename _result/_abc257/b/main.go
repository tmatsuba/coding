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
	sc.Split(bufio.ScanWords)
	N := nextInt()
	K := nextInt()
	Q := nextInt()
	var a []int

	for j := 0; j < K; j++ {
		a = append(a, nextInt())
	}

	for k := 0; k < Q; k ++ {
		l := nextInt()
		if l == K && a[l-1] < N {
			a[l-1]++
			continue
		}
		if l < K && a[l-1] != a[l]-1 {
			a[l-1]++
			continue
		}
	}
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), " "), "[]"))
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
