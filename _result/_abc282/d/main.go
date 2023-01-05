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
		ux := nextInt()
		vx := nextInt()

		if v, ok:= t[ux]; ok {
			t[ux] = append(v, vx)
		} else {
			t[ux] = []int{vx}
		}

		if v, ok:= t[vx]; ok {
			t[vx] = append(v, ux)
		} else {
			t[vx] = []int{ux}
		}
	}

	ng := 0
	r := map[int]int{}
	for i:=1; i<=N; i++ {
		cntm := map[int]int{}
		cnt0, cnt1 := 0,0

		if _, ok := r[i]; ok {continue}
		if ! dfs(i, t, r , cntm, 0) {
			fmt.Println(0)
			return
		}
		
		cnt0 = cntm[0]
		cnt1 = cntm[1]

		if cnt0 >= 2 {
			ng = ng + ((cnt0 - 1) * cnt0) / 2

		}

		if cnt1 >= 2 {
			ng = ng + ((cnt1 - 1) * cnt1) / 2

		}

	}
	cnt := ((N - 1) * N) / 2 - M - ng
	fmt.Println(cnt)
}

func dfs(i int, t map[int][]int, r map[int]int, cntm map[int]int, cnt int) bool{
	if color, ok := r[i]; ok {
		return  color == (cnt)%2 
	} else {
		r[i] = (cnt)%2 


		if c, ok := cntm[cnt%2]; ok {
			c++
			cntm[cnt%2] = c
		} else {
			cntm[cnt%2] = 1
		}

		if a, ok := t[i]; ok {
			cnt++
			for _, v := range a {

				if ! dfs(v, t, r, cntm, cnt) {
					return false
				}
			}
		}
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

func ArrayJoin(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
