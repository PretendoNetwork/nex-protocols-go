// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SubscriberContent is unknown
type SubscriberContent struct {
	nex.Structure
	Unknown1 uint64
	Unknown2 string
	Unknown3 []byte
	Unknown4 uint64
	Unknown5 []string
	Unknown6 *nex.DateTime
}

// ExtractFromStream extracts a SubscriberContent structure from a stream
func (subscriberContent *SubscriberContent) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	subscriberContent.Unknown1, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown1 from stream. %s", err.Error())
	}

	subscriberContent.Unknown2, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown2 from stream. %s", err.Error())
	}

	subscriberContent.Unknown3, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown3 from stream. %s", err.Error())
	}

	subscriberContent.Unknown4, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown4 from stream. %s", err.Error())
	}

	subscriberContent.Unknown5, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown5 from stream. %s", err.Error())
	}

	subscriberContent.Unknown6, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Unknown6 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SubscriberContent and returns a byte array
func (subscriberContent *SubscriberContent) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(subscriberContent.Unknown1)
	stream.WriteString(subscriberContent.Unknown2)
	stream.WriteQBuffer(subscriberContent.Unknown3)
	stream.WriteUInt64LE(subscriberContent.Unknown4)
	stream.WriteListString(subscriberContent.Unknown5)
	stream.WriteDateTime(subscriberContent.Unknown6)

	return stream.Bytes()
}

// Copy returns a new copied instance of SubscriberContent
func (subscriberContent *SubscriberContent) Copy() nex.StructureInterface {
	copied := NewSubscriberContent()

	copied.Unknown1 = subscriberContent.Unknown1
	copied.Unknown2 = subscriberContent.Unknown2
	copied.Unknown3 = make([]byte, len(subscriberContent.Unknown3))

	copy(copied.Unknown3, subscriberContent.Unknown3)

	copied.Unknown4 = subscriberContent.Unknown4
	copied.Unknown5 = make([]string, len(subscriberContent.Unknown5))

	copy(copied.Unknown5, subscriberContent.Unknown5)

	if subscriberContent.Unknown6 != nil {
		copied.Unknown6 = subscriberContent.Unknown6.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberContent *SubscriberContent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SubscriberContent)

	if subscriberContent.Unknown1 != other.Unknown1 {
		return false
	}

	if subscriberContent.Unknown2 != other.Unknown2 {
		return false
	}

	if !bytes.Equal(subscriberContent.Unknown3, other.Unknown3) {
		return false
	}

	if subscriberContent.Unknown4 != other.Unknown4 {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, subscriberContent.StructureVersion()))
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
