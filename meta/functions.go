package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	GenerateRandDefines()
}

func GenerateRandDefines() {
	raw := RandSeq62()

	str := fmt.Sprintf("package shorten\n\n")
	str += fmt.Sprintf("// you can replace all the config by a random one generate by meta/functions.go\n\n")
	str += fmt.Sprintf("const (\n")
	str += fmt.Sprintf("\tidEncodeStr = \"%s\"\n", string(raw))
	str += fmt.Sprintf("\tidStrLen    = %d\n", len(raw))
	str += fmt.Sprintf(")\n\n")
	str += fmt.Sprintf("var charMap = map[byte]int64{\n")
	for i := 0; i < len(raw); i++ {
		str += fmt.Sprintf("\t'%c': %d,\n", raw[i], i)
	}
	str += fmt.Sprintf("}\n")

	fmt.Println(str)
}

// RandSeq62 generate random sequences of 62 chars
func RandSeq62() []byte {
	raw := []byte{}
	for i := 0; i < 10; i++ {
		raw = append(raw, byte('0'+i))
	}
	for i := 0; i < 26; i++ {
		raw = append(raw, byte('a'+i))
	}
	for i := 0; i < 26; i++ {
		raw = append(raw, byte('A'+i))
	}

	rand.Seed(time.Now().Unix())
	for i := 1; i < len(raw); i++ {
		randIdx := rand.Intn(i)
		raw[i], raw[randIdx] = raw[randIdx], raw[i]
	}

	return raw
}
