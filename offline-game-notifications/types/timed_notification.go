// Package types implements all the types used by the OfflineGameNotifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	notification_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
)

// TimedNotification is a type with the OfflineGameNotifications protocol
type TimedNotification struct {
	types.Structure
	timestamp    types.DateTime
	notification notification_types.NotificationEvent
}

// WriteTo writes the TimdeNotification to the given writable
func (tn TimedNotification) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	tn.timestamp.WriteTo(contentWritable)
	tn.notification.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tn.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TimedNotificatoin from the given readable
func (tn *TimedNotification) ExtractFrom(readable types.Readable) error {
	if err := tn.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification header. %s", err.Error())
	}

	if err := tn.timestamp.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification.timestamp. %s", err.Error())
	}

	if err := tn.notification.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification.notification. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of TimedNotification
func (tn TimedNotification) Copy() types.RVType {
	copied := NewTimedNotification()

	copied.StructureVersion = tn.StructureVersion
	copied.timestamp = tn.timestamp
	copied.notification = tn.Copy().(notification_types.NotificationEvent)

	return copied
}

// Equals checks if the given TimedNotification contains the same data as the current TimedNotification
func (tn TimedNotification) Equals(o types.RVType) bool {
	if _, ok := o.(TimedNotification); !ok {
		return false
	}

	other := o.(TimedNotification)

	if tn.StructureVersion != other.StructureVersion {
		return false
	}

	if tn.timestamp != other.timestamp {
		return false
	}

	return tn.notification.Equals(other.notification)
}

// CopyRef copies the current value of the TimedNotification
// and returns a pointer to the new copy
func (tn TimedNotification) CopyRef() types.RVTypePtr {
	copied := tn.Copy().(TimedNotification)
	return &copied
}

// Deref takes a pointer of TimedNotification
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (tn *TimedNotification) Deref() types.RVType {
	return *tn
}

// String returns the string representation of the TimedNotification
func (tn TimedNotification) String() string {
	return tn.FormatToString(0)
}

// FormatToString pretty-prints the TimedNotification using the provided indentation level
func (tn TimedNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TimeNotification{\n}")
	fmt.Fprintf(&b, "%stimestamp: %s,\n", indentationValues, tn.timestamp)
	fmt.Fprintf(&b, "%sNotification: %s,\n", indentationValues, tn.notification.FormatToString(indentationLevel+1))

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTimedNotification returns a new TimedNotification
func NewTimedNotification() TimedNotification {
	return TimedNotification{
		timestamp:    types.NewDateTime(0),
		notification: notification_types.NewNotificationEvent(),
	}

}
