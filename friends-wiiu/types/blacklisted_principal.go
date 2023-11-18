// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// BlacklistedPrincipal contains information about a blocked user
type BlacklistedPrincipal struct {
	nex.Structure
	*nex.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	GameKey            *GameKey
	BlackListedSince   *nex.DateTime
}

// Bytes encodes the BlacklistedPrincipal and returns a byte array
func (blacklistedPrincipal *BlacklistedPrincipal) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(blacklistedPrincipal.PrincipalBasicInfo)
	stream.WriteStructure(blacklistedPrincipal.GameKey)
	stream.WriteDateTime(blacklistedPrincipal.BlackListedSince)

	return stream.Bytes()
}

// ExtractFromStream extracts a BlacklistedPrincipal structure from a stream
func (blacklistedPrincipal *BlacklistedPrincipal) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	blacklistedPrincipal.PrincipalBasicInfo, err = nex.StreamReadStructure(stream, NewPrincipalBasicInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.PrincipalBasicInfo. %s", err.Error())
	}

	blacklistedPrincipal.GameKey, err = nex.StreamReadStructure(stream, NewGameKey())
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.GameKey. %s", err.Error())
	}

	blacklistedPrincipal.BlackListedSince, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.BlackListedSince. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BlacklistedPrincipal
func (blacklistedPrincipal *BlacklistedPrincipal) Copy() nex.StructureInterface {
	copied := NewBlacklistedPrincipal()

	copied.SetStructureVersion(blacklistedPrincipal.StructureVersion())

	if blacklistedPrincipal.ParentType() != nil {
		copied.Data = blacklistedPrincipal.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.PrincipalBasicInfo = blacklistedPrincipal.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.GameKey = blacklistedPrincipal.GameKey.Copy().(*GameKey)
	copied.BlackListedSince = blacklistedPrincipal.BlackListedSince.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (blacklistedPrincipal *BlacklistedPrincipal) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BlacklistedPrincipal)

	if blacklistedPrincipal.StructureVersion() != other.StructureVersion() {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, blacklistedPrincipal.StructureVersion()))

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
