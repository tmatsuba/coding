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

	N, K := nextInt(),nextInt()
	s := nextLine()

	starts := []int{0}

	one, zero := false, false
	for i:=0; i<N; i++ {

		if s[i] == '0' {
			zero,one = true,false
		} else {
			if zero && !one && i != 0 {
				starts = append(starts, i)
			} 
			zero,one = false,true
		}
	}

	k := K
	cnt := 0
	end := false
	r:=0

	for i:=0; i<len(starts); i++ {
		zero = false
		for j:=r; j<N; j++ {
			if j == N-1 {
				end = true
			}

			if s[j] == '0' && !zero {
				k--
				zero = true
			} else if s[j] == '1' {
				zero = false
			}

			if k < 0 {
				if 1 <= K {
					k = 1
				} else  {
					k = K
				}
				break
			} else {
				cnt = Max(cnt, j - starts[i] + 1)
				r = j
			}

		}
		if end {
			break
		}
	}

	fmt.Println(cnt)
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
