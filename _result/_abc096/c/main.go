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
	H := nextInt()
	W := nextInt()
	var sl []string

	for i:=0;i<H;i++ {
		sl = append(sl, nextLine())
	}

	for i,s := range sl {
		for j:=0;j<W;j++ {
			if s[j] == '#' {

				ok := false

				if j-1>=0 {
					if s[j-1] == '#' {ok = true}
				}

				if j+1<W {
					if s[j+1] == '#' {ok = true}
				}

				if i-1>=0 {
					if sl[i-1][j] == '#' {ok = true}
				}

				if i+1<H {
					if sl[i+1][j] == '#' {ok = true}
				}

				if !ok {
					fmt.Println("No")
					return
				}
			}
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
