// Package circular implements a circular buffer
package circular

import "errors"

// the 4th iteration of the test suite is in use
const testVersion = 4

// Buffer implements a bounded, circular buffer
type Buffer struct {
	read   int    // the next position from which we'll read data
	write  int    // the next position where we'll write data
	length int    // number of valid elements we have to read
	buf    []byte // our underlying byte store
}

// ErrEmptyBuffer is thrown when we try to read from a fallow buffer
var ErrEmptyBuffer = errors.New("circular : Attempted to read from an empty buffer")

// ErrFullBuffer is thrown when we try to write to a buffer at capacity
var ErrFullBuffer = errors.New("circular : Attempted to write to a full buffer")

// NewBuffer creates a circular buffer with capacity equal to size
func NewBuffer(size int) *Buffer {
	// if the size is not-positive, we should fail
	if size <= 0 {
		return nil
	}
	b := new(Buffer)
	//here we initialize the actual backing slice
	b.buf = make([]byte, size)
	return b
}

// Cap returns maximum capacity of this circular buffer
func (buf *Buffer) Cap() int {
	//simply we get the size of the underlying array
	return len(buf.buf)
}

// Len returns the number of occupied slots in the circular buffer
func (buf *Buffer) Len() int {
	return buf.length
}

// nextPos returns the next position in the circular buffer after x
func (buf *Buffer) nextPos(current int) int {
	//since this is a circular buffer, we need to increment, but also modulo in-case we have to wrap around
	return (current + 1) % buf.Cap()
}

// nextReadPos returns the next reading position in the buffer
func (buf *Buffer) nextReadPos() int {
	return buf.nextPos(buf.read)
}

// nextWritePos returns the next writing position in the buffer
func (buf *Buffer) nextWritePos() int {
	return buf.nextPos(buf.write)
}

// Reset effectively clears out the buffers, this is done by moving the read and write heads back to 0
func (buf *Buffer) Reset() {
	buf.read, buf.write = 0, 0
	buf.length = 0
}

// ReadByte gives back the next byte in the circular buffer and increments our read head if possible
func (buf *Buffer) ReadByte() (b byte, err error) {
	if buf.Len() <= 0 {
		err = ErrEmptyBuffer
		return
	}
	b = buf.buf[buf.read]
	buf.read = buf.nextReadPos()
	buf.length--
	return
}

// WriteByte inserts a new byte into the circular buffer if possible, errors if full
func (buf *Buffer) WriteByte(b byte) (err error) {
	if buf.Len() >= buf.Cap() {
		err = ErrFullBuffer
		return
	}
	buf.buf[buf.write] = b
	buf.write = buf.nextWritePos()
	buf.length++
	return
}

// Overwrite attempts to write normally, but if that's not possible, overwrites the oldest data we have
func (buf *Buffer) Overwrite(b byte) {
	// if we aren't at capacity, this translates to a simple write
	if buf.Len() < buf.Cap() {
		buf.WriteByte(b)
		return
	}
	//otherwise we need to do some magic i.e. we write anyway, but advance both the read and write heads
	buf.buf[buf.write] = b
	buf.write = buf.nextWritePos()
	buf.read = buf.nextReadPos()
}
