package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcounts(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopcountsLoop(x uint64) int {
	var cnt byte

	for i := range pc {
		cnt += pc[byte(x>>(i*8))]
	}

	return int(cnt)
}

func PopcountsShift(x uint64) int {
	var cnt int

	for i := 0; i < 64; i++ {
		if byte(x>>i) == 1 {
			cnt++
		}
	}

	return cnt
}
