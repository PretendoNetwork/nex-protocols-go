// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SubscriberPostContentParam is unknown
type SubscriberPostContentParam struct {
	nex.Structure
	Unknown1 []string
	Unknown2 string
	Unknown3 []byte
}

// ExtractFromStream extracts a SubscriberPostContentParam structure from a stream
func (subscriberPostContentParam *SubscriberPostContentParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	subscriberPostContentParam.Unknown1, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown1 from stream. %s", err.Error())
	}

	subscriberPostContentParam.Unknown2, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown2 from stream. %s", err.Error())
	}

	subscriberPostContentParam.Unknown3, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown3 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SubscriberPostContentParam and returns a byte array
func (subscriberPostContentParam *SubscriberPostContentParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListString(subscriberPostContentParam.Unknown1)
	stream.WriteString(subscriberPostContentParam.Unknown2)
	stream.WriteQBuffer(subscriberPostContentParam.Unknown3)

	return stream.Bytes()
}

// Copy returns a new copied instance of SubscriberPostContentParam
func (subscriberPostContentParam *SubscriberPostContentParam) Copy() nex.StructureInterface {
	copied := NewSubscriberPostContentParam()

	copied.Unknown1 = make([]string, len(subscriberPostContentParam.Unknown1))

	copy(copied.Unknown1, subscriberPostContentParam.Unknown1)

	copied.Unknown2 = subscriberPostContentParam.Unknown2
	copied.Unknown3 = make([]byte, len(subscriberPostContentParam.Unknown3))

	copy(copied.Unknown3, subscriberPostContentParam.Unknown3)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberPostContentParam *SubscriberPostContentParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SubscriberPostContentParam)

	if len(subscriberPostContentParam.Unknown1) != len(other.Unknown1) {
		return false
	}

	for i := 0; i < len(subscriberPostContentParam.Unknown1); i++ {
		if subscriberPostContentParam.Unknown1[i] != other.Unknown1[i] {
			return false
		}
	}

	if subscriberPostContentParam.Unknown2 != other.Unknown2 {
		return false
	}

	if !bytes.Equal(subscriberPostContentParam.Unknown3, other.Unknown3) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (subscriberPostContentParam *SubscriberPostContentParam) String() string {
	return subscriberPostContentParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (subscriberPostContentParam *SubscriberPostContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberPostContentParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, subscriberPostContentParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUnknown1: %v,\n", indentationValues, subscriberPostContentParam.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %q,\n", indentationValues, subscriberPostContentParam.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %x\n", indentationValues, subscriberPostContentParam.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberPostContentParam returns a new SubscriberPostContentParam
func NewSubscriberPostContentParam() *SubscriberPostContentParam {
	return &SubscriberPostContentParam{}
}
