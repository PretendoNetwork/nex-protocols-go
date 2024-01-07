// Package types implements all the types used by the Utility protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	types.Structure
	NexUniqueID         *types.PrimitiveU64
	NexUniqueIDPassword *types.PrimitiveU64
}

// WriteTo writes the UniqueIDInfo to the given writable
func (uniqueIDInfo *UniqueIDInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uniqueIDInfo.NexUniqueID.WriteTo(contentWritable)
	uniqueIDInfo.NexUniqueIDPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	uniqueIDInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UniqueIDInfo from the given readable
func (uniqueIDInfo *UniqueIDInfo) ExtractFrom(readable types.Readable) error {
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
func (uniqueIDInfo *UniqueIDInfo) Copy() types.RVType {
	copied := NewUniqueIDInfo()

	copied.StructureVersion = uniqueIDInfo.StructureVersion

	copied.NexUniqueID = uniqueIDInfo.NexUniqueID
	copied.NexUniqueIDPassword = uniqueIDInfo.NexUniqueIDPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (uniqueIDInfo *UniqueIDInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*UniqueIDInfo); !ok {
		return false
	}

	other := o.(*UniqueIDInfo)

	if uniqueIDInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !uniqueIDInfo.NexUniqueID.Equals(other.NexUniqueID) {
		return false
	}

	if !uniqueIDInfo.NexUniqueIDPassword.Equals(other.NexUniqueIDPassword) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, uniqueIDInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %d,\n", indentationValues, uniqueIDInfo.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sNexUniqueIDPassword: %d\n", indentationValues, uniqueIDInfo.NexUniqueIDPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	return &UniqueIDInfo{}
}
