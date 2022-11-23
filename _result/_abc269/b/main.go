package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	// buf := make([]byte, 1024*1024)
	// sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)


	S := make([]string, 10)
	for i:=0; i<10; i++ {
		S[i] = nextLine() 
	}

	A,B,C,D := 0,10,0,10
	
	for i,Si := range S {
		if strings.Contains(Si, "#") {
			if A == 0 {
				A = i + 1
				for j:=0; j<10; j++ {
					if Si[j] == '#' {
						if C == 0 {
							C = j + 1
						}
					} else {
						if C != 0 && D == 10 {
							D = j
						}
					}
				}
			} else {
				continue
			}
		} else {
			if A != 0 && B == 10 {
				B = i
			}
			continue
		}
	}
	fmt.Println(A, B)
	fmt.Println(C, D)
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
