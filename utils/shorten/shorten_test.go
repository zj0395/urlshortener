package shorten

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const minNum = 0xffffffff + 1
const maxNum = 62*62*62*62*62*62 - 0xffffffff
const numbers = maxNum - minNum

func TestIntConvert(t *testing.T) {
	testCnt := 10 * 1000 * 1000

	for testCnt > 0 {
		testCnt--
		randNum := rand.Int63n(numbers) + minNum
		midVal := intConvert(randNum)
		res := intRecover(midVal)
		if randNum != res {
			t.Fatalf("intRecover(intConvert(%d)) = %d, excepted:%d, midVal:%d", randNum, res, randNum, midVal)
		}
	}
}

func TestIDShorten(t *testing.T) {
	t.Logf("min:%d, max:%d, numbers:%d", minNum, maxNum, numbers)
	testCnt := 1000 * 1000

	rand.Seed(time.Now().Unix())

	for testCnt > 0 {
		testCnt--

		randNum := rand.Int63n(numbers) + minNum
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
	IDRecover(IDShorten(1111111111111))
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
		var begin int64 = minNum + rand.Int63n(numbers-oneLoopCnt)
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
