package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

type Buffer struct {
	size, writeIndex int
	data             []byte
}

func NewBuffer(capacity int) *Buffer {
	return &Buffer{data: make([]byte, capacity)}
}

// The question doesn't say but reading a byte frees up that slot.
func (b *Buffer) ReadByte() (byte, error) {
	if b.isEmpty() {
		return 0, errors.New("empty buffer")
	}
	val := b.data[b.readIndex()]
	b.size--
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isFull() {
		return errors.New("full buffer")
	}
	b.data[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % cap(b.data)
	b.size++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.isFull() {
		b.writeIndex = b.readIndex()
		b.size--
	}
	_ = b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.size = 0
}

// Go remainder (%) operator returns a result with the same sign as the dividend,
// which is different from the mathematical definition of a modulus, which always
// returns a non-negative result.
func mod(a, b int) int {
	return (a%b + b) % b
}

func (b *Buffer) readIndex() int {
	return mod(b.writeIndex-b.size, cap(b.data))
}

func (b *Buffer) isFull() bool {
	return b.size == cap(b.data)
}

func (b *Buffer) isEmpty() bool {
	return b.size == 0
}
