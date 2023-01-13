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

	H := nextInt()
	W := nextInt()
	N := nextInt()

	Z := make([][]int, H+1)

	for i:=0; i<=H; i++ {
		S := make([]int, W+2)
		Z[i] = S
	}

	for i:=0; i<N; i++ {

		a := nextInt()
		b := nextInt()
		c := nextInt()
		d := nextInt()

		for i:=a; i<=c; i++ {
			Z[i][b] = Z[i][b] + 1
			Z[i][d+1] = Z[i][d+1] - 1
		}

	}
	for i:=1; i<=H; i++ {
		for j:=1; j<=W; j++ {
			Z[i][j] = Z[i][j-1] + Z[i][j]
		}
	}

	for i:=1; i<=H; i++ {
		for j:=1; j<=W; j++ {
			fmt.Print(Z[i][j])
			if j != W {
				fmt.Print(" ")				
			}
		}
		fmt.Println("")
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
