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

	_, M := nextInt(), nextInt()
	ret := 0
	m := map[int][]int{}

	for i:=0; i<M; i++ {
		key := nextInt()
		_, ok := m[key]
		if ok {
			m[key] = append(m[key], nextInt())
		} else {
			m[key] = []int{nextInt()}
		}
	}

	for a,v1 := range m {
		for _,b := range v1 {
			v2, ok := m[b] 
			if ok {
				for _, c := range v2 {
					x, ok1 := m[a]
					y, ok2 := m[c]
					if ok1 {
						for _, v := range x {
							if v == c {
								ret++
								break
							}
						}
					}
					if ok2 {
						for _, v := range y {
							if v == a {
								ret++
								break
							}
						}
					}

				}
			}
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
