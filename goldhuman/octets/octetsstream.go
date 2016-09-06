package octets

import (
"fmt"
)

const (
MAXSPARE = 16384
)

type OctetsStream struct {
	data []byte
	pos uint32
	tranpos uint32
}

func NewOctetsStream() *OctetsStream {
	b := make([]byte,0, 16)
	return &OctetsStream{b,0,0}
}

func NewOctetsStreamByData(data []byte) *OctetsStream {
	b := make([]byte, len(data), cap(data))
	copy(b, data)
	return &OctetsStream{b,0,0}
}

func NewOctetsStreamByStream(octetsStream *OctetsStream) *OctetsStream {
	b := make([]byte, len(octetsStream.data), cap(octetsStream.data))
	copy(b, octetsStream.data)
	pos := octetsStream.pos
	return &OctetsStream{b,pos,0}
}

func (octetsStream *OctetsStream) push_bool(num bool) {
	var numint int8
	if num == true {
		numint = 1
	}
	octetsStream.push_int8(numint)
}

func (octetsStream *OctetsStream) push_int8(num int8) {
	octetsStream.data = append(octetsStream.data, byte(num))
}

func (octetsStream *OctetsStream) push_int16(num int16) {
	octetsStream.data = append(octetsStream.data, byte(num >> 8))
	octetsStream.data = append(octetsStream.data, byte(num))
}

func (octetsStream *OctetsStream) push_int32(num int32) {
	octetsStream.data = append(octetsStream.data, byte(num >> 24))
	octetsStream.data = append(octetsStream.data, byte(num >> 16))
	octetsStream.data = append(octetsStream.data, byte(num >> 8))
	octetsStream.data = append(octetsStream.data, byte(num))
}

func (octetsStream *OctetsStream) push_int64(num int64) {
	octetsStream.data = append(octetsStream.data, byte(num >> 56))
	octetsStream.data = append(octetsStream.data, byte(num >> 48))
	octetsStream.data = append(octetsStream.data, byte(num >> 40))
	octetsStream.data = append(octetsStream.data, byte(num >> 32))
	octetsStream.data = append(octetsStream.data, byte(num >> 24))
	octetsStream.data = append(octetsStream.data, byte(num >> 16))
	octetsStream.data = append(octetsStream.data, byte(num >> 8))
	octetsStream.data = append(octetsStream.data, byte(num))
}

func (octetsStream *OctetsStream) compact_uint32(x int32) {
	if x < 0x40 {
		octetsStream.push_int8((int8)x)
		return
	} else if x < 0x4000 {
		octetsStream.push_int16((int16)(x|0x8000))
		return
	} else if x < 0x20000000 {
		octetsStream.push_int32(x|0xc0000000)
		return
	}
	octetsStream.push_int8((int8)0xe0)
	octetsStream.push_int32(x)
}

func (octetsStream *OctetsStream) compact_sint32(x int32) {
	if x >= 0 {
		if x < 0x40 {
			octetsStream.push_int8((int8)x)
			return
		} else if x < 0x2000 {
			octetsStream.push_int16((int16)(x|0x8000))
			return
		} else if x < 0x10000000 {
			octetsStream.push_int32(x|0xc0000000)
			return
		}
		octetsStream.push_int8((int8)0xe0)
		octetsStream.push_ine32(x)
	} else {
		x = -x
		if x < 0x40 {
			octetsStream.push_int8((int8)(x|0x40))
			return
		} else if x < 0x2000 {
			octetsStream.push_int16(int16)(x|0xa000))
			return
		} else if x < 0x10000000 {
			octetsStream.push_int32(x|0xd0000000)
			return
		}
		octetsStream.push_int8((int8)0xf0);
		octetsStream.push_int32(x)
	}
}

func (octetsStream *OctetsStream) pop_int8() (n int8, error) {
}

func (octetsStream *

func (octetsStream *OctetsStream) uncompact_uint32(x uint32) {
}
