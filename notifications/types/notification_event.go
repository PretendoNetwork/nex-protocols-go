// Package types implements all the types used by the Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NotificationEvent holds general purpose notification data
type NotificationEvent struct {
	types.Structure
	PIDSource *types.PID
	Type      *types.PrimitiveU32
	Param1    *types.PrimitiveU32
	Param2    *types.PrimitiveU32
	StrParam  string
	Param3    *types.PrimitiveU32
}

// WriteTo writes the NotificationEvent to the given writable
func (notificationEvent *NotificationEvent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nexVersion := stream.Server.LibraryVersion()

	notificationEvent.PIDSource.WriteTo(contentWritable)
	notificationEvent.Type.WriteTo(contentWritable)
	notificationEvent.Param1.WriteTo(contentWritable)
	notificationEvent.Param2.WriteTo(contentWritable)
	notificationEvent.StrParam.WriteTo(contentWritable)

	if nexVersion.GreaterOrEqual("3.4.0") {
		notificationEvent.Param3.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of NotificationEvent
func (notificationEvent *NotificationEvent) Copy() types.RVType {
	copied := NewNotificationEvent()

	copied.StructureVersion = notificationEvent.StructureVersion

	copied.PIDSource = notificationEvent.PIDSource.Copy()
	copied.Type = notificationEvent.Type
	copied.Param1 = notificationEvent.Param1
	copied.Param2 = notificationEvent.Param2
	copied.StrParam = notificationEvent.StrParam
	copied.Param3 = notificationEvent.Param3

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notificationEvent *NotificationEvent) Equals(o types.RVType) bool {
	if _, ok := o.(*NotificationEvent); !ok {
		return false
	}

	other := o.(*NotificationEvent)

	if notificationEvent.StructureVersion != other.StructureVersion {
		return false
	}

	if !notificationEvent.PIDSource.Equals(other.PIDSource) {
		return false
	}

	if !notificationEvent.Type.Equals(other.Type) {
		return false
	}

	if !notificationEvent.Param1.Equals(other.Param1) {
		return false
	}

	if !notificationEvent.Param2.Equals(other.Param2) {
		return false
	}

	if !notificationEvent.StrParam.Equals(other.StrParam) {
		return false
	}

	if !notificationEvent.Param3.Equals(other.Param3) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, notificationEvent.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPIDSource: %s,\n", indentationValues, notificationEvent.PIDSource.FormatToString(indentationLevel+1)))
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
