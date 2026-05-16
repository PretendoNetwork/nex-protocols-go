// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberPostContentParam is a type within the Shop protocol
type SubscriberPostContentParam struct {
	types.Structure
	Topic   types.List[types.String]
	Message types.String
	Binary  types.QBuffer
}

// WriteTo writes the SubscriberPostContentParam to the given writable
func (spcp SubscriberPostContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	spcp.Topic.WriteTo(contentWritable)
	spcp.Message.WriteTo(contentWritable)
	spcp.Binary.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	spcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberPostContentParam from the given readable
func (spcp *SubscriberPostContentParam) ExtractFrom(readable types.Readable) error {
	if err := spcp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam header. %s", err.Error())
	}

	if err := spcp.Topic.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Topic. %s", err.Error())
	}

	if err := spcp.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Message. %s", err.Error())
	}

	if err := spcp.Binary.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Binary. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberPostContentParam
func (spcp SubscriberPostContentParam) Copy() types.RVType {
	copied := NewSubscriberPostContentParam()

	copied.StructureVersion = spcp.StructureVersion
	copied.Topic = spcp.Topic.Copy().(types.List[types.String])
	copied.Message = spcp.Message.Copy().(types.String)
	copied.Binary = spcp.Binary.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given SubscriberPostContentParam contains the same data as the current SubscriberPostContentParam
func (spcp SubscriberPostContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriberPostContentParam); !ok {
		return false
	}

	other := o.(SubscriberPostContentParam)

	if spcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !spcp.Topic.Equals(other.Topic) {
		return false
	}

	if !spcp.Message.Equals(other.Message) {
		return false
	}

	return spcp.Binary.Equals(other.Binary)
}

// CopyRef copies the current value of the SubscriberPostContentParam
// and returns a pointer to the new copy
func (spcp SubscriberPostContentParam) CopyRef() types.RVTypePtr {
	copied := spcp.Copy().(SubscriberPostContentParam)
	return &copied
}

// Deref takes a pointer to the SubscriberPostContentParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (spcp *SubscriberPostContentParam) Deref() types.RVType {
	return *spcp
}

// String returns the string representation of the SubscriberPostContentParam
func (spcp SubscriberPostContentParam) String() string {
	return spcp.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberPostContentParam using the provided indentation level
func (spcp SubscriberPostContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberPostContentParam{\n")
	b.WriteString(fmt.Sprintf("%sTopic: %s,\n", indentationValues, spcp.Topic))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, spcp.Message))
	b.WriteString(fmt.Sprintf("%sBinary: %s,\n", indentationValues, spcp.Binary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberPostContentParam returns a new SubscriberPostContentParam
func NewSubscriberPostContentParam() SubscriberPostContentParam {
	return SubscriberPostContentParam{
		Topic:   types.NewList[types.String](),
		Message: types.NewString(""),
		Binary:  types.NewQBuffer(nil),
	}

}
