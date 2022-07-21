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
	sc.Split(bufio.ScanWords)
	N := nextInt()
	Y := nextInt()

	a := 10000
	b := 5000
	c := 1000

	var ac,	bc,	cc int

	exact := false
	for ac=0;ac<=N;ac++ {
		for bc=0;ac+bc<=N;bc++{
			cc = N-ac-bc
			if ac * a + bc * b + cc * c == Y {
				exact = true
				break
			}
		}
		if exact {break}
	}

	if exact {
		fmt.Printf("%v %v %v\n", ac, bc, cc)
	} else {
		fmt.Println("-1 -1 -1")
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
