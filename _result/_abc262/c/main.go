package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"

)

func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	m := map[int]int{}
	c := 0
	ret := 0

	for i:=0; i<N; i++ {
		v := nextInt()
		if i+1 == v {
			c++ 
		} else  {
		   m[i+1] = v
		}
	}

	for k,v := range m{
		elem, ok  := m[v]
		if ok && elem == k{
			ret++
			m[v] = -1
		}
	}
	ret = ret + combination(c, 2)
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

func permutation(n int, k int) int {
    v := 1
    if 0 < k && k <= n {
        for i := 0; i < k; i++ {
            v *= (n - i)
        }
    } else if k > n {
        v = 0
    }
    return v
}

func factorial(n int) int {
    return permutation(n, n - 1)
}

func combination(n int, k int) int {
    return permutation(n, k) / factorial(k)
}


