package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
)

func main() {
	N := nextLine()
	ret := math.MaxInt64

	for bits:=0;bits<(1<<uint64(len(N)));bits++ {
		// bitsの各要素についてのループ
		sum := 0
		on := 0
		for i:=0;i<len(N);i++ {
			// bitsのi個目の要素を右シフトしつつ状態がonかどうかAND演算でチェック
			if (bits>>uint64(i))&1 == 1 {
				sum += int(N[i])
				on++
			}
		}
		if on > 0 && sum % 3 == 0 {
			ret = Min(ret, len(N) - on)
		}
    }

	if ret == math.MaxInt64 {ret = -1}
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
