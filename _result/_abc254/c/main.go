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
	K := nextInt()

	var a []int
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	ret := "Yes"
	for i := 0; i < K; i++ {
		var b []int
		for j := 0; i + j * K < N; j++ {
			b = append(b, a[i+j*K])
		}
		sort.Sort(sort.IntSlice(b))

		for j := 0; i + j * K < N; j++ {
			a[i+j*K] = b[j]
		}
	}


	r := make([]int, len(a))
	copy(r, a)
	sort.Sort(sort.IntSlice(a))

	for i, v := range r {
		if v != a[i] {
			ret = "No"
			break
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
