
package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	// sc.Split(bufio.ScanWords)

	N := nextInt()
	dp := make([]int, N+1)

	for i:=0; i<N+1; i++{
		dp[i] = i
	}

	for j:=6; j<=N; j=j*6 {
		dp[j] = 1
	}

	for j:=9; j<=N; j=j*9 {
		dp[j] = 1
	}
	
	for i:=1; i<=N; i++ {
		
		for j:=6; j+i<=N; j=j*6 {
			dp[i+j] = Min(dp[i+j], dp[i]+1)
		}
		for j:=9; j+i<=N; j=j*9 {
			dp[i+j] = Min(dp[i+j], dp[i]+1)
		}
	}

	fmt.Println(dp[N])
}


var sc = bufio.NewScanner(os.Stdin)

func getMaxI(v, c int) int {
	for i := 1;;i++ {
		if pow(c, i) > c {
			return i - 1
		}
	}
}

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
