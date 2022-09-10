package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
)

func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	d := map[int]int{}
	for i:=0; i<N; i++ {
		v := nextInt()

		var mod int
		if (i-1)%N < 0 {
			mod = (i-1)%N + N
		} else {
			mod = (i-1)%N
		}
		
		d1 := mod - v
		d2 := i - v
		d3 := (i+1)%N - v

		if d1 < 0 {
			d1 += N
		}
		d[d1]++

		if d2 < 0 {
			d2 += N
		}
		d[d2]++

		if d3 < 0 {
			d3 += N
		}
		d[d3]++
	}
	ret := 0
	for _, v := range d {
		ret = Max(ret, v)
	}
	fmt.Println(ret)
}

func check(d []int, v int) bool {
	for _, e := range d {
		if e == v {
			return true
		}
	}
	return false
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

func Abs(x int) int {
	return int(math.Abs(float64(x)))
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

func combination(set []string, n int) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}
