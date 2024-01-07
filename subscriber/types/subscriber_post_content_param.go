// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SubscriberPostContentParam is unknown
type SubscriberPostContentParam struct {
	types.Structure
	Unknown1 *types.List[*types.String]
	Unknown2 string
	Unknown3 []byte
}

// ExtractFrom extracts the SubscriberPostContentParam from the given readable
func (subscriberPostContentParam *SubscriberPostContentParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = subscriberPostContentParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SubscriberPostContentParam header. %s", err.Error())
	}

	err = subscriberPostContentParam.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown1 from stream. %s", err.Error())
	}

	err = subscriberPostContentParam.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown2 from stream. %s", err.Error())
	}

	subscriberPostContentParam.Unknown3, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberPostContentParam.Unknown3 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SubscriberPostContentParam to the given writable
func (subscriberPostContentParam *SubscriberPostContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	subscriberPostContentParam.Unknown1.WriteTo(contentWritable)
	subscriberPostContentParam.Unknown2.WriteTo(contentWritable)
	stream.WriteQBuffer(subscriberPostContentParam.Unknown3)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SubscriberPostContentParam
func (subscriberPostContentParam *SubscriberPostContentParam) Copy() types.RVType {
	copied := NewSubscriberPostContentParam()

	copied.StructureVersion = subscriberPostContentParam.StructureVersion

	copied.Unknown1 = make(*types.List[*types.String], len(subscriberPostContentParam.Unknown1))

	copy(copied.Unknown1, subscriberPostContentParam.Unknown1)

	copied.Unknown2 = subscriberPostContentParam.Unknown2
	copied.Unknown3 = make([]byte, len(subscriberPostContentParam.Unknown3))

	copy(copied.Unknown3, subscriberPostContentParam.Unknown3)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberPostContentParam *SubscriberPostContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberPostContentParam); !ok {
		return false
	}

	other := o.(*SubscriberPostContentParam)

	if subscriberPostContentParam.StructureVersion != other.StructureVersion {
		return false
	}

	if len(subscriberPostContentParam.Unknown1) != len(other.Unknown1) {
		return false
	}

	for i := 0; i < len(subscriberPostContentParam.Unknown1); i++ {
		if subscriberPostContentParam.Unknown1[i] != other.Unknown1[i] {
			return false
		}
	}

	if !subscriberPostContentParam.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !subscriberPostContentParam.Unknown3.Equals(other.Unknown3) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, subscriberPostContentParam.StructureVersion))
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
