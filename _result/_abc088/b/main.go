package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"sort"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	var a []int

	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice := 0
	bob := 0
	for i, v := range(a) {
		if i % 2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}
	fmt.Println(alice - bob)
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
