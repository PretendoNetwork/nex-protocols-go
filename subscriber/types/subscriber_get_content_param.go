// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SubscriberGetContentParam is unknown
type SubscriberGetContentParam struct {
	nex.Structure
	Unknown1 string
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint64
}

// ExtractFromStream extracts a SubscriberGetContentParam structure from a stream
func (subscriberGetContentParam *SubscriberGetContentParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	subscriberGetContentParam.Unknown1, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown1 from stream. %s", err.Error())
	}

	subscriberGetContentParam.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown2 from stream. %s", err.Error())
	}

	subscriberGetContentParam.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown3 from stream. %s", err.Error())
	}

	subscriberGetContentParam.Unknown4, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberGetContentParam.Unknown4 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SubscriberGetContentParam and returns a byte array
func (subscriberGetContentParam *SubscriberGetContentParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(subscriberGetContentParam.Unknown1)
	stream.WriteUInt32LE(subscriberGetContentParam.Unknown2)
	stream.WriteUInt32LE(subscriberGetContentParam.Unknown3)
	stream.WriteUInt64LE(subscriberGetContentParam.Unknown4)

	return stream.Bytes()
}

// Copy returns a new copied instance of SubscriberGetContentParam
func (subscriberGetContentParam *SubscriberGetContentParam) Copy() nex.StructureInterface {
	copied := NewSubscriberGetContentParam()

	copied.Unknown1 = subscriberGetContentParam.Unknown1
	copied.Unknown2 = subscriberGetContentParam.Unknown2
	copied.Unknown3 = subscriberGetContentParam.Unknown3
	copied.Unknown4 = subscriberGetContentParam.Unknown4

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberGetContentParam *SubscriberGetContentParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SubscriberGetContentParam)

	if subscriberGetContentParam.Unknown1 != other.Unknown1 {
		return false
	}

	if subscriberGetContentParam.Unknown2 != other.Unknown2 {
		return false
	}

	if subscriberGetContentParam.Unknown3 != other.Unknown3 {
		return false
	}

	if subscriberGetContentParam.Unknown4 != other.Unknown4 {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, subscriberGetContentParam.StructureVersion()))
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
