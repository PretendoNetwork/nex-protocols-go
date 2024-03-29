// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// BufferQueueParam holds data for the DataStore (Super Mario Maker) protocol
type BufferQueueParam struct {
	nex.Structure
	DataID uint64
	Slot   uint32
}

// ExtractFromStream extracts a BufferQueueParam structure from a stream
func (bufferQueueParam *BufferQueueParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	bufferQueueParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.DataID from stream. %s", err.Error())
	}

	bufferQueueParam.Slot, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.Slot from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the BufferQueueParam and returns a byte array
func (bufferQueueParam *BufferQueueParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(bufferQueueParam.DataID)
	stream.WriteUInt32LE(bufferQueueParam.Slot)

	return stream.Bytes()
}

// Copy returns a new copied instance of BufferQueueParam
func (bufferQueueParam *BufferQueueParam) Copy() nex.StructureInterface {
	copied := NewBufferQueueParam()

	copied.SetStructureVersion(bufferQueueParam.StructureVersion())

	copied.DataID = bufferQueueParam.DataID
	copied.Slot = bufferQueueParam.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bufferQueueParam *BufferQueueParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BufferQueueParam)

	if bufferQueueParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if bufferQueueParam.DataID != other.DataID {
		return false
	}

	if bufferQueueParam.Slot != other.Slot {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (bufferQueueParam *BufferQueueParam) String() string {
	return bufferQueueParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (bufferQueueParam *BufferQueueParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BufferQueueParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, bufferQueueParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, bufferQueueParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %d,\n", indentationValues, bufferQueueParam.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() *BufferQueueParam {
	return &BufferQueueParam{
		DataID: 0,
		Slot:   0,
	}
}
