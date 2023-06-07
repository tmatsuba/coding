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

	N,M,K := nextInt(),nextInt(),nextInt()
	sumA := make([]int, N+1)
	sumB := make([]int, M+1)

	endA := 0
	for i:=0; i<N; i++ {
		sumA[i+1] = sumA[i] + nextInt()

		if sumA[i+1] <= K {
			endA = i+1
		}
	}

	endB := 0	
	for i:=0; i<M; i++ {
		sumB[i+1] = sumB[i] + nextInt()

		if sumA[endA] + sumB[i+1] <= K {
			endB = i+1
		}
	}

	if endA + endB == N + M {
		fmt.Println(N + M)
		return
	}
	
	ret := endA + endB

//	fmt.Println(sumA, sumB, endA, endB)
	for endA > 0 {
		endA--
		for endB < M && sumA[endA] + sumB[endB+1] <= K  {
			endB++
			ret = Max(ret, endA + endB)
		}
	}

	fmt.Println(ret)
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
