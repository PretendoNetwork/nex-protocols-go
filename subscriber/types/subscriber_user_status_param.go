// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberUserStatusParam is a type within the Shop protocol
type SubscriberUserStatusParam struct {
	types.Structure
	Value types.QBuffer
}

// WriteTo writes the SubscriberUserStatusParam to the given writable
func (u SubscriberUserStatusParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	u.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	u.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberUserStatusParam from the given readable
func (u *SubscriberUserStatusParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = u.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusParam header. %s", err.Error())
	}

	err = u.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusParam.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberUserStatusParam
func (u SubscriberUserStatusParam) Copy() types.RVType {
	copied := NewSubscriberUserStatusParam()

	copied.StructureVersion = u.StructureVersion
	copied.Value = u.Value.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given SubscriberUserStatusParam contains the same data as the current SubscriberUserStatusParam
func (u SubscriberUserStatusParam) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriberUserStatusParam); !ok {
		return false
	}

	other := o.(SubscriberUserStatusParam)

	if u.StructureVersion != other.StructureVersion {
		return false
	}

	return u.Value.Equals(other.Value)
}

// CopyRef copies the current value of the SubscriberUserStatusParam
// and returns a pointer to the new copy
func (u SubscriberUserStatusParam) CopyRef() types.RVTypePtr {
	copied := u.Copy().(SubscriberUserStatusParam)
	return &copied
}

// Deref takes a pointer to the SubscriberUserStatusParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (u *SubscriberUserStatusParam) Deref() types.RVType {
	return *u
}

// String returns the string representation of the SubscriberUserStatusParam
func (u SubscriberUserStatusParam) String() string {
	return u.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberUserStatusParam using the provided indentation level
func (u SubscriberUserStatusParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberUserStatusParam{\n")
	b.WriteString(fmt.Sprintf("%sValue: %s,\n", indentationValues, u.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberUserStatusParam returns a new SubscriberUserStatusParam
func NewSubscriberUserStatusParam() SubscriberUserStatusParam {
	return SubscriberUserStatusParam{
		Value: types.NewQBuffer(nil),
	}

}
