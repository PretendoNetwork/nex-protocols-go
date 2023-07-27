// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// TextMessage is a data structure used by the Message Delivery protocol
type TextMessage struct {
	nex.Structure
	*UserMessage
	StrTextBody string
}

// Bytes encodes the TextMessage and returns a byte array
func (textMessage *TextMessage) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(textMessage.StrTextBody)

	return stream.Bytes()
}

// ExtractFromStream extracts a TextMessage structure from a stream
func (textMessage *TextMessage) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	textMessage.StrTextBody, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract TextMessage.StrTextBody from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TextMessage
func (textMessage *TextMessage) Copy() nex.StructureInterface {
	copied := NewTextMessage()

	copied.UserMessage = textMessage.UserMessage.Copy().(*UserMessage)
	copied.SetParentType(copied.UserMessage)

	copied.StrTextBody = textMessage.StrTextBody

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (textMessage *TextMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*TextMessage)

	if !textMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if textMessage.StrTextBody != other.StrTextBody {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, textMessage.StructureVersion()))
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
