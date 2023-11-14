// Package types implements all the types used by the Message Delivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MessageRecipient is a data structure used by the Message Delivery protocol
type MessageRecipient struct {
	nex.Structure
	UIRecipientType uint32
	PrincipalID     *nex.PID
	GatheringID     uint32
}

// ExtractFromStream extracts a MessageRecipient structure from a stream
func (messageRecipient *MessageRecipient) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	messageRecipient.UIRecipientType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.UIRecipientType from stream. %s", err.Error())
	}

	messageRecipient.PrincipalID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.PrincipalID from stream. %s", err.Error())
	}

	messageRecipient.GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.GatheringID from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (messageRecipient *MessageRecipient) Copy() nex.StructureInterface {
	copied := NewMessageRecipient()

	copied.SetStructureVersion(messageRecipient.StructureVersion())

	copied.UIRecipientType = messageRecipient.UIRecipientType
	copied.PrincipalID = messageRecipient.PrincipalID.Copy()
	copied.GatheringID = messageRecipient.GatheringID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (messageRecipient *MessageRecipient) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MessageRecipient)

	if messageRecipient.StructureVersion() != other.StructureVersion() {
		return false
	}

	if messageRecipient.UIRecipientType != other.UIRecipientType {
		return false
	}

	if !messageRecipient.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if messageRecipient.GatheringID != other.GatheringID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, messageRecipient.StructureVersion()))
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
