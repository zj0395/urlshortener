package shorten

// This file was auto-generated by `meta/random_seq.go`,
// you can change it to your own seq by running
// `go run meta/random_seq.go > utils/shorten/defines.go`

const (
	idEncodeStr = "{{printf "%s" .idEncodeArr}}"
	idStrLen    = {{len .idEncodeArr}}
)

var idCharMap = map[byte]int64{
{{range $index, $element := .idEncodeArr}} {{printf "\t%q: %d,\n" $element $index}} {{end}}
}
