package main

import (
	"bytes"
	"go/format"
	"math/rand"
	"os"
	"text/template"
	"time"
)

func main() {
	GenerateRandDefinesV2()
}

func GenerateRandDefinesV2() {
	raw := RandSeq62()

	tpl, err := template.ParseFiles("./utils/shorten/defines.go.tpl")
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"idEncodeArr": raw,
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		panic(err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(p)
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
