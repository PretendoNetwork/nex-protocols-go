// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SubscriberContent is unknown
type SubscriberContent struct {
	types.Structure
	Unknown1 *types.PrimitiveU64
	Unknown2 string
	Unknown3 []byte
	Unknown4 *types.PrimitiveU64
	Unknown5 *types.List[*types.String]
	Unknown6 *types.DateTime
}

// ExtractFrom extracts the SubscriberContent from the given readable
func (subscriberContent *SubscriberContent) ExtractFrom(readable types.Readable) error {
	var err error

	if err = subscriberContent.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SubscriberContent header. %s", err.Error())
	}

	err = subscriberContent.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown1 from stream. %s", err.Error())
	}

	err = subscriberContent.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown2 from stream. %s", err.Error())
	}

	subscriberContent.Unknown3, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown3 from stream. %s", err.Error())
	}

	err = subscriberContent.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown4 from stream. %s", err.Error())
	}

	err = subscriberContent.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown5 from stream. %s", err.Error())
	}

	err = subscriberContent.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown6 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SubscriberContent to the given writable
func (subscriberContent *SubscriberContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	subscriberContent.Unknown1.WriteTo(contentWritable)
	subscriberContent.Unknown2.WriteTo(contentWritable)
	stream.WriteQBuffer(subscriberContent.Unknown3)
	subscriberContent.Unknown4.WriteTo(contentWritable)
	subscriberContent.Unknown5.WriteTo(contentWritable)
	subscriberContent.Unknown6.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	subscriberContent.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SubscriberContent
func (subscriberContent *SubscriberContent) Copy() types.RVType {
	copied := NewSubscriberContent()

	copied.StructureVersion = subscriberContent.StructureVersion

	copied.Unknown1 = subscriberContent.Unknown1
	copied.Unknown2 = subscriberContent.Unknown2
	copied.Unknown3 = make([]byte, len(subscriberContent.Unknown3))

	copy(copied.Unknown3, subscriberContent.Unknown3)

	copied.Unknown4 = subscriberContent.Unknown4
	copied.Unknown5 = make(*types.List[*types.String], len(subscriberContent.Unknown5))

	copy(copied.Unknown5, subscriberContent.Unknown5)

	copied.Unknown6 = subscriberContent.Unknown6.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberContent *SubscriberContent) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberContent); !ok {
		return false
	}

	other := o.(*SubscriberContent)

	if subscriberContent.StructureVersion != other.StructureVersion {
		return false
	}

	if !subscriberContent.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !subscriberContent.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !subscriberContent.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !subscriberContent.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if len(subscriberContent.Unknown5) != len(other.Unknown5) {
		return false
	}

	for i := 0; i < len(subscriberContent.Unknown5); i++ {
		if subscriberContent.Unknown5[i] != other.Unknown5[i] {
			return false
		}
	}

	return subscriberContent.Unknown6.Equals(other.Unknown6)
}

// String returns a string representation of the struct
func (subscriberContent *SubscriberContent) String() string {
	return subscriberContent.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (subscriberContent *SubscriberContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberContent{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, subscriberContent.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, subscriberContent.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %q,\n", indentationValues, subscriberContent.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %x,\n", indentationValues, subscriberContent.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d,\n", indentationValues, subscriberContent.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %v,\n", indentationValues, subscriberContent.Unknown5))

	if subscriberContent.Unknown6 != nil {
		b.WriteString(fmt.Sprintf("%sUnknown6: %s\n", indentationValues, subscriberContent.Unknown6.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown6: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberContent returns a new SubscriberContent
func NewSubscriberContent() *SubscriberContent {
	return &SubscriberContent{}
}
