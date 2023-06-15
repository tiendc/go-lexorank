package lexorank

func incrementChar(char byte) (byte, bool) {
	if char == 'z' {
		return 0, false
	}
	if char == '9' {
		return 'A', true
	}
	if char == 'Z' {
		return 'a', true
	}
	return char + 1, true
}

func incrementChars(chars []byte) []byte {
	for i := len(chars) - 1; i >= 0; i-- {
		next, ok := incrementChar(chars[i])
		if ok {
			chars[i] = next
			return chars
		} else {
			chars[i] = '1'
		}
	}
	return append(chars, '1')
}

func decrementChar(char byte) (byte, bool) {
	if char == '0' {
		return 0, false
	}
	if char == 'a' {
		return 'Z', true
	}
	if char == 'A' {
		return '9', true
	}
	return char - 1, true
}
