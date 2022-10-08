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
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	M := nextInt()

	m := make([][]int, N)
	for i:=range m {
		m[i] = make([]int, N)
	}
	
	for i := 0; i < M; i++ {
		k := nextInt()
		x := make([]int, k)
		for j:=0; j<k; j++ {
			x[j] = nextInt()
		}

		for p:=0; p<k; p++ {
			for q:=p; q<k; q++ {
				m[x[p]-1][x[q]-1] = 1
				m[x[q]-1][x[p]-1] = 1
			}
		}
	}
	cnt := 0
	for _,r := range m {
		for _, w := range r {
			cnt = cnt + w
		}
	}

	if cnt == N*N {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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

func combination(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}
