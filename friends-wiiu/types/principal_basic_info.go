// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PrincipalBasicInfo contains user account and Mii data
type PrincipalBasicInfo struct {
	types.Structure
	*types.Data
	PID     *types.PID
	NNID    string
	Mii     *MiiV2
	Unknown *types.PrimitiveU8
}

// WriteTo writes the PrincipalBasicInfo to the given writable
func (principalInfo *PrincipalBasicInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	principalInfo.PID.WriteTo(contentWritable)
	principalInfo.NNID.WriteTo(contentWritable)
	principalInfo.Mii.WriteTo(contentWritable)
	principalInfo.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	principalInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrincipalBasicInfo from the given readable
func (principalInfo *PrincipalBasicInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = principalInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read PrincipalBasicInfo header. %s", err.Error())
	}

	err = principalInfo.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.PID. %s", err.Error())
	}

	err = principalInfo.NNID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.NNID. %s", err.Error())
	}

	err = principalInfo.Mii.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Mii. %s", err.Error())
	}

	err = principalInfo.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalBasicInfo
func (principalInfo *PrincipalBasicInfo) Copy() types.RVType {
	copied := NewPrincipalBasicInfo()

	copied.StructureVersion = principalInfo.StructureVersion

	copied.Data = principalInfo.Data.Copy().(*types.Data)

	copied.PID = principalInfo.PID.Copy()
	copied.NNID = principalInfo.NNID
	copied.Mii = principalInfo.Mii.Copy().(*MiiV2)
	copied.Unknown = principalInfo.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalInfo *PrincipalBasicInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*PrincipalBasicInfo); !ok {
		return false
	}

	other := o.(*PrincipalBasicInfo)

	if principalInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !principalInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !principalInfo.PID.Equals(other.PID) {
		return false
	}

	if !principalInfo.NNID.Equals(other.NNID) {
		return false
	}

	if !principalInfo.Mii.Equals(other.Mii) {
		return false
	}

	if !principalInfo.Unknown.Equals(other.Unknown) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (principalInfo *PrincipalBasicInfo) String() string {
	return principalInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (principalInfo *PrincipalBasicInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalBasicInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, principalInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, principalInfo.PID))
	b.WriteString(fmt.Sprintf("%sNNID: %q,\n", indentationValues, principalInfo.NNID))

	if principalInfo.Mii != nil {
		b.WriteString(fmt.Sprintf("%sMii: %s,\n", indentationValues, principalInfo.Mii.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMii: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown: %d\n", indentationValues, principalInfo.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalBasicInfo returns a new PrincipalBasicInfo
func NewPrincipalBasicInfo() *PrincipalBasicInfo {
	return &PrincipalBasicInfo{}
}
