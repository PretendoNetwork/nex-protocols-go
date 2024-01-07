// Package types implements all the types used by the AAUser protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// ApplicationInfo contains the title ID and version for a title
type ApplicationInfo struct {
	types.Structure
	*types.Data
	TitleID      *types.PrimitiveU64
	TitleVersion *types.PrimitiveU16
}

// WriteTo writes the ApplicationInfo to the given writable
func (applicationInfo *ApplicationInfo) WriteTo(writable types.Writable) {
	applicationInfo.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	applicationInfo.TitleID.WriteTo(contentWritable)
	applicationInfo.TitleVersion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	applicationInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ApplicationInfo from the given readable
func (applicationInfo *ApplicationInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = applicationInfo.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.Data. %s", err.Error())
	}

	if err = applicationInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ApplicationInfo header. %s", err.Error())
	}

	err = applicationInfo.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleID. %s", err.Error())
	}

	err = applicationInfo.TitleVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ApplicationInfo.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ApplicationInfo
func (applicationInfo *ApplicationInfo) Copy() types.RVType {
	copied := NewApplicationInfo()

	copied.StructureVersion = applicationInfo.StructureVersion

	copied.Data = applicationInfo.Data.Copy().(*types.Data)

	copied.TitleID = applicationInfo.TitleID.Copy().(*types.PrimitiveU64)
	copied.TitleVersion = applicationInfo.TitleVersion.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (applicationInfo *ApplicationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ApplicationInfo); !ok {
		return false
	}

	other := o.(*ApplicationInfo)

	if applicationInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !applicationInfo.Data.Equals(other.Data) {
		return false
	}

	if !applicationInfo.TitleID.Equals(other.TitleID) {
		return false
	}

	if !applicationInfo.TitleVersion.Equals(other.TitleVersion) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, applicationInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTitleID: %s,\n", indentationValues, applicationInfo.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %s\n", indentationValues, applicationInfo.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewApplicationInfo returns a new ApplicationInfo
func NewApplicationInfo() *ApplicationInfo {
	return &ApplicationInfo{
		Data: types.NewData(),
		TitleID: types.NewPrimitiveU64(0),
		TitleVersion: types.NewPrimitiveU16(0),
	}
}
