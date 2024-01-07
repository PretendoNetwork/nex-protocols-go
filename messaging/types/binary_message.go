// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// BinaryMessage is a data structure used by the Message Delivery protocol
type BinaryMessage struct {
	types.Structure
	*UserMessage
	BinaryBody []byte
}

// WriteTo writes the BinaryMessage to the given writable
func (binaryMessage *BinaryMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	stream.WriteQBuffer(binaryMessage.BinaryBody)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BinaryMessage from the given readable
func (binaryMessage *BinaryMessage) ExtractFrom(readable types.Readable) error {
	var err error

	if err = binaryMessage.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BinaryMessage header. %s", err.Error())
	}

	binaryMessage.BinaryBody, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract BinaryMessage.BinaryBody from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BinaryMessage
func (binaryMessage *BinaryMessage) Copy() types.RVType {
	copied := NewBinaryMessage()

	copied.StructureVersion = binaryMessage.StructureVersion

	copied.UserMessage = binaryMessage.UserMessage.Copy().(*UserMessage)

	copied.BinaryBody = make([]byte, len(binaryMessage.BinaryBody))

	copy(copied.BinaryBody, binaryMessage.BinaryBody)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (binaryMessage *BinaryMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*BinaryMessage); !ok {
		return false
	}

	other := o.(*BinaryMessage)

	if binaryMessage.StructureVersion != other.StructureVersion {
		return false
	}

	if !binaryMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !binaryMessage.BinaryBody.Equals(other.BinaryBody) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (binaryMessage *BinaryMessage) String() string {
	return binaryMessage.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (binaryMessage *BinaryMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BinaryMessage{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, binaryMessage.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, binaryMessage.StructureVersion))
	b.WriteString(fmt.Sprintf("%sBinaryBody: %x\n", indentationValues, binaryMessage.BinaryBody))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBinaryMessage returns a new BinaryMessage
func NewBinaryMessage() *BinaryMessage {
	binaryMessage := &BinaryMessage{}
	binaryMessage.UserMessage = NewUserMessage()
	binaryMessage.SetParentType(binaryMessage.UserMessage)

	return binaryMessage
}
