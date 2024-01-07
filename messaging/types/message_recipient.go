// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MessageRecipient is a data structure used by the Message Delivery protocol
type MessageRecipient struct {
	types.Structure
	UIRecipientType *types.PrimitiveU32
	PrincipalID     *types.PID
	GatheringID     *types.PrimitiveU32
}

// ExtractFrom extracts the MessageRecipient from the given readable
func (messageRecipient *MessageRecipient) ExtractFrom(readable types.Readable) error {
	var err error

	if err = messageRecipient.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MessageRecipient header. %s", err.Error())
	}

	err = messageRecipient.UIRecipientType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.UIRecipientType from stream. %s", err.Error())
	}

	err = messageRecipient.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.PrincipalID from stream. %s", err.Error())
	}

	err = messageRecipient.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.GatheringID from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (messageRecipient *MessageRecipient) Copy() types.RVType {
	copied := NewMessageRecipient()

	copied.StructureVersion = messageRecipient.StructureVersion

	copied.UIRecipientType = messageRecipient.UIRecipientType
	copied.PrincipalID = messageRecipient.PrincipalID.Copy()
	copied.GatheringID = messageRecipient.GatheringID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (messageRecipient *MessageRecipient) Equals(o types.RVType) bool {
	if _, ok := o.(*MessageRecipient); !ok {
		return false
	}

	other := o.(*MessageRecipient)

	if messageRecipient.StructureVersion != other.StructureVersion {
		return false
	}

	if !messageRecipient.UIRecipientType.Equals(other.UIRecipientType) {
		return false
	}

	if !messageRecipient.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !messageRecipient.GatheringID.Equals(other.GatheringID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (messageRecipient *MessageRecipient) String() string {
	return messageRecipient.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (messageRecipient *MessageRecipient) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MessageRecipient{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, messageRecipient.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUIRecipientType: %d,\n", indentationValues, messageRecipient.UIRecipientType))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, messageRecipient.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %d\n", indentationValues, messageRecipient.GatheringID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	return &MessageRecipient{}
}
