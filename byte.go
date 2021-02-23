package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type Byte []byte

func NewByteFromInt(i uint32) Byte {

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func NAByte() Byte {
	return nil
}

func (b Byte) ToInt() uint32 {
	return binary.BigEndian.Uint32(b)
}

func (b Byte) BinaryString() string {

	if b == nil {
		return "N/A"
	}

	var out []string
	for _, n := range b {
		out = append(out, fmt.Sprintf("%08b", n))
	}
	return strings.Join(out, ".")
}

func (b Byte) String() string {

	if b == nil {
		return "N/A"
	}

	var out []string
	for _, n := range b {
		out = append(out, fmt.Sprintf("%d", n))
	}
	return strings.Join(out, ".")
}
