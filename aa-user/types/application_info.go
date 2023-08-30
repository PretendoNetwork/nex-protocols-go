// Package types implements all the types used by the AAUser protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ApplicationInfo contains the title ID and version for a title
type ApplicationInfo struct {
	nex.Structure
	*nex.Data
	TitleID      uint64
	TitleVersion uint16
}

// Bytes encodes the ApplicationInfo and returns a byte array
func (applicationInfo *ApplicationInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(applicationInfo.TitleID)
	stream.WriteUInt16LE(applicationInfo.TitleVersion)

	return stream.Bytes()
}

// ExtractFromStream extracts a ApplicationInfo structure from a stream
func (applicationInfo *ApplicationInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	applicationInfo.TitleID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleID. %s", err.Error())
	}

	applicationInfo.TitleVersion, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ApplicationInfo
func (applicationInfo *ApplicationInfo) Copy() nex.StructureInterface {
	copied := NewApplicationInfo()

	copied.SetStructureVersion(applicationInfo.StructureVersion())

	copied.Data = applicationInfo.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.TitleID = applicationInfo.TitleID
	copied.TitleVersion = applicationInfo.TitleVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (applicationInfo *ApplicationInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ApplicationInfo)

	if applicationInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !applicationInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if applicationInfo.TitleID != other.TitleID {
		return false
	}

	if applicationInfo.TitleVersion != other.TitleVersion {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (applicationInfo *ApplicationInfo) String() string {
	return applicationInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (applicationInfo *ApplicationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ApplicationInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, applicationInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTitleID: %d,\n", indentationValues, applicationInfo.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %d\n", indentationValues, applicationInfo.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewApplicationInfo returns a new ApplicationInfo
func NewApplicationInfo() *ApplicationInfo {
	return &ApplicationInfo{}
}
