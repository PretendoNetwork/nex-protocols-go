// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// BlacklistedPrincipal contains information about a blocked user
type BlacklistedPrincipal struct {
	types.Structure
	*types.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	GameKey            *GameKey
	BlackListedSince   *types.DateTime
}

// WriteTo writes the BlacklistedPrincipal to the given writable
func (blacklistedPrincipal *BlacklistedPrincipal) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	blacklistedPrincipal.PrincipalBasicInfo.WriteTo(contentWritable)
	blacklistedPrincipal.GameKey.WriteTo(contentWritable)
	blacklistedPrincipal.BlackListedSince.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	blacklistedPrincipal.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BlacklistedPrincipal from the given readable
func (blacklistedPrincipal *BlacklistedPrincipal) ExtractFrom(readable types.Readable) error {
	var err error

	if err = blacklistedPrincipal.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BlacklistedPrincipal header. %s", err.Error())
	}

	err = blacklistedPrincipal.PrincipalBasicInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.PrincipalBasicInfo. %s", err.Error())
	}

	err = blacklistedPrincipal.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.GameKey. %s", err.Error())
	}

	err = blacklistedPrincipal.BlackListedSince.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.BlackListedSince. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BlacklistedPrincipal
func (blacklistedPrincipal *BlacklistedPrincipal) Copy() types.RVType {
	copied := NewBlacklistedPrincipal()

	copied.StructureVersion = blacklistedPrincipal.StructureVersion

	copied.Data = blacklistedPrincipal.Data.Copy().(*types.Data)

	copied.PrincipalBasicInfo = blacklistedPrincipal.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.GameKey = blacklistedPrincipal.GameKey.Copy().(*GameKey)
	copied.BlackListedSince = blacklistedPrincipal.BlackListedSince.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (blacklistedPrincipal *BlacklistedPrincipal) Equals(o types.RVType) bool {
	if _, ok := o.(*BlacklistedPrincipal); !ok {
		return false
	}

	other := o.(*BlacklistedPrincipal)

	if blacklistedPrincipal.StructureVersion != other.StructureVersion {
		return false
	}

	if !blacklistedPrincipal.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !blacklistedPrincipal.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if !blacklistedPrincipal.GameKey.Equals(other.GameKey) {
		return false
	}

	if !blacklistedPrincipal.BlackListedSince.Equals(other.BlackListedSince) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (blacklistedPrincipal *BlacklistedPrincipal) String() string {
	return blacklistedPrincipal.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (blacklistedPrincipal *BlacklistedPrincipal) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BlacklistedPrincipal{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, blacklistedPrincipal.StructureVersion))

	if blacklistedPrincipal.PrincipalBasicInfo != nil {
		b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: %s,\n", indentationValues, blacklistedPrincipal.PrincipalBasicInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: nil,\n", indentationValues))
	}

	if blacklistedPrincipal.GameKey != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, blacklistedPrincipal.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil,\n", indentationValues))
	}

	if blacklistedPrincipal.BlackListedSince != nil {
		b.WriteString(fmt.Sprintf("%sBlackListedSince: %s\n", indentationValues, blacklistedPrincipal.BlackListedSince.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBlackListedSince: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBlacklistedPrincipal returns a new BlacklistedPrincipal
func NewBlacklistedPrincipal() *BlacklistedPrincipal {
	return &BlacklistedPrincipal{}
}
