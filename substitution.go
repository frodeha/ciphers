package main

import "fmt"

type SubstitutionCipher struct {
}

func NewSubstitutionCipher() SubstitutionCipher {
	return SubstitutionCipher{}
}

func (c SubstitutionCipher) Encrypt(s string) (string, error) {
	out := ""
	for idx, b := range s {
		if b == 32 {
			continue
		}

		if b > 122 || b < 65 {
			return "", fmt.Errorf("Found invalid character %#U at position %d", b, idx)
		}

		encodedByte := b + 4
		if encodedByte > 122 {
			encodedByte = 64 + encodedByte - 122
		}

		out = fmt.Sprintf("%s%s", out, string(encodedByte))
	}

	return out, nil
}

func (c SubstitutionCipher) Decrypt(s string) (string, error) {
	out := ""
	for idx, b := range s {
		if b == 32 {
			continue
		}

		if b > 122 || b < 65 {
			return "", fmt.Errorf("Found invalid character %#U at position %d", b, idx)
		}

		encodedByte := b - 4
		if encodedByte < 65 {
			encodedByte = 123 - (65 - encodedByte)
		}

		out = fmt.Sprintf("%s%s", out, string(encodedByte))
	}

	return out, nil
}
