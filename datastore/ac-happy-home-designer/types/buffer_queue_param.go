// Package types implements all the types used by the DataStoreACHappyHomeDesigner protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BufferQueueParam is a type within the DataStoreACHappyHomeDesigner protocol
type BufferQueueParam struct {
	types.Structure

	DataId *types.PrimitiveU64
	Slot   *types.PrimitiveU32
}

// WriteTo writes the BufferQueueParam to the given variable
func (bqp *BufferQueueParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bqp.DataId.WriteTo(contentWritable)
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

	err = bqp.DataId.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.DataId. %s", err.Error())
	}

	err = bqp.Slot.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreFileServerGetObject
func (bqp *BufferQueueParam) Copy() types.RVType {
	copied := NewBufferQueueParam()

	copied.DataId = bqp.DataId
	copied.Slot = bqp.Slot

	return copied
}

// Equals checks if the given BufferQueueParam contains the same data as the current BufferQueueParam
func (bqp *BufferQueueParam) Equals(o types.RVType) bool {
	if _, ok := o.(*BufferQueueParam); !ok {
		return false
	}

	other := o.(*BufferQueueParam)

	if !bqp.DataId.Equals(other.DataId) {
		return false
	}

	return bqp.Slot.Equals(other.Slot)
}

// String returns the string representation of the BufferQueueParam
func (bqp *BufferQueueParam) String() string {
	return bqp.FormatToString(0)
}

// FormatToString pretty-prints the BufferQueueParam using the provided indentation level
func (bqp *BufferQueueParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BufferQueueParam{\n")
	b.WriteString(fmt.Sprintf("%sDataId: %s,\n", indentationValues, bqp.DataId))
	b.WriteString(fmt.Sprintf("%sSlot: %s,\n", indentationValues, bqp.Slot))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() *BufferQueueParam {
	bqp := &BufferQueueParam{
		DataId: types.NewPrimitiveU64(0),
		Slot:   types.NewPrimitiveU32(0),
	}

	return bqp
}
