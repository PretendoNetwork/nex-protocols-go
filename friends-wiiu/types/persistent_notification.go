// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PersistentNotification contains unknown data
type PersistentNotification struct {
	types.Structure
	*types.Data
	Unknown1 *types.PrimitiveU64
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU32
	Unknown5 string
}

// ExtractFrom extracts the PersistentNotification from the given readable
func (notification *PersistentNotification) ExtractFrom(readable types.Readable) error {
	var err error

	if err = notification.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read PersistentNotification header. %s", err.Error())
	}

	err = notification.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown1. %s", err.Error())
	}

	err = notification.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown2. %s", err.Error())
	}

	err = notification.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown3. %s", err.Error())
	}

	err = notification.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown4. %s", err.Error())
	}

	err = notification.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown5. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentNotification
func (notification *PersistentNotification) Copy() types.RVType {
	copied := NewPersistentNotification()

	copied.StructureVersion = notification.StructureVersion

	copied.Data = notification.Data.Copy().(*types.Data)

	copied.Unknown1 = notification.Unknown1
	copied.Unknown2 = notification.Unknown2
	copied.Unknown3 = notification.Unknown3
	copied.Unknown4 = notification.Unknown4
	copied.Unknown5 = notification.Unknown5

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notification *PersistentNotification) Equals(o types.RVType) bool {
	if _, ok := o.(*PersistentNotification); !ok {
		return false
	}

	other := o.(*PersistentNotification)

	if notification.StructureVersion != other.StructureVersion {
		return false
	}

	if !notification.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !notification.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !notification.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !notification.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !notification.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !notification.Unknown5.Equals(other.Unknown5) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, notification.StructureVersion))
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
