package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
)

var X int
func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	X = nextInt()

	L := make([][]int, N)
	for i:=0; i<N; i++ {
		l := nextInt()
		la := make([]int, l)
		for j:=0; j<l; j++ {
			la[j] = nextInt()
		}
		L[i] = la
	}

	

	fmt.Println(dfs(L, 1, 0))
}

func dfs(L [][]int, sum, ret int) int {

	la := L[0]
	for _,v := range la {
		if len(L) == 1 && sum * v == X {
			ret++
		} else if sum * v > X {
			continue
		} else {
			if len(L) > 1 {
				ret = dfs(L[1:], sum * v, ret)
			}
		}
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
