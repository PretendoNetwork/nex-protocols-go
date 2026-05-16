// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberGetContentParam is a type within the Shop protocol
type SubscriberGetContentParam struct {
	types.Structure
	Topic            types.String
	Size             types.UInt32
	Offset           types.UInt32
	MinimumContentID types.UInt64
}

// WriteTo writes the SubscriberGetContentParam to the given writable
func (sgcp SubscriberGetContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sgcp.Topic.WriteTo(contentWritable)
	sgcp.Size.WriteTo(contentWritable)
	sgcp.Offset.WriteTo(contentWritable)
	sgcp.MinimumContentID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sgcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberGetContentParam from the given readable
func (sgcp *SubscriberGetContentParam) ExtractFrom(readable types.Readable) error {
	if err := sgcp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam header. %s", err.Error())
	}

	if err := sgcp.Topic.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Topic. %s", err.Error())
	}

	if err := sgcp.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Size. %s", err.Error())
	}

	if err := sgcp.Offset.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Offset. %s", err.Error())
	}

	if err := sgcp.MinimumContentID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.MinimumContentID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberGetContentParam
func (sgcp SubscriberGetContentParam) Copy() types.RVType {
	copied := NewSubscriberGetContentParam()

	copied.StructureVersion = sgcp.StructureVersion
	copied.Topic = sgcp.Topic.Copy().(types.String)
	copied.Size = sgcp.Size.Copy().(types.UInt32)
	copied.Offset = sgcp.Offset.Copy().(types.UInt32)
	copied.MinimumContentID = sgcp.MinimumContentID.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given SubscriberGetContentParam contains the same data as the current SubscriberGetContentParam
func (sgcp SubscriberGetContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriberGetContentParam); !ok {
		return false
	}

	other := o.(SubscriberGetContentParam)

	if sgcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sgcp.Topic.Equals(other.Topic) {
		return false
	}

	if !sgcp.Size.Equals(other.Size) {
		return false
	}

	if !sgcp.Offset.Equals(other.Offset) {
		return false
	}

	return sgcp.MinimumContentID.Equals(other.MinimumContentID)
}

// CopyRef copies the current value of the SubscriberGetContentParam
// and returns a pointer to the new copy
func (sgcp SubscriberGetContentParam) CopyRef() types.RVTypePtr {
	copied := sgcp.Copy().(SubscriberGetContentParam)
	return &copied
}

// Deref takes a pointer to the SubscriberGetContentParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sgcp *SubscriberGetContentParam) Deref() types.RVType {
	return *sgcp
}

// String returns the string representation of the SubscriberGetContentParam
func (sgcp SubscriberGetContentParam) String() string {
	return sgcp.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberGetContentParam using the provided indentation level
func (sgcp SubscriberGetContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberGetContentParam{\n")
	b.WriteString(fmt.Sprintf("%sTopic: %s,\n", indentationValues, sgcp.Topic))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, sgcp.Size))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, sgcp.Offset))
	b.WriteString(fmt.Sprintf("%sMinimumContentID: %s,\n", indentationValues, sgcp.MinimumContentID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberGetContentParam returns a new SubscriberGetContentParam
func NewSubscriberGetContentParam() SubscriberGetContentParam {
	return SubscriberGetContentParam{
		Topic:            types.NewString(""),
		Size:             types.NewUInt32(0),
		Offset:           types.NewUInt32(0),
		MinimumContentID: types.NewUInt64(0),
	}

}
