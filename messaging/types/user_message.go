// Package messaging_types implements all the types used by the Message Delivery protocol
package messaging_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// UserMessage is a data structure used by the Message Delivery protocol
type UserMessage struct {
	nex.Structure
	*nex.Data
	UIID             uint32
	UIParentID       uint32
	PIDSender        uint32
	Receptiontime    *nex.DateTime
	UILifeTime       uint32
	UIFlags          uint32
	StrSubject       string
	StrSender        string
	MessageRecipient *MessageRecipient
}

// ExtractFromStream extracts a UserMessage structure from a stream
func (userMessage *UserMessage) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	userMessage.UIID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIID from stream. %s", err.Error())
	}

	userMessage.UIParentID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIParentID from stream. %s", err.Error())
	}

	userMessage.PIDSender, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.PIDSender from stream. %s", err.Error())
	}

	userMessage.Receptiontime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.Receptiontime from stream. %s", err.Error())
	}

	userMessage.UILifeTime, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UILifeTime from stream. %s", err.Error())
	}

	userMessage.UIFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIFlags from stream. %s", err.Error())
	}

	userMessage.StrSubject, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSubject from stream. %s", err.Error())
	}

	userMessage.StrSender, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSender from stream. %s", err.Error())
	}

	messageRecipient, err := stream.ReadStructure(NewMessageRecipient())
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.MessageRecipient from stream. %s", err.Error())
	}

	userMessage.MessageRecipient = messageRecipient.(*MessageRecipient)

	return nil
}

// Copy returns a new copied instance of UserMessage
func (userMessage *UserMessage) Copy() nex.StructureInterface {
	copied := NewUserMessage()

	copied.Data = userMessage.Data.Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.UIID = userMessage.UIID
	copied.UIParentID = userMessage.UIParentID
	copied.PIDSender = userMessage.PIDSender
	copied.Receptiontime = userMessage.Receptiontime.Copy()
	copied.UILifeTime = userMessage.UILifeTime
	copied.UIFlags = userMessage.UIFlags
	copied.StrSubject = userMessage.StrSubject
	copied.StrSender = userMessage.StrSender
	copied.MessageRecipient = userMessage.MessageRecipient.Copy().(*MessageRecipient)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (userMessage *UserMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*UserMessage)

	if !userMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if userMessage.UIID != other.UIID {
		return false
	}

	if userMessage.UIParentID != other.UIParentID {
		return false
	}

	if userMessage.PIDSender != other.PIDSender {
		return false
	}

	if !userMessage.Receptiontime.Equals(other.Receptiontime) {
		return false
	}

	if userMessage.UILifeTime != other.UILifeTime {
		return false
	}

	if userMessage.UIFlags != other.UIFlags {
		return false
	}

	if userMessage.StrSubject != other.StrSubject {
		return false
	}

	if userMessage.StrSender != other.StrSender {
		return false
	}

	if !userMessage.MessageRecipient.Equals(other.MessageRecipient) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (userMessage *UserMessage) String() string {
	return userMessage.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (userMessage *UserMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserMessage{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, userMessage.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, userMessage.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUIID: %d,\n", indentationValues, userMessage.UIID))
	b.WriteString(fmt.Sprintf("%sUIParentID: %d,\n", indentationValues, userMessage.UIParentID))
	b.WriteString(fmt.Sprintf("%sPIDSender: %d,\n", indentationValues, userMessage.PIDSender))

	if userMessage.Receptiontime != nil {
		b.WriteString(fmt.Sprintf("%sReceptiontime: %s,\n", indentationValues, userMessage.Receptiontime))
	} else {
		b.WriteString(fmt.Sprintf("%sReceptiontime: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUILifeTime: %d,\n", indentationValues, userMessage.UILifeTime))
	b.WriteString(fmt.Sprintf("%sUIFlags: %d,\n", indentationValues, userMessage.UIFlags))
	b.WriteString(fmt.Sprintf("%sStrSubject: %q,\n", indentationValues, userMessage.StrSubject))
	b.WriteString(fmt.Sprintf("%sStrSender: %q,\n", indentationValues, userMessage.StrSender))

	if userMessage.MessageRecipient != nil {
		b.WriteString(fmt.Sprintf("%sMessageRecipient: %s,\n", indentationValues, userMessage.MessageRecipient))
	} else {
		b.WriteString(fmt.Sprintf("%sMessageRecipient: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUserMessage returns a new UserMessage
func NewUserMessage() *UserMessage {
	userMessage := &UserMessage{}
	userMessage.Data = nex.NewData()
	userMessage.SetParentType(userMessage.Data)

	return userMessage
}
