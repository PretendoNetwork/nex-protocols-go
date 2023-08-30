// Package types implements all the types used by the Utility protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	nex.Structure
	NexUniqueID         uint64
	NexUniqueIDPassword uint64
}

// Bytes encodes the UniqueIDInfo and returns a byte array
func (uniqueIDInfo *UniqueIDInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueID)
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueIDPassword)

	return stream.Bytes()
}

// ExtractFromStream extracts a UniqueIDInfo structure from a stream
func (uniqueIDInfo *UniqueIDInfo) ExtractFromStream(stream *nex.StreamIn) error {
	nexUniqueID, err := stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UniqueIDInfo.NexUniqueID from stream. %s", err.Error())
	}

	nexUniqueIDPassword, err := stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UniqueIDInfo.NexUniqueIDPassword from stream. %s", err.Error())
	}

	uniqueIDInfo.NexUniqueID = nexUniqueID
	uniqueIDInfo.NexUniqueIDPassword = nexUniqueIDPassword

	return nil
}

// Copy returns a new copied instance of UniqueIDInfo
func (uniqueIDInfo *UniqueIDInfo) Copy() nex.StructureInterface {
	copied := NewUniqueIDInfo()

	copied.SetStructureVersion(uniqueIDInfo.StructureVersion())

	copied.NexUniqueID = uniqueIDInfo.NexUniqueID
	copied.NexUniqueIDPassword = uniqueIDInfo.NexUniqueIDPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (uniqueIDInfo *UniqueIDInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*UniqueIDInfo)

	if uniqueIDInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if uniqueIDInfo.NexUniqueID != other.NexUniqueID {
		return false
	}

	if uniqueIDInfo.NexUniqueIDPassword != other.NexUniqueIDPassword {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (uniqueIDInfo *UniqueIDInfo) String() string {
	return uniqueIDInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (uniqueIDInfo *UniqueIDInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UniqueIDInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, uniqueIDInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %d,\n", indentationValues, uniqueIDInfo.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sNexUniqueIDPassword: %d\n", indentationValues, uniqueIDInfo.NexUniqueIDPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	return &UniqueIDInfo{}
}
