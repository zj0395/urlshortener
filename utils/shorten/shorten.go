package shorten

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
		res += charMap[code[i]] * pow
		pow *= idStrLen
	}
	res = intRecover(res)
	return res
}

// intConvert convert a int to another
func intConvert(raw int64) int64 {
	val := raw & 0x0fffffff00000000
	// only convert int32
	val += (raw & 0xff000000) >> 16
	val += (raw & 0x00ff0000) >> 16
	val += (raw & 0x0000ff00) << 8
	val += (raw & 0x000000ff) << 24
	return val
}

// intRecover undo intConvert
func intRecover(val int64) int64 {
	raw := val & 0x0fffffff00000000

	raw += (val & 0xff000000) >> 24
	raw += (val & 0x00ff0000) >> 8
	raw += (val & 0x000000ff) << 16
	raw += (val & 0x0000ff00) << 16
	return raw
}
