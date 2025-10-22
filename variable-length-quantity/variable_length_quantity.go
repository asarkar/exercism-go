package variablelengthquantity

import (
	"errors"
	"slices"
)

func EncodeVarint(input []uint32) []byte {
	res := make([]byte, 0)
	for _, i := range input {
		res = append(res, encode(i)...)
	}
	return res
}

func encode(i uint32) []byte {
	if i == 0 {
		return []byte{0}
	}
	res := make([]byte, 0)
	for i > 0 {
		// Extract the lowest 7 bits, and set the MSB (continuation bit).
		b := byte(i&0x7F | 0x80)
		i >>= 7
		res = append(res, b)
	}
	// Clear the MSB on the last byte.
	res[0] &= (0x80 - 1)
	slices.Reverse(res)

	return res
}

func DecodeVarint(input []byte) ([]uint32, error) {
	res := make([]uint32, 0)
	buf := make([]byte, 0)
	for _, b := range input {
		buf = append(buf, b)
		if b&0x80 == 0 {
			res = append(res, decode(buf))
			buf = []byte{}
		}
	}
	if len(buf) > 0 {
		return nil, errors.New("incomplete sequence")
	}
	return res, nil
}

func decode(input []byte) uint32 {
	var res uint32
	for _, b := range input {
		res <<= 7
		res |= uint32(b & 0x7F)
	}
	return res
}
