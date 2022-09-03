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
	// sc.Split(bufio.ScanWords)

	S := nextLine()

	if S[0] == '1' {
		fmt.Println("No")
		return
	}

	// 4が倒れている時
	if S[3] == '0' {
		if S[6] == '1' && (S[1] == '1' || S[2] == '1' || S[4] == '1' || S[5] == '1' || S[7] == '1' || S[8] == '1' || S[9] == '1') {
			fmt.Println("Yes")
			return
		}
	} 
	if S[5] == '0' {
		if S[9] == '1' && (S[1] == '1' || S[2] == '1' || S[3] == '1' || S[4] == '1' || S[6] == '1' || S[7] == '1' || S[8] == '1') {
			fmt.Println("Yes")
			return
		}
	} 
	if S[4] == '0' {
		if (S[1] == '1' || S[3] == '1' || S[6] == '1' || S[7] == '1') && (S[2] == '1' || S[5] == '1' || S[8] == '1' || S[9] == '1') {
			fmt.Println("Yes")
			return
		}
	} 
	if S[1] == '0' && S[7] == '0' {
		if (S[3] == '1' || S[6] == '1') && (S[2] == '1' || S[4] == '1' || S[5] == '1' || S[8] == '1' || S[9] == '1') {
			fmt.Println("Yes")
			return
		}
	} 
	if S[2] == '0' && S[8] == '0' {
		if (S[5] == '1' || S[9] == '1') && (S[1] == '1' || S[3] == '1' || S[4] == '1' || S[6] == '1' || S[7] == '1') {
			fmt.Println("Yes")
			return
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
