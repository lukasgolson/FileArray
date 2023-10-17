package internal

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"github.com/lukasgolson/FileArray/serialization"
	"reflect"
)

type NonSliceType interface {
	~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~complex64 | ~complex128 | ~string | ~rune
}

type GenericSerializer[T NonSliceType] struct {
	data T
}

func (g GenericSerializer[T]) SerializeToBinaryStream(buffer []byte) error {
	bufferSize := len(buffer)
	requiredSize := int(g.StrideLength())

	// Check buffer size
	if bufferSize < requiredSize {
		return errors.New("buffer size is too small")
	}

	buf := bytes.NewBuffer(buffer[:0])
	if err := binary.Write(buf, binary.LittleEndian, g.data); err != nil {
		return err
	}

	return nil
}

func (g GenericSerializer[T]) DeserializeFromBinaryStream(buffer []byte) (GenericSerializer[T], error) {
	bufferSize := len(buffer)
	requiredSize := int(g.StrideLength())

	// Check buffer size
	if bufferSize < requiredSize {
		return GenericSerializer[T]{}, errors.New("buffer size is too small")
	}

	buf := bytes.NewReader(buffer)
	var data T
	if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
		return GenericSerializer[T]{}, err
	}

	return GenericSerializer[T]{data: data}, nil
}

func (g GenericSerializer[T]) StrideLength() serialization.Length {
	return serialization.Length(binary.Size(g.data))
}

func (g GenericSerializer[T]) IDByte() byte {
	typeString := reflect.TypeOf(g.data).String()
	hasher := sha1.New()
	hasher.Write([]byte(typeString))
	hash := hasher.Sum(nil)
	return hash[0]
}
