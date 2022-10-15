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

	X := nextInt()
	K := nextInt()

	for i:=0; i<K; i++ { 
		m := X % pow(10,i+1)/pow(10,i)
//		m := (X / pow(10, i)) *pow(10, i) - (X / pow(10, (i+1))) * pow(10, i+1)
 		// fmt.Println("m:", m)
 		// fmt.Println("X:", X)
// 		fmt.Println("src:", (X / pow(10, i)) * pow(10, i))
// 		fmt.Println("dst:", (X / pow(10, (i+1))) * pow(10, (i+1)))
		if m >= 5 {
			X = ((X / pow(10, (i+1))) + 1) * pow(10, (i+1))
		} else {
			X = ((X / pow(10, (i+1)))) * pow(10, (i+1))			
		}
	}

	fmt.Println(X)
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
