package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
	"strings"
	"sort"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	M := nextInt()
	a := make([]bool, N+1)
	r := []string{}

	for i:= 0; i<M; i++ {
		a[nextInt()] = true
	}

	m := []int{}
	for i:=1; i<=N; i++ {
		if a[i] {
			m = append(m, i)
		} else {
			if len(m) != 0 {
				m = append(m, i)
				sort.Sort(sort.Reverse(sort.IntSlice(m)))
				for _,v:=range m {
					r = append(r, fmt.Sprint(v))
				}
				m = []int{}
			} else {
				r = append(r, fmt.Sprint(i))
			}
		}
	}

	if len(m) != 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(m)))
		for _,v:=range m {
			r = append(r, fmt.Sprint(v))
		}
	}

	fmt.Println(strings.Join(r, " "))
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

func ArrayJoin(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
