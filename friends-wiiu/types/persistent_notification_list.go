// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PersistentNotificationList is a type within the FriendsWiiU protocol
type PersistentNotificationList struct {
	types.Structure
	types.Data
	Notifications types.List[PersistentNotification]
}

// ObjectID returns the object identifier of the type
func (pnl PersistentNotificationList) ObjectID() types.RVType {
	return pnl.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (pnl PersistentNotificationList) DataObjectID() types.RVType {
	return types.NewString("PersistentNotificationList")
}

// WriteTo writes the PersistentNotificationList to the given writable
func (pnl PersistentNotificationList) WriteTo(writable types.Writable) {
	pnl.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	pnl.Notifications.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pnl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PersistentNotificationList from the given readable
func (pnl *PersistentNotificationList) ExtractFrom(readable types.Readable) error {
	var err error

	err = pnl.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotificationList.Data. %s", err.Error())
	}

	err = pnl.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotificationList header. %s", err.Error())
	}

	err = pnl.Notifications.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotificationList.Notifications. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentNotificationList
func (pnl PersistentNotificationList) Copy() types.RVType {
	copied := NewPersistentNotificationList()

	copied.StructureVersion = pnl.StructureVersion
	copied.Data = pnl.Data.Copy().(types.Data)
	copied.Notifications = pnl.Notifications.Copy().(types.List[PersistentNotification])

	return copied
}

// Equals checks if the given PersistentNotificationList contains the same data as the current PersistentNotificationList
func (pnl PersistentNotificationList) Equals(o types.RVType) bool {
	if _, ok := o.(PersistentNotificationList); !ok {
		return false
	}

	other := o.(PersistentNotificationList)

	if pnl.StructureVersion != other.StructureVersion {
		return false
	}

	if !pnl.Data.Equals(other.Data) {
		return false
	}

	return pnl.Notifications.Equals(other.Notifications)
}

// CopyRef copies the current value of the PersistentNotificationList
// and returns a pointer to the new copy
func (pnl PersistentNotificationList) CopyRef() types.RVTypePtr {
	copied := pnl.Copy().(PersistentNotificationList)
	return &copied
}

// Deref takes a pointer to the PersistentNotificationList
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pnl *PersistentNotificationList) Deref() types.RVType {
	return *pnl
}

// String returns the string representation of the PersistentNotificationList
func (pnl PersistentNotificationList) String() string {
	return pnl.FormatToString(0)
}

// FormatToString pretty-prints the PersistentNotificationList using the provided indentation level
func (pnl PersistentNotificationList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentNotificationList{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, pnl.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNotifications: %s,\n", indentationValues, pnl.Notifications))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentNotificationList returns a new PersistentNotificationList
func NewPersistentNotificationList() PersistentNotificationList {
	return PersistentNotificationList{
		Data:          types.NewData(),
		Notifications: types.NewList[PersistentNotification](),
	}

}
