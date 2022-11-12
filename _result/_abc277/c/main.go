package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
)

var	m = map[int][]int{}
func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()

	for i:=0; i<N; i++ {
		a := nextInt()
		b := nextInt()

		if _,ok := m[a]; ok {
			m[a] = append(m[a], b)
		} else {
			m[a] = []int{b}
		}

		if _,ok := m[b]; ok {
			m[b] = append(m[b], a)
		} else {
			m[b] = []int{a}
		}
	}
	fmt.Println(recur(1, map[int]bool{}))
}

func recur(now int, gone map[int]bool) int {
	next_list,ok := m[now]
	
	if !ok {
		return now
	}

	gone[now] = true

	ret := now
	for _,v := range next_list {
		if _,ok := gone[v]; ok {
			continue
		}
		ret =  Max(ret, Max(v, recur(v, gone)))
	}
	return ret
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
