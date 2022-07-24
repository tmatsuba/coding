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
	sc.Split(bufio.ScanWords)
	
	N := nextInt()
	M := nextInt()
	X := map[int]int{}
	C := map[int]int{}
	
	dp := [][]int{}
	for i := 0; i<=N; i++ {
		v := []int{}
		for j := 0; j<=N; j++ {
			v = append(v, math.MinInt64)
		}
		dp = append(dp, v)
	}
	
	for i:=0; i<N; i++ {
		X[i+1] = nextInt()
	}

	for i:=0; i<M; i++ {
		C[nextInt()] = nextInt()
	}


	dp[0][0] = 0
	for i:=1; i<=N; i++ {
		for j:=0; j<=i; j++ {
			if j == 0{
				tmp := 0
				for _,value := range dp[i-1] {
					tmp = Max(tmp, value)
				}
				dp[i][0] = tmp
			} else {
				dp[i][j] = dp[i-1][j-1] + X[i] + C[j]
			}
		}
	}

	ans := 0
	for _,v := range dp[N] {
		ans = Max(ans, v)
	}
	fmt.Println(ans)
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
