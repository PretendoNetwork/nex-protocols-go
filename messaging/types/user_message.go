// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// UserMessage is a data structure used by the Message Delivery protocol
type UserMessage struct {
	types.Structure
	*types.Data
	UIID             *types.PrimitiveU32
	UIParentID       *types.PrimitiveU32
	PIDSender        *types.PID
	Receptiontime    *types.DateTime
	UILifeTime       *types.PrimitiveU32
	UIFlags          *types.PrimitiveU32
	StrSubject       string
	StrSender        string
	MessageRecipient *MessageRecipient
}

// ExtractFrom extracts the UserMessage from the given readable
func (userMessage *UserMessage) ExtractFrom(readable types.Readable) error {
	var err error

	if err = userMessage.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read UserMessage header. %s", err.Error())
	}

	err = userMessage.UIID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIID from stream. %s", err.Error())
	}

	err = userMessage.UIParentID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIParentID from stream. %s", err.Error())
	}

	err = userMessage.PIDSender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.PIDSender from stream. %s", err.Error())
	}

	err = userMessage.Receptiontime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.Receptiontime from stream. %s", err.Error())
	}

	err = userMessage.UILifeTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UILifeTime from stream. %s", err.Error())
	}

	err = userMessage.UIFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIFlags from stream. %s", err.Error())
	}

	err = userMessage.StrSubject.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSubject from stream. %s", err.Error())
	}

	err = userMessage.StrSender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSender from stream. %s", err.Error())
	}

	err = userMessage.MessageRecipient.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.MessageRecipient from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserMessage
func (userMessage *UserMessage) Copy() types.RVType {
	copied := NewUserMessage()

	copied.StructureVersion = userMessage.StructureVersion

	copied.Data = userMessage.Data.Copy().(*types.Data)

	copied.UIID = userMessage.UIID
	copied.UIParentID = userMessage.UIParentID
	copied.PIDSender = userMessage.PIDSender.Copy()
	copied.Receptiontime = userMessage.Receptiontime.Copy()
	copied.UILifeTime = userMessage.UILifeTime
	copied.UIFlags = userMessage.UIFlags
	copied.StrSubject = userMessage.StrSubject
	copied.StrSender = userMessage.StrSender
	copied.MessageRecipient = userMessage.MessageRecipient.Copy().(*MessageRecipient)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (userMessage *UserMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*UserMessage); !ok {
		return false
	}

	other := o.(*UserMessage)

	if userMessage.StructureVersion != other.StructureVersion {
		return false
	}

	if !userMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !userMessage.UIID.Equals(other.UIID) {
		return false
	}

	if !userMessage.UIParentID.Equals(other.UIParentID) {
		return false
	}

	if !userMessage.PIDSender.Equals(other.PIDSender) {
		return false
	}

	if !userMessage.Receptiontime.Equals(other.Receptiontime) {
		return false
	}

	if !userMessage.UILifeTime.Equals(other.UILifeTime) {
		return false
	}

	if !userMessage.UIFlags.Equals(other.UIFlags) {
		return false
	}

	if !userMessage.StrSubject.Equals(other.StrSubject) {
		return false
	}

	if !userMessage.StrSender.Equals(other.StrSender) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, userMessage.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUIID: %d,\n", indentationValues, userMessage.UIID))
	b.WriteString(fmt.Sprintf("%sUIParentID: %d,\n", indentationValues, userMessage.UIParentID))
	b.WriteString(fmt.Sprintf("%sPIDSender: %s,\n", indentationValues, userMessage.PIDSender.FormatToString(indentationLevel+1)))

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
	userMessage.Data = types.NewData()
	userMessage.SetParentType(userMessage.Data)

	return userMessage
}
