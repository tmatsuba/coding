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
	h := []int{}
	ret := make([]int, N)
	
	for i:=0; i<N; i++ {
		h = append(h, nextInt())
	}

	for i:=0; i<N; i++ {
		if i<1 {
			ret[i] = 0
			continue
		}
		if i<2 {
			ret[i] = int(math.Abs(float64(h[i])-float64(h[i-1])))
		} else {
			onesteps := math.Abs(float64(h[i])-float64(h[i-1]))
			twosteps := math.Abs(float64(h[i])-float64(h[i-2]))
			ret[i] = Min(ret[i-1]+int(onesteps), ret[i-2]+int(twosteps))
		}
	}
	fmt.Println(ret[len(ret)-1])
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

func nextFloat() float64 {
    sc.Scan()
    i, e := strconv.ParseFloat(sc.Text(), 64)
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
