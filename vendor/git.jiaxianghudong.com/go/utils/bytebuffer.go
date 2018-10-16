package utils

// 加入8字节
func Put8bit(buf []byte, n byte) []byte {
	return append(buf, n)
}

// 加入16字节
func Put16bit(buf []byte, n uint16) []byte {
	var by [2]byte

	by[0] = byte((n >> 8) & 0xff)
	by[1] = byte(n & 0xff)

	return append(buf, by[:]...)
}

// 加入32字节
func Put32bit(buf []byte, n uint32) []byte {
	var by [4]byte

	by[0] = byte((n >> 24) & 0xff)
	by[1] = byte((n >> 16) & 0xff)
	by[2] = byte((n >> 8) & 0xff)
	by[3] = byte(n & 0xff)

	return append(buf, by[:]...)
}

// 加入64字节
func Put64bit(buf []byte, n uint64) []byte {
	var by [8]byte

	by[0] = byte((n >> 56) & 0xff)
	by[1] = byte((n >> 48) & 0xff)
	by[2] = byte((n >> 40) & 0xff)
	by[3] = byte((n >> 32) & 0xff)
	by[4] = byte((n >> 24) & 0xff)
	by[5] = byte((n >> 16) & 0xff)
	by[6] = byte((n >> 8) & 0xff)
	by[7] = byte(n & 0xff)

	return append(buf, by[:]...)
}

// 获取8bit
func Get8bit(buf []byte, start int) byte {
	return buf[start]
}

// 获取16bit
func Get16bit(buf []byte, start int) uint16 {
	var ret uint16

	ret = uint16(buf[start]) << 8
	ret |= uint16(buf[start+1])

	return ret
}

// 获取32big
func Get32bit(buf []byte, start int) uint32 {
	var ret uint32

	ret = uint32(buf[start]) << 24
	ret |= uint32(buf[start+1]) << 16
	ret |= uint32(buf[start+2]) << 8
	ret |= uint32(buf[start+3])

	return ret
}

// 获取64bit
func Get64bit(buf []byte, start int) uint64 {
	var ret uint64

	ret = uint64(buf[start]) << 56
	ret |= uint64(buf[start+1]) << 48
	ret |= uint64(buf[start+2]) << 40
	ret |= uint64(buf[start+3]) << 32
	ret |= uint64(buf[start+4]) << 24
	ret |= uint64(buf[start+5]) << 16
	ret |= uint64(buf[start+6]) << 8
	ret |= uint64(buf[start+7])

	return ret
}
