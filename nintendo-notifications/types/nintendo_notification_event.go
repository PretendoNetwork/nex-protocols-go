// Package types implements all the types used by the Nintendo Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// NintendoNotificationEvent is used to send data about a notification event to a client
type NintendoNotificationEvent struct {
	nex.Structure
	Type       uint32
	SenderPID  *nex.PID
	DataHolder *nex.DataHolder
}

// Bytes encodes the NintendoNotificationEvent and returns a byte array
func (nintendoNotificationEvent *NintendoNotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(nintendoNotificationEvent.Type)
	stream.WritePID(nintendoNotificationEvent.SenderPID)
	stream.WriteDataHolder(nintendoNotificationEvent.DataHolder)

	return stream.Bytes()
}

// Copy returns a new copied instance of NintendoNotificationEvent
func (nintendoNotificationEvent *NintendoNotificationEvent) Copy() nex.StructureInterface {
	copied := NewNintendoNotificationEvent()

	copied.SetStructureVersion(nintendoNotificationEvent.StructureVersion())

	copied.Type = nintendoNotificationEvent.Type
	copied.SenderPID = nintendoNotificationEvent.SenderPID.Copy()
	copied.DataHolder = nintendoNotificationEvent.DataHolder.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoNotificationEvent *NintendoNotificationEvent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoNotificationEvent)

	if nintendoNotificationEvent.StructureVersion() != other.StructureVersion() {
		return false
	}

	if nintendoNotificationEvent.Type != other.Type {
		return false
	}

	if !nintendoNotificationEvent.SenderPID.Equals(other.SenderPID) {
		return false
	}

	if !nintendoNotificationEvent.DataHolder.Equals(other.DataHolder) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (nintendoNotificationEvent *NintendoNotificationEvent) String() string {
	return nintendoNotificationEvent.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (nintendoNotificationEvent *NintendoNotificationEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEvent{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, nintendoNotificationEvent.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sType: %d,\n", indentationValues, nintendoNotificationEvent.Type))
	b.WriteString(fmt.Sprintf("%sSenderPID: %s,\n", indentationValues, nintendoNotificationEvent.SenderPID.FormatToString(indentationLevel+1)))

	if nintendoNotificationEvent.DataHolder != nil {
		b.WriteString(fmt.Sprintf("%sDataHolder: %s\n", indentationValues, nintendoNotificationEvent.DataHolder.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDataHolder: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEvent returns a new NintendoNotificationEvent
func NewNintendoNotificationEvent() *NintendoNotificationEvent {
	return &NintendoNotificationEvent{}
}
