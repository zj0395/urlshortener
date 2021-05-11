package shorten

const (
	minNum = 62*62*62*62*62 + 1              // bigger than 62^5
	maxNum = 62*62*62*62*62*62 - 0x3ffffffff // less than 62^6
)

// IDShorten Convert int64 to number base62
func IDShorten(id int64) string {
	id = intConvert(id)
	res := make([]byte, 0, 6)
	for id > 0 {
		res = append(res, idEncodeStr[id%idStrLen])
		id /= idStrLen
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

// IDRecover Convert from string to int64
func IDRecover(code string) int64 {
	var res int64
	var pow int64 = 1
	for i := len(code) - 1; i >= 0; i-- {
		res += idCharMap[code[i]] * pow
		pow *= idStrLen
	}
	res = intRecover(res)
	return res
}

// intConvert convert a int to another
func intConvert(raw int64) int64 {
	val := raw & 0x0ffffffc00000000
	// only convert int34 0x3ffffffff
	val += (raw & 0x3ff000000) >> 16
	val += (raw & 0x000ff0000) >> 16
	val += (raw & 0x00000ff00) << 10
	val += (raw & 0x0000000ff) << 26

	val += minNum
	return val
}

// intRecover undo intConvert
func intRecover(val int64) int64 {
	val -= minNum
	raw := val & 0x0ffffffc00000000

	raw += (val & 0x3fc000000) >> 26
	raw += (val & 0x003fc0000) >> 10
	raw += (val & 0x0000000ff) << 16
	raw += (val & 0x00003ff00) << 16
	return raw
}
