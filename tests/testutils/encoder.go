package testutils

import "encoding/binary"

func EncodeHName(hname uint32) []byte {
	encoded_hname := make([]byte, 4)
	binary.LittleEndian.PutUint32(encoded_hname, hname)
	return encoded_hname
}
