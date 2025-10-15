package encoder

import "github.com/yenoxx/voltkeep/hashes"

const MAGIC_1 uint64 = 123454321
const MAGIC_2 uint64 = 9797
const MAGIC_3 uint64 = 7979

type Encoder struct {
	key   uint64
	ckey  uint64
	bytes []byte
}

func CreateEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Begin(pass string, bytes []byte) *Encoder {
	e.key = hashes.DJB2(pass)

	nbytes := make([]byte, len(bytes))
	copy(nbytes, bytes)
	e.bytes = nbytes

	return e
}

func (e *Encoder) ChangePass(pass string) *Encoder {
	e.key = hashes.DJB2(pass)

	return e
}

func (e *Encoder) Encode() *Encoder {
	e.ckey = e.key
	for i := range len(e.bytes) {
		val, nkey := getByte(e.ckey)
		e.bytes[i] = wrapPositive(e.bytes[i], val)
		e.ckey = nkey
	}
	return e
}

func (e *Encoder) Decode() *Encoder {
	e.ckey = e.key
	for i := range len(e.bytes) {
		val, nkey := getByte(e.ckey)
		e.bytes[i] = wrapNegative(e.bytes[i], val)
		e.ckey = nkey
	}
	return e
}

func (e *Encoder) Bytes() []byte {
	nbytes := make([]byte, len(e.bytes))
	copy(nbytes, e.bytes)

	return nbytes
}

func getByte(ckey uint64) (byte, uint64) {
	val := ((ckey << 2) + (ckey * MAGIC_1 >> 2) + (ckey / MAGIC_2)) % 256
	nkey := (ckey << 1) + (ckey * MAGIC_1 >> 1) + (ckey / MAGIC_3)

	return byte(val), nkey
}

func wrapPositive(a byte, b byte) byte {
	c := int(a) + int(b)
	if c > 255 {
		c -= 256
	}
	return byte(c)
}

func wrapNegative(a byte, b byte) byte {
	c := int(a) - int(b)
	if c < 0 {
		c += 256
	}
	return byte(c)
}
