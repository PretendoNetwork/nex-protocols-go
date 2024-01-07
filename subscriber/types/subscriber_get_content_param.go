// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SubscriberGetContentParam is unknown
type SubscriberGetContentParam struct {
	types.Structure
	Unknown1 string
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU64
}

// ExtractFrom extracts the SubscriberGetContentParam from the given readable
func (subscriberGetContentParam *SubscriberGetContentParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = subscriberGetContentParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SubscriberGetContentParam header. %s", err.Error())
	}

	err = subscriberGetContentParam.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown1 from stream. %s", err.Error())
	}

	err = subscriberGetContentParam.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown2 from stream. %s", err.Error())
	}

	err = subscriberGetContentParam.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown3 from stream. %s", err.Error())
	}

	err = subscriberGetContentParam.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown4 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SubscriberGetContentParam to the given writable
func (subscriberGetContentParam *SubscriberGetContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	subscriberGetContentParam.Unknown1.WriteTo(contentWritable)
	subscriberGetContentParam.Unknown2.WriteTo(contentWritable)
	subscriberGetContentParam.Unknown3.WriteTo(contentWritable)
	subscriberGetContentParam.Unknown4.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	subscriberGetContentParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SubscriberGetContentParam
func (subscriberGetContentParam *SubscriberGetContentParam) Copy() types.RVType {
	copied := NewSubscriberGetContentParam()

	copied.StructureVersion = subscriberGetContentParam.StructureVersion

	copied.Unknown1 = subscriberGetContentParam.Unknown1
	copied.Unknown2 = subscriberGetContentParam.Unknown2
	copied.Unknown3 = subscriberGetContentParam.Unknown3
	copied.Unknown4 = subscriberGetContentParam.Unknown4

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberGetContentParam *SubscriberGetContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberGetContentParam); !ok {
		return false
	}

	other := o.(*SubscriberGetContentParam)

	if subscriberGetContentParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !subscriberGetContentParam.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !subscriberGetContentParam.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !subscriberGetContentParam.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !subscriberGetContentParam.Unknown4.Equals(other.Unknown4) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (subscriberGetContentParam *SubscriberGetContentParam) String() string {
	return subscriberGetContentParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (subscriberGetContentParam *SubscriberGetContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberGetContentParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, subscriberGetContentParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %q,\n", indentationValues, subscriberGetContentParam.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, subscriberGetContentParam.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, subscriberGetContentParam.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d\n", indentationValues, subscriberGetContentParam.Unknown4))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberGetContentParam returns a new SubscriberGetContentParam
func NewSubscriberGetContentParam() *SubscriberGetContentParam {
	return &SubscriberGetContentParam{}
}
