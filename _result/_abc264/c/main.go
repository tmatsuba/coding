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

	H1 := nextInt()
	W1 := nextInt()
	a := [][]int{}

	for i:=0; i<H1; i++ {
		h := []int{}
		for j:= 0; j<W1; j++ {
			h = append(h, nextInt())
		}
		a = append(a, h)
	}

	H2 := nextInt()
	W2 := nextInt()
	b := [][]int{}

	for i:=0; i<H2; i++ {
		h := []int{}
		for j:= 0; j<W2; j++ {
			h = append(h, nextInt())
		}
		b = append(b, h)
	}

	h3 := [][]int{}
	for bits:=0;bits<(1<<uint64(H1));bits++ {
		h := []int{}
		for i:=0;i<H1;i++ {
			if (bits>>uint64(i))&1 == 1 {
				h = append(h, i)
			}
		}

		if len(h) == H2 {
			h3 = append(h3, h)
		}
	}

	w3 := [][]int{}
	for bits:=0;bits<(1<<uint64(W1));bits++ {
		w := []int{}
		for i:=0;i<W1;i++ {
			if (bits>>uint64(i))&1 == 1 {
				w = append(w, i)
			}
		}

		if len(w) == W2 {
			w3 = append(w3, w)
		}
	}




	for _,h := range h3 {
		for _,w := range w3 {

			match := true
			for i,v1 := range h {
				for  j,v2 := range w {
					if a[v1][v2] != b[i][j] {
						match = false
						break
					}
				}
				if !match {
					break
				}
			}

			if match {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
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
