// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// TextMessage is a data structure used by the Message Delivery protocol
type TextMessage struct {
	types.Structure
	*UserMessage
	StrTextBody string
}

// WriteTo writes the TextMessage to the given writable
func (textMessage *TextMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	textMessage.StrTextBody.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	textMessage.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TextMessage from the given readable
func (textMessage *TextMessage) ExtractFrom(readable types.Readable) error {
	var err error

	if err = textMessage.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read TextMessage header. %s", err.Error())
	}

	err = textMessage.StrTextBody.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TextMessage.StrTextBody from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TextMessage
func (textMessage *TextMessage) Copy() types.RVType {
	copied := NewTextMessage()

	copied.StructureVersion = textMessage.StructureVersion

	copied.UserMessage = textMessage.UserMessage.Copy().(*UserMessage)

	copied.StrTextBody = textMessage.StrTextBody

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (textMessage *TextMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*TextMessage); !ok {
		return false
	}

	other := o.(*TextMessage)

	if textMessage.StructureVersion != other.StructureVersion {
		return false
	}

	if !textMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !textMessage.StrTextBody.Equals(other.StrTextBody) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (textMessage *TextMessage) String() string {
	return textMessage.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (textMessage *TextMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TextMessage{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, textMessage.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, textMessage.StructureVersion))
	b.WriteString(fmt.Sprintf("%sStrTextBody: %q\n", indentationValues, textMessage.StrTextBody))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTextMessage returns a new TextMessage
func NewTextMessage() *TextMessage {
	textMessage := &TextMessage{}
	textMessage.UserMessage = NewUserMessage()
	textMessage.SetParentType(textMessage.UserMessage)

	return textMessage
}
