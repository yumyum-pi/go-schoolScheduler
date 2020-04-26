package utils

import "fmt"

// B256 store number in bytes
type B256 []byte

// Encode convert an int to b256 format
func (b *B256) Encode(i int) error {
	v := uint32(i)
	_ = (*b)[0] // to throw error when length is 0

	switch len(*b) {
	case 1: // 256
		if v > (1<<8 - 1) {
			return fmt.Errorf("> max no. allowed=%v", (1<<8 - 1))
		}
		(*b)[0] = byte(v)
	case 2: // 65536
		if v > (1<<16 - 1) {
			return fmt.Errorf("> max no. allowed=%v", (1<<16 - 1))
		}
		(*b)[1] = byte(v)
		(*b)[0] = byte(v >> 8)
	case 3: // 16777216
		if v > (1<<24 - 1) {
			return fmt.Errorf("> max no. allowed=%v", (1<<24 - 1))
		}
		(*b)[2] = byte(v)
		(*b)[1] = byte(v >> 8)
		(*b)[0] = byte(v >> 16)
	case 4: //4294967296
		(*b)[3] = byte(v)
		(*b)[2] = byte(v >> 8)
		(*b)[1] = byte(v >> 16)
		(*b)[0] = byte(v >> 24)
	}
	return nil
}

// Decode convert an b256 to int
func (b *B256) Decode() int {
	switch len(*b) {
	case 4:
		return int((*b)[3]) | int((*b)[2])<<8 | int((*b)[1])<<16 | int((*b)[0])<<32
	case 3:
		return int((*b)[2]) | int((*b)[1])<<8 | int((*b)[0])<<16
	case 2:
		return int((*b)[1]) | int((*b)[0])<<8
	case 1:
		return int((*b)[0])
	}

	return 0
}
