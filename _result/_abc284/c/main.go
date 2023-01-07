package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
	"strings"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	M := nextInt()
	t := map[int][]int{}
	for i:=0; i<M; i++ {
		u := nextInt()
		v := nextInt()

		if a, ok := t[u]; ok {
			t[u] = append(a, v)
		} else {
			t[u] = []int{v}
		}

		if a, ok := t[v]; ok {
			t[v] = append(a, u)
		} else {
			t[v] = []int{u}
		}
	}

	r := map[int]bool{}
	cnt := 0
	for i:=1; i<=N; i++ {
		if _,ok := r[i]; ok {
			continue
		} else {
			cnt++
			dfs(t, r, i)
		}
	}
	fmt.Println(cnt)
}

func dfs(t map[int][]int, r map[int]bool, i int) {
	if _,ok := r[i]; ok {
		return
	}

	if a,ok := t[i]; ok {
		r[i] = true
		for _,j := range a {
			dfs(t, r, j)
		}
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
