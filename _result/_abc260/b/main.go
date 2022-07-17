package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
	"sort"
)

func main() {
	sc.Split(bufio.ScanWords)
	
	N := nextInt()
	X:= nextInt()
	Y := nextInt()
	Z := nextInt()

	A := map[int]int{}
	B := map[int]int{}
	C := map[int]int{}

	ret := []int{}

	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	
	a := List{}
    for k, v := range A {
        e := Entry{k, v}
        a = append(a, e)
    }
	sort.Sort(a)

	for i := 0; i < X; i++ {
		ret = append(ret, a[i].id)
	}

	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}

	b := List{}
    for k, v := range B {
        e := Entry{k, v}
        b = append(b, e)
    }
	sort.Sort(b)

	for i := 0; i < N; i++ {
		skip := false
		for _, v := range ret {
			if v == b[i].id {skip = true}
		}
		if skip {continue} 
		if Y <= 0 {break}
		ret = append(ret, b[i].id)
		Y--
	}

	for i := 0; i < N; i++ {
		C[i] = A[i] + B[i]
	}

	c := List{}
    for k, v := range C {
        e := Entry{k, v}
        c = append(c, e)
    }


	sort.Sort(c)

	for i := 0; i < N;i++ {
		skip := false
		for _, v := range ret {
			if v == c[i].id {skip = true}
 		}
		if skip {continue} 
		if Z <= 0 {break}
		ret = append(ret, c[i].id)
		Z--
	}

	sort.Ints(ret)

	for _, v := range ret {
		fmt.Println(v + 1)
	}
}

type Entry struct {
    id  int
    value int
}

type List []Entry

func (l List) Len() int {
    return len(l)
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
    if l[i].value == l[j].value {
        return (l[i].id < l[j].id)
    } else {
        return (l[i].value > l[j].value)
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
