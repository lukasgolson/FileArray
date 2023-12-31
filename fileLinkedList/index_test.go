package fileLinkedList

import (
	"encoding/binary"
	"github.com/lukasgolson/FileArray/serialization"
	"testing"
)

func TestIndexSerialization(t *testing.T) {
	// Create a sample indexEntry
	itemID := serialization.Offset(1)
	offset := serialization.Offset(2)
	length := serialization.Length(3)
	index := newIndexEntry(itemID, offset, length)

	// Serialize the indexEntry
	buffer := make([]byte, index.StrideLength())
	err := index.SerializeToBinaryStream(buffer)
	if err != nil {
		t.Fatalf("Serialization failed: %v", err)
	}

	// Deserialize the serialized data
	deserializedIndex, err := index.DeserializeFromBinaryStream(buffer)
	if err != nil {
		t.Fatalf("Deserialization failed: %v", err)
	}

	// Compare the original indexEntry with the deserialized indexEntry
	if index != deserializedIndex {
		t.Errorf("Expected %v, got %v", index, deserializedIndex)
	}
}

func TestIndexStrideLength(t *testing.T) {
	// Create a sample indexEntry
	itemID := serialization.Offset(1)
	offset := serialization.Offset(2)
	length := serialization.Length(3)
	index := newIndexEntry(itemID, offset, length)

	// Calculate the expected stride length manually
	expectedStride := serialization.Length(binary.Size(itemID) + binary.Size(offset) + binary.Size(length))

	// Get the stride length from the indexEntry
	stride := index.StrideLength()

	// Compare the calculated stride with the expected stride
	if stride != expectedStride {
		t.Errorf("Expected stride length %d, got %d", expectedStride, stride)
	}
}
