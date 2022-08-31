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

	N := nextInt()
	M := nextInt()
	a := []int{}
	ret := make([]int, N+1)

	for i:=0; i<M; i++ {
		a = append(a, nextInt())
	}

	for i:=0; i<N+1; i++ {
		
		if i == 0 {
			ret[i] = 1

		} else	if i == 1 {
			if indexOf(a, i) == -1 {
				ret[i] = 1
			}
		} else { 
			if indexOf(a, i) != -1 {
				ret[i] = 0
			}  else  if ret[i-1] == 0 && ret[i-2] == 0 {
				fmt.Println(0)
				return
			}  else  if ret[i-1] != 0 && ret[i-2] != 0 {
				ret[i] = (ret[i-1] + ret[i-2])%1000000007
			} else if ret[i-1] !=0 {
				ret[i] = ret[i-1]%1000000007
			} else if ret[i-2] != 0 {
				ret[i] = ret[i-2]%1000000007
			}
		}
	}
	fmt.Println(ret[len(ret)-1])
}

func indexOf(a []int, e int) int {
	for i:=0; i<len(a); i++ {

		if a[i] > e {
			a = a[:i]
			return -1
		}
		if a[i] == e {
			return i
		}
	}
	return -1
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
