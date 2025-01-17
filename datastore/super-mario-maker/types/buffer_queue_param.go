// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BufferQueueParam is a type within the DataStore protocol
type BufferQueueParam struct {
	types.Structure
	DataID types.UInt64
	Slot   types.UInt32
}

// WriteTo writes the BufferQueueParam to the given writable
func (bqp BufferQueueParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bqp.DataID.WriteTo(contentWritable)
	bqp.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bqp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BufferQueueParam from the given readable
func (bqp *BufferQueueParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = bqp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam header. %s", err.Error())
	}

	err = bqp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.DataID. %s", err.Error())
	}

	err = bqp.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BufferQueueParam
func (bqp BufferQueueParam) Copy() types.RVType {
	copied := NewBufferQueueParam()

	copied.StructureVersion = bqp.StructureVersion
	copied.DataID = bqp.DataID.Copy().(types.UInt64)
	copied.Slot = bqp.Slot.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given BufferQueueParam contains the same data as the current BufferQueueParam
func (bqp BufferQueueParam) Equals(o types.RVType) bool {
	if _, ok := o.(BufferQueueParam); !ok {
		return false
	}

	other := o.(BufferQueueParam)

	if bqp.StructureVersion != other.StructureVersion {
		return false
	}

	if !bqp.DataID.Equals(other.DataID) {
		return false
	}

	return bqp.Slot.Equals(other.Slot)
}

// CopyRef copies the current value of the BufferQueueParam
// and returns a pointer to the new copy
func (bqp BufferQueueParam) CopyRef() types.RVTypePtr {
	copied := bqp.Copy().(BufferQueueParam)
	return &copied
}

// Deref takes a pointer to the BufferQueueParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (bqp *BufferQueueParam) Deref() types.RVType {
	return *bqp
}

// String returns the string representation of the BufferQueueParam
func (bqp BufferQueueParam) String() string {
	return bqp.FormatToString(0)
}

// FormatToString pretty-prints the BufferQueueParam using the provided indentation level
func (bqp BufferQueueParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BufferQueueParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, bqp.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, bqp.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() BufferQueueParam {
	return BufferQueueParam{
		DataID: types.NewUInt64(0),
		Slot:   types.NewUInt32(0),
	}

}
