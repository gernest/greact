package vdom

import (
	"fmt"
	"io"
)

// NewIndexedByteReader returns a new IndexedByteReader which will
// read from buf.
func NewIndexedByteReader(buf []byte) *IndexedByteReader {
	return &IndexedByteReader{buf: buf}
}

// IndexedByteReader satisfies io.Reader and io.ByteReader and also
// adds some additional methods for searching the buffer and returning
// the current offset.
type IndexedByteReader struct {
	buf []byte
	off int
}

// Read satisfies io.Reader.
func (r *IndexedByteReader) Read(p []byte) (int, error) {
	n := copy(p, r.buf[r.off:])
	r.off += n
	return n, nil
}

// ReadByte satisfies io.ByteReader. We expect that xml.Decoder will
// upgrade the IndexedByteReader to a io.ByteReader and call this method
// instead of Read.
func (r *IndexedByteReader) ReadByte() (byte, error) {
	if r.off >= len(r.buf) {
		// Reached the end of the buffer
		return 0, io.EOF
	}
	c := r.buf[r.off]
	r.off++
	return c, nil
}

// Offset returns the current offset position for r, i.e.,
// the number of bytes that have been read so far.
func (r *IndexedByteReader) Offset() int {
	return r.off
}

// BackwardsSearch starts at start and iterates backwards through r.buf[min:max]
// until it finds b. It returns the index of b if b was found within the given interval
// and -1 if it was not. It returns an error if min or max is outside the bounds of r.buf.
func (r *IndexedByteReader) BackwardsSearch(min int, max int, b byte) (int, error) {
	if min >= len(r.buf) || min < 0 {
		return -1, fmt.Errorf("Error in BackwardsSearch min %d is out of bounds. r has buf of length %d", min, len(r.buf))
	}
	if max >= len(r.buf) || max < min {
		return -1, fmt.Errorf("Error in BackwardsSearch max %d is out of bounds. r has buf of length %d and min was %d", min, len(r.buf), min)
	}
	for j := max; j >= min; j-- {
		if r.buf[j] == b {
			return j, nil
		}
	}
	return -1, nil
}
