
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1001), 1001001)
	sc.Split(bufio.ScanWords)
	N := nextInt()
	A := nextLine()

	var a []int
	var c []int

	for i := 0; i < N; i++ {
		if A[i] == '1' {
			a = append(a, nextInt())
		} else if A[i] == '0' {
			c = append(c, nextInt())
		}
	}
	sort.Ints(a)
	sort.Ints(c)
	a = append(a, math.MaxInt64)
	var ret int

	for i, aw := range(a) {
		r := sort.SearchInts(c, aw)
		ret = Max(ret, len(a) - 1 - i + r)
	}


	fmt.Println(ret)
}


var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

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
