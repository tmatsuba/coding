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

	Ax, Ay := nextFloat(), nextFloat()
	Bx, By := nextFloat(), nextFloat()
	Cx, Cy := nextFloat(), nextFloat()
	Dx, Dy := nextFloat(), nextFloat()

	AB := []float64{Bx - Ax, By - Ay}
	BC := []float64{Cx - Bx, Cy - By}
	CD := []float64{Dx - Cx, Dy - Cy}
	DA := []float64{Ax - Dx, Ay - Dy}

	ABBC := AB[0] * BC[1] - AB[1] * BC[0]
	BCCD := BC[0] * CD[1] - BC[1] * CD[0]
	CDDA := CD[0] * DA[1] - CD[1] * DA[0]
	DAAB := DA[0] * AB[1] - DA[1] * AB[0]

	if ABBC > 0 && BCCD > 0 && CDDA > 0 && DAAB > 0 {
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

func nextFloat() float64 {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return float64(i)
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
