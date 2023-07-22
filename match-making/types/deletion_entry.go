// Package match_making_types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package match_making_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DeletionEntry holds an entry for a deletion
type DeletionEntry struct {
	nex.Structure
	IDGathering uint32
	PID         uint32
	UIReason    uint32
}

// ExtractFromStream extracts a DeletionEntry structure from a stream
func (deletionEntry *DeletionEntry) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	deletionEntry.IDGathering, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.IDGathering. %s", err.Error())
	}

	deletionEntry.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.PID. %s", err.Error())
	}

	deletionEntry.UIReason, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DeletionEntry.UIReason. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DeletionEntry and returns a byte array
func (deletionEntry *DeletionEntry) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(deletionEntry.IDGathering)
	stream.WriteUInt32LE(deletionEntry.PID)
	stream.WriteUInt32LE(deletionEntry.UIReason)

	return stream.Bytes()
}

// Copy returns a new copied instance of DeletionEntry
func (deletionEntry *DeletionEntry) Copy() nex.StructureInterface {
	copied := NewDeletionEntry()

	copied.IDGathering = deletionEntry.IDGathering
	copied.PID = deletionEntry.PID
	copied.UIReason = deletionEntry.UIReason

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (deletionEntry *DeletionEntry) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DeletionEntry)

	if deletionEntry.IDGathering != other.IDGathering {
		return false
	}

	if deletionEntry.PID != other.PID {
		return false
	}

	if deletionEntry.UIReason != other.UIReason {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, deletionEntry.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sIDGathering: %d,\n", indentationValues, deletionEntry.IDGathering))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, deletionEntry.PID))
	b.WriteString(fmt.Sprintf("%sUIReason: %d\n", indentationValues, deletionEntry.UIReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDeletionEntry returns a new DeletionEntry
func NewDeletionEntry() *DeletionEntry {
	return &DeletionEntry{}
}
