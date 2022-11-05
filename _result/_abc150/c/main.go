package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
)

var N int
var P []int
var Q []int
var p_cnt int
var q_cnt int
var cnt int = 0

func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N = nextInt()
	P = make([]int, N)
	Q = make([]int, N)


	for i:=0; i<N; i++ {
		P[i] = nextInt()
	}
	for i:=0; i<N; i++ {
		Q[i] = nextInt()
	}

	dfs(0, []int{}, map[int]bool{}) 

	fmt.Println(Abs(p_cnt - q_cnt))
}

func dfs(i int, a []int, m map[int]bool) {
	if len(a) == N {

		cnt++
		if SliceEqual(a, P) {
			p_cnt = cnt
		}

		if SliceEqual(a, Q) {
			q_cnt = cnt
		}
		return
	}

	if i == 0 {
		m = map[int]bool{}
	}

	for j:=1; j<=N; j++ {
		if v, ok := m[j]; ok && v {
			continue
		}
		a = append(a, j)
		m[j] = true

		dfs(j, a, m)
		if p_cnt == 0 || q_cnt == 0 {
			a = a[:len(a)-1]
			m[j] = false
		} else {

			return
		}
	}
}

func SliceEqual(a,b []int) bool {
	if len(a) != len(b) {return false}

	for i, v := range a {
		if v != b[i] {return false}
	}
	return true
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
