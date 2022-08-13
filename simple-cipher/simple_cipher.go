package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
	shiftDistance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance <= 25 && distance >= -25 && distance != 0 {
		return shift{shiftDistance: distance}
	}
	return nil
}

func (c shift) Encode(input string) string {
	var sb strings.Builder
	input = strings.ToLower(input)
	for i := range input {
		if unicode.IsLetter(rune(input[i])) {
			sb.WriteByte('a' + (input[i]-'a'+byte(c.shiftDistance)+26)%26)
		}
	}
	return string(sb.String())
}

func (c shift) Decode(input string) string {
	var sb strings.Builder
	for i := range input {
		if unicode.IsLetter(rune(input[i])) {
			sb.WriteByte('a' + (input[i]-'a'-byte(c.shiftDistance)+26)%26)
		}
	}
	return string(sb.String())
}

func NewVigenere(key string) Cipher {
	count := 0
	for i := range key {
		if !unicode.IsLetter(rune(key[i])) {
			return nil
		}
		if !unicode.IsLower(rune(key[i])) {
			return nil
		}
		count += int(key[i] - 'a')
	}
	if count == 0 {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	var sb strings.Builder
	input = strings.ToLower(input)
	effectiveKey := v.key

	for i := range input {
		if unicode.IsLetter(rune(input[i])) {
			sb.WriteRune(rune(input[i]))
		}
	}
	effectiveInput := sb.String()
	sb.Reset()
	if len(effectiveInput) > len(v.key) {
		effectiveKey = strings.Repeat(effectiveKey, len(effectiveInput)/len(v.key))
		leftOver := len(effectiveInput) - (len(effectiveInput)/len(v.key))*len(v.key)
		if leftOver > 0 {
			effectiveKey += v.key[0:leftOver]
		}
	}
	for i := range effectiveInput {
		sb.WriteByte('a' + (effectiveInput[i]-'a'+byte(effectiveKey[i]-'a')+26)%26)
	}
	return string(sb.String())
}

func (v vigenere) Decode(input string) string {
	var sb strings.Builder
	input = strings.ToLower(input)
	effectiveKey := v.key
	if len(input) > len(v.key) {
		effectiveKey = strings.Repeat(effectiveKey, len(input)/len(v.key))
		leftOver := len(input) - (len(input)/len(v.key))*len(v.key)
		if leftOver > 0 {
			effectiveKey += v.key[0:leftOver]
		}
	}
	for i := range input {
		if unicode.IsLetter(rune(input[i])) {
			sb.WriteByte('a' + (input[i]-'a'-byte(effectiveKey[i]-'a')+26)%26)
		}
	}
	return string(sb.String())
}
