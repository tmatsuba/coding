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
	c11, c12, c13, c21, c22, c23, c31, c32, c33 := nextInt(),nextInt(),nextInt(),nextInt(),nextInt(),nextInt(),nextInt(),nextInt(),nextInt()
	ret := "No"
	var a1,a2,a3,b1,b2,b3 int

	m := Max(Max(Max(Max(Max(Max(Max(Max(c11, c12),c13), c21),c22,),c23),c31),c32),c33)
		
	for a1=0;a1<=m;a1++ {
		b1 = c11-a1
		b2 = c12-a1
		b3 = c13-a1
		a2 = c21-b1
		a3 = c31-b1

		if a2+b2==c22 && a2+b3==c23 && a3+b2==c32 && a3+b3==c33 {
			ret = "Yes"
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
