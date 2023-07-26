// Package types implements all the types used by the Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// NotificationEvent holds general purpose notification data
type NotificationEvent struct {
	nex.Structure
	PIDSource uint32
	Type      uint32
	Param1    uint32
	Param2    uint32
	StrParam  string
	Param3    uint32
}

// Bytes encodes the NotificationEvent and returns a byte array
func (notificationEvent *NotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	nexVersion := stream.Server.NEXVersion()

	stream.WriteUInt32LE(notificationEvent.PIDSource)
	stream.WriteUInt32LE(notificationEvent.Type)
	stream.WriteUInt32LE(notificationEvent.Param1)
	stream.WriteUInt32LE(notificationEvent.Param2)
	stream.WriteString(notificationEvent.StrParam)

	if nexVersion.Major >= 3 && nexVersion.Minor >= 4 {
		stream.WriteUInt32LE(notificationEvent.Param3)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of NotificationEvent
func (notificationEvent *NotificationEvent) Copy() nex.StructureInterface {
	copied := NewNotificationEvent()

	copied.PIDSource = notificationEvent.PIDSource
	copied.Type = notificationEvent.Type
	copied.Param1 = notificationEvent.Param1
	copied.Param2 = notificationEvent.Param2
	copied.StrParam = notificationEvent.StrParam
	copied.Param3 = notificationEvent.Param3

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notificationEvent *NotificationEvent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NotificationEvent)

	if notificationEvent.PIDSource != other.PIDSource {
		return false
	}

	if notificationEvent.Type != other.Type {
		return false
	}

	if notificationEvent.Param1 != other.Param1 {
		return false
	}

	if notificationEvent.Param2 != other.Param2 {
		return false
	}

	if notificationEvent.StrParam != other.StrParam {
		return false
	}

	if notificationEvent.Param3 != other.Param3 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (notificationEvent *NotificationEvent) String() string {
	return notificationEvent.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (notificationEvent *NotificationEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NotificationEvent{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, notificationEvent.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPIDSource: %d,\n", indentationValues, notificationEvent.PIDSource))
	b.WriteString(fmt.Sprintf("%sType: %d,\n", indentationValues, notificationEvent.Type))
	b.WriteString(fmt.Sprintf("%sParam1: %d,\n", indentationValues, notificationEvent.Param1))
	b.WriteString(fmt.Sprintf("%sParam2: %d,\n", indentationValues, notificationEvent.Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %q,\n", indentationValues, notificationEvent.StrParam))
	b.WriteString(fmt.Sprintf("%sParam3: %d\n", indentationValues, notificationEvent.Param3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	return &NotificationEvent{}
}
