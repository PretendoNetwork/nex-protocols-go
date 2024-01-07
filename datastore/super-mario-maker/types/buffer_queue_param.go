// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// BufferQueueParam holds data for the DataStore (Super Mario Maker) protocol
type BufferQueueParam struct {
	types.Structure
	DataID *types.PrimitiveU64
	Slot   *types.PrimitiveU32
}

// ExtractFrom extracts the BufferQueueParam from the given readable
func (bufferQueueParam *BufferQueueParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = bufferQueueParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BufferQueueParam header. %s", err.Error())
	}

	err = bufferQueueParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.DataID from stream. %s", err.Error())
	}

	err = bufferQueueParam.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.Slot from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the BufferQueueParam to the given writable
func (bufferQueueParam *BufferQueueParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bufferQueueParam.DataID.WriteTo(contentWritable)
	bufferQueueParam.Slot.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bufferQueueParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of BufferQueueParam
func (bufferQueueParam *BufferQueueParam) Copy() types.RVType {
	copied := NewBufferQueueParam()

	copied.StructureVersion = bufferQueueParam.StructureVersion

	copied.DataID = bufferQueueParam.DataID.Copy().(*types.PrimitiveU64)
	copied.Slot = bufferQueueParam.Slot.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bufferQueueParam *BufferQueueParam) Equals(o types.RVType) bool {
	if _, ok := o.(*BufferQueueParam); !ok {
		return false
	}

	other := o.(*BufferQueueParam)

	if bufferQueueParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !bufferQueueParam.DataID.Equals(other.DataID) {
		return false
	}

	if !bufferQueueParam.Slot.Equals(other.Slot) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, bufferQueueParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, bufferQueueParam.DataID))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, bufferQueueParam.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() *BufferQueueParam {
	return &BufferQueueParam{
		DataID: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU32(0),
	}
}
