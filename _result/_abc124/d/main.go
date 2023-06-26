package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
	"math/bits"
	"strings"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, K := nextInt(),nextInt()
	s := nextLine()

	starts := []int{0}

	one, zero := false, false
	for i:=0; i<N; i++ {

		if s[i] == '0' {
			zero,one = true,false
		} else {
			if zero && !one && i != 0 {
				starts = append(starts, i)
			} 
			zero,one = false,true
		}
	}

	zero = false
	k := K
	cnt := 0
	end := false
	r:=0

//	fmt.Println(starts)
	for i:=0; i<len(starts); i++ {
		zero = false
		for j:=r; j<N; j++ {
			if j == N-1 {
				end = true
			}

			if s[j] == '0' && !zero {
				k--
				zero = true
			} else if s[j] == '1' {
				zero = false
			}

			if k < 0 {
//				fmt.Println(starts[i], j, "K: ", K, "k: ", k)
				if 1 <= K {
					k = 1
				} else  {
					k = K
				}
				break
			} else {
//				fmt.Println(starts[i], j, k, j-starts[i]+1)
				cnt = Max(cnt, j - starts[i] + 1)
				r = j
			}

		}
		if end {
			break
		}
	}


	// l, r := 0, 0
	// p := 1
	// ret := 0

	// for r < N {
	// 	p = p * s[r]

	// 	if p <= K {
	// 		ret = Max(ret, r-l+1)
	// 		r++
	// 		continue
	// 	} else {
	// 		for l < N {
	// 			p = p / s[l]

	// 			if p <= K {
	// 				l++
	// 				r++
	// 				break
	// 			}
				
	// 			if r <= l {
	// 				l++
	// 				r = l
	// 				p = 1
	// 				break
	// 			}
	// 			l++
	// 		}
	// 	}
	// }

	// fmt.Println(ret)


//	starts := []int{}

	// lone := false
	// zero := false
	// rone := false
	// lrs := [][]int{}



	// next_l := 0
	// lr := []int{}


	// for i:=0; i<N; i++ {
	// 	s := S[i]
	// 	if lone && rone {
	// 		if s == '0' {
	// 			lr = append(lr, i-1)
	// 			lrs = append(lrs, lr) 
	// 			lone, zero, rone = false, false, false
	// 			lr = []int{next_l}
	// 		} 
	// 	} else if lone && !rone {
	// 		if s == '1' && !zero{
	// 			rone = true
	// 			next_l = i
	// 		} else {
	// 			zero = true
	// 		}
	// 	} else if !lone && !rone {
	// 		lone = true
	// 		lr = append(lr, i)
	// 	} 

	// 	if i == N-1 {
	// 		lr = append(lr, i)
	// 		lrs = append(lrs, lr) 
	// 	}
	// }

	// fmt.Println(lrs)
	// fmt.Println(K)
	// 	if !zero {
	// 		if s == '0' {
 	// 			sum[i+1] = sum[i] + 1
	// 			ret = Max(ret, cnt)
	// 			cnt = 0
	// 			zero = true
	// 		} else {
	// 			cnt++
	// 			sum[i+1] = sum[i]
	// 			zero = false
	// 		}
	// 	} else if zero {
	// 		if s == '1' {
	// 			starts = append(starts, i)
	// 			cnt++
	// 			zero = false
	// 		} else {
	// 			ret = Max(ret, cnt)
	// 			cnt = 0
	// 		}
	// 		sum[i+1] = sum[i]
	// 	}
	// }

	// if K == 0 {
	// 	fmt.Println(ret)
	// 	return
	// }


	// i,  next := 1, 1
	// v := 0
	
	// for i<N+1 {
	// 	if i != 1 {
	// 		v = sum[i]
	// 	}
	// 	cnt := 0
	// 	for j:=i; j<N+1; j++ {

	// 		if (i != 0 && i == next && v != sum[j]) {
	// 			next = j
	// 		} else if j == N {
	// 			next = j + 1
	// 		}

	// 		if sum[j] - v > K {
	// 			break
	// 		}
	// 		cnt++
	// 		fmt.Println("i:", i, "j:", j, "cnt: ", cnt)
	// 	}


	// 	ret = Max(ret, cnt)
	// 	i = next
	// }
	
	// for i:=1; i<len(sum); i++{
	// 	if i == 0 {continue}
	// 	if i + 1 >= len(sum) {break}

	// 	if sum[i] < sum[i+1] {
	// 		starts = append(starts, i)
	// 	}
	// }
//	fmt.Println(sum)
//	fmt.Println(starts)

	// for i, start := range(starts) {
	// 	end := start
		


	// 	for j:=i; j<len(starts); j++ {
	// 		if starts[j] > 0  {
	// 			end = starts[j] - 1
	// 		} else {
	// 			end = start
	// 		}
	// 		if sum[starts[j]] - sum[start] > K {
	// 			ret = Max(ret, end - start)
	// 			break
	// 		}
	// 	}
	// }


	// start := 0
	// end := 0
	// next := 0
	// for end <= N {
	// 	start = next
	// 	zero := false
	// 	k := K
	// 	for i:=start; i<N; i++ {
			
	// 		if S[i] == '1' {
	// 			if zero {
	// 				k--
	// 				if start == next {
	// 					next = i
	// 				}
	// 			}
	// 			zero = false
	// 		} else {
	// 			if !zero && k <= 0 {
	// 				break
	// 			}

	// 			zero = true
	// 		}
	// 		end = i + 1
	// 	}
	// 	ret = Max(ret, end - start)

	// 	if start == next {
	// 		break
	// 	}
	// }

	fmt.Println(cnt)
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

func ArrayJoin(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
