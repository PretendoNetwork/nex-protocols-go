// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PersistentNotification contains unknown data
type PersistentNotification struct {
	nex.Structure
	*nex.Data
	Unknown1 uint64
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 string
}

// ExtractFromStream extracts a PersistentNotification structure from a stream
func (notification *PersistentNotification) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	notification.Unknown1, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown1. %s", err.Error())
	}

	notification.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown2. %s", err.Error())
	}

	notification.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown3. %s", err.Error())
	}

	notification.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown4. %s", err.Error())
	}

	notification.Unknown5, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown5. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentNotification
func (notification *PersistentNotification) Copy() nex.StructureInterface {
	copied := NewPersistentNotification()

	copied.SetStructureVersion(notification.StructureVersion())

	if notification.ParentType() != nil {
		copied.Data = notification.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.Unknown1 = notification.Unknown1
	copied.Unknown2 = notification.Unknown2
	copied.Unknown3 = notification.Unknown3
	copied.Unknown4 = notification.Unknown4
	copied.Unknown5 = notification.Unknown5

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notification *PersistentNotification) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PersistentNotification)

	if notification.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !notification.ParentType().Equals(other.ParentType()) {
		return false
	}

	if notification.Unknown1 != other.Unknown1 {
		return false
	}

	if notification.Unknown2 != other.Unknown2 {
		return false
	}

	if notification.Unknown3 != other.Unknown3 {
		return false
	}

	if notification.Unknown4 != other.Unknown4 {
		return false
	}

	if notification.Unknown5 != other.Unknown5 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (notification *PersistentNotification) String() string {
	return notification.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (notification *PersistentNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PersistentNotification{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, notification.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, notification.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, notification.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, notification.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d,\n", indentationValues, notification.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %q\n", indentationValues, notification.Unknown5))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPersistentNotification returns a new PersistentNotification
func NewPersistentNotification() *PersistentNotification {
	return &PersistentNotification{}
}
