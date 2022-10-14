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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	W := nextInt()
	N := nextInt()
	K := nextInt()
	A := make([][]int, N+1)
//	dp := [20][20][20]int{}
	dp := [51][51][10001]int{}
	ret := 0

	for i:=0; i<N; i++ {
		e := make([]int, 2)
		e[0] = nextInt()
		e[1] = nextInt()
		A[i] = e
	}

	for i:=0; i<N; i++ {
		item_w := A[i][0]
		item_v := A[i][1]

		for k:=0; k<=K; k++ {
			for w :=0; w <= W; w++ {
				
				dp[i+1][k][w] = Max(dp[i][k][w], dp[i+1][k][w])

				if w - item_w >= 0 && k < K {
					dp[i+1][k+1][w] = Max(dp[i+1][k+1][w], dp[i][k][w-item_w]+item_v)
				}

			}	

		}
//		fmt.Println(dp[i+1])
	}

	for i:=0; i<=N; i++ {
		for k:=0; k<=K; k++ {
			for w:=0; w<=W; w++ {
				ret = Max(ret, dp[i][k][w])
			}
		}
	}
	fmt.Println(ret)	
//	fmt.Println(dp[N +1][W][0])
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
