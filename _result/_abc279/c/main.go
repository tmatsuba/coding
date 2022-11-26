package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	H := nextInt()
	W := nextInt()

	SWi := make([][]rune, W)
	TWi := make([][]rune, W)

	for i := range SWi {
		SWi[i] = make([]rune, H)
	}

	for i := range TWi {
		TWi[i] = make([]rune, H)
	}

	for i:=0; i<H; i++ {
		Si := nextLine()

		for j:=0; j<W; j++ {
			SWi[j][i] = rune(Si[j])
		}
	}

	for i:=0; i<H; i++ {
		Ti := nextLine()

		for j:=0; j<W; j++ {
			TWi[j][i] = rune(Ti[j])
		}
	}

	arr_swi := make([]string, W)
	for i, s := range SWi {
		arr_swi[i] = string(s)
	}

	arr_twi := make([]string, W)
	for i, s := range TWi {
		arr_twi[i] = string(s)
	}

	sort.Strings(arr_swi)
	sort.Strings(arr_twi)

	for i:=0; i<W; i++ {
		if arr_swi[i] != arr_twi[i] {
			fmt.Println("No")
			return
		}
	}

 	fmt.Println("Yes")
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
