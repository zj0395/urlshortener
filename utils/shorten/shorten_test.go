package shorten

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIntConvert(t *testing.T) {
	check := func(testNum int64, msg string) {
		midVal := intConvert(testNum)
		res := intRecover(midVal)
		if testNum != res {
			t.Fatalf("%s intRecover(intConvert(0x%x)) = 0x%x, excepted:0x%x, midVal:0x%x", msg, testNum, res, testNum, midVal)
		}
	}

	// base
	f := []int64{0x7ff000000, 0x700ff0000, 0x70000ff00, 0, maxNum}
	for _, v := range f {
		check(v, "base")
	}

	testCnt := 10 * 1000 * 1000
	rand.Seed(time.Now().Unix())

	for testCnt > 0 {
		testCnt--
		number := rand.Int63n(maxNum)
		check(number, "random")
	}
}

func TestIDShorten(t *testing.T) {
	t.Logf("min:%d, max:%d, maxNum:%d", 0, maxNum, maxNum)
	testCnt := 1000 * 1000

	rand.Seed(time.Now().Unix())

	for testCnt > 0 {
		testCnt--

		randNum := rand.Int63n(maxNum)
		middleVal := IDShorten(randNum)
		res := IDRecover(middleVal)
		if len(middleVal) != 6 {
			t.Fatalf("len(IDShorten(%d)) != 6, val:%s", randNum, middleVal)
		}
		if res != randNum {
			t.Fatalf("IDRecover(IDShorten(%d)) = %d, excepted:%d, middle:%s\n", randNum, res, randNum, middleVal)
		}
	}
}

func BenchmarkIDShorten(t *testing.B) {
	for i := 0; i < t.N; i++ {
		IDRecover(IDShorten(0x7ffffffff))
	}
}

func TestIDIncr(t *testing.T) {
	enable := true

	if !enable {
		return
	}

	rand.Seed(time.Now().Unix())

	var loopCnt int64 = 1000
	const oneLoopCnt int64 = 1000
	res := [7]int{}

	for loopCnt > 0 {
		loopCnt--
		var begin int64 = rand.Int63n(maxNum - oneLoopCnt)
		var end int64 = begin + oneLoopCnt

		before, cur := IDShorten(begin-1), ""

		for begin < end {
			cur = IDShorten(begin)
			begin++

			diffCnt := 0
			for i := 0; i < len(before); i++ {
				if before[i] != cur[i] {
					diffCnt++
				}
			}
			// t.Logf("%d, %s, %s", begin, cur, before)
			before = cur

			res[diffCnt]++
		}
	}

	str := "When number incr by one, the string change\n"
	str += "\tchangeCharNum\tcnt\n"
	for i := 1; i < len(res); i++ {
		str += fmt.Sprintf("\t%d\t%d\t\n", i, res[i])
	}
	t.Logf(str)
}
