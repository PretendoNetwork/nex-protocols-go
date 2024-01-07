// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PersistentNotificationList contains unknown data
type PersistentNotificationList struct {
	types.Structure
	*types.Data
	Notifications []*PersistentNotification
}

// ExtractFrom extracts the PersistentNotificationList from the given readable
func (notificationList *PersistentNotificationList) ExtractFrom(readable types.Readable) error {
	notifications, err := nex.StreamReadListStructure(stream, NewPersistentNotification())
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotificationList.Notifications. %s", err.Error())
	}

	notificationList.Notifications = notifications

	return nil
}

// Copy returns a new copied instance of PersistentNotificationList
func (notificationList *PersistentNotificationList) Copy() types.RVType {
	copied := NewPersistentNotificationList()

	copied.StructureVersion = notificationList.StructureVersion

	copied.Data = notificationList.Data.Copy().(*types.Data)

	copied.Notifications = make([]*PersistentNotification, len(notificationList.Notifications))

	for i := 0; i < len(notificationList.Notifications); i++ {
		copied.Notifications[i] = notificationList.Notifications[i].Copy().(*PersistentNotification)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notificationList *PersistentNotificationList) Equals(o types.RVType) bool {
	if _, ok := o.(*PersistentNotificationList); !ok {
		return false
	}

	other := o.(*PersistentNotificationList)

	if notificationList.StructureVersion != other.StructureVersion {
		return false
	}

	if !notificationList.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(notificationList.Notifications) != len(other.Notifications) {
		return false
	}

	for i := 0; i < len(notificationList.Notifications); i++ {
		if !notificationList.Notifications[i].Equals(other.Notifications[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (notificationList *PersistentNotificationList) String() string {
	return notificationList.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (notificationList *PersistentNotificationList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentNotificationList{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, notificationList.StructureVersion))

	if len(notificationList.Notifications) == 0 {
		b.WriteString(fmt.Sprintf("%sNotifications: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNotifications: [\n", indentationValues))

		for i := 0; i < len(notificationList.Notifications); i++ {
			str := notificationList.Notifications[i].FormatToString(indentationLevel + 2)
			if i == len(notificationList.Notifications)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentNotificationList returns a new PersistentNotificationList
func NewPersistentNotificationList() *PersistentNotificationList {
	return &PersistentNotificationList{}
}
