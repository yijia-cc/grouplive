package security

type CaesarCipher struct {
	offset int32
}

func (c CaesarCipher) Encrypt(input string) string {
	return shift(input, c.offset)
}

func (c CaesarCipher) Decrypt(input string) string {
	return shift(input, -c.offset)
}

func shift(input string, offset int32) string {
	if len(input) == 0 {
		return ""
	}

	output := make([]rune, 0)
	for _, curChar := range input {
		output = append(output, curChar+offset)
	}
	return string(output)
}

func NewCaesarCipher(offset int) CaesarCipher {
	return CaesarCipher{offset: int32(offset)}
}
