// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DeletionEntry holds an entry for a deletion
type DeletionEntry struct {
	types.Structure
	IDGathering *types.PrimitiveU32
	PID         *types.PID
	UIReason    *types.PrimitiveU32
}

// ExtractFrom extracts the DeletionEntry from the given readable
func (deletionEntry *DeletionEntry) ExtractFrom(readable types.Readable) error {
	var err error

	if err = deletionEntry.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DeletionEntry header. %s", err.Error())
	}

	err = deletionEntry.IDGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.IDGathering. %s", err.Error())
	}

	err = deletionEntry.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.PID. %s", err.Error())
	}

	err = deletionEntry.UIReason.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.UIReason. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DeletionEntry to the given writable
func (deletionEntry *DeletionEntry) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	deletionEntry.IDGathering.WriteTo(contentWritable)
	deletionEntry.PID.WriteTo(contentWritable)
	deletionEntry.UIReason.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	deletionEntry.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DeletionEntry
func (deletionEntry *DeletionEntry) Copy() types.RVType {
	copied := NewDeletionEntry()

	copied.StructureVersion = deletionEntry.StructureVersion

	copied.IDGathering = deletionEntry.IDGathering
	copied.PID = deletionEntry.PID.Copy()
	copied.UIReason = deletionEntry.UIReason

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (deletionEntry *DeletionEntry) Equals(o types.RVType) bool {
	if _, ok := o.(*DeletionEntry); !ok {
		return false
	}

	other := o.(*DeletionEntry)

	if deletionEntry.StructureVersion != other.StructureVersion {
		return false
	}

	if !deletionEntry.IDGathering.Equals(other.IDGathering) {
		return false
	}

	if !deletionEntry.PID.Equals(other.PID) {
		return false
	}

	if !deletionEntry.UIReason.Equals(other.UIReason) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (deletionEntry *DeletionEntry) String() string {
	return deletionEntry.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (deletionEntry *DeletionEntry) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DeletionEntry{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, deletionEntry.StructureVersion))
	b.WriteString(fmt.Sprintf("%sIDGathering: %d,\n", indentationValues, deletionEntry.IDGathering))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, deletionEntry.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUIReason: %d\n", indentationValues, deletionEntry.UIReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDeletionEntry returns a new DeletionEntry
func NewDeletionEntry() *DeletionEntry {
	return &DeletionEntry{}
}
