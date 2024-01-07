// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PrincipalPreference contains unknown data
type PrincipalPreference struct {
	types.Structure
	*types.Data
	ShowOnlinePresence  *types.PrimitiveBool
	ShowCurrentTitle    *types.PrimitiveBool
	BlockFriendRequests *types.PrimitiveBool
}

// WriteTo writes the PrincipalPreference to the given writable
func (principalPreference *PrincipalPreference) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	principalPreference.ShowOnlinePresence.WriteTo(contentWritable)
	principalPreference.ShowCurrentTitle.WriteTo(contentWritable)
	principalPreference.BlockFriendRequests.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	principalPreference.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrincipalPreference from the given readable
func (principalPreference *PrincipalPreference) ExtractFrom(readable types.Readable) error {
	var err error

	if err = principalPreference.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read PrincipalPreference header. %s", err.Error())
	}

	err = principalPreference.ShowOnlinePresence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowOnlinePresence. %s", err.Error())
	}

	err = principalPreference.ShowCurrentTitle.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowCurrentTitle. %s", err.Error())
	}

	err = principalPreference.BlockFriendRequests.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.BlockFriendRequests. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalPreference
func (principalPreference *PrincipalPreference) Copy() types.RVType {
	copied := NewPrincipalPreference()

	copied.StructureVersion = principalPreference.StructureVersion

	copied.Data = principalPreference.Data.Copy().(*types.Data)

	copied.ShowOnlinePresence = principalPreference.ShowOnlinePresence
	copied.ShowCurrentTitle = principalPreference.ShowCurrentTitle
	copied.BlockFriendRequests = principalPreference.BlockFriendRequests

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalPreference *PrincipalPreference) Equals(o types.RVType) bool {
	if _, ok := o.(*PrincipalPreference); !ok {
		return false
	}

	other := o.(*PrincipalPreference)

	if principalPreference.StructureVersion != other.StructureVersion {
		return false
	}

	if !principalPreference.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !principalPreference.ShowOnlinePresence.Equals(other.ShowOnlinePresence) {
		return false
	}

	if !principalPreference.ShowCurrentTitle.Equals(other.ShowCurrentTitle) {
		return false
	}

	if !principalPreference.BlockFriendRequests.Equals(other.BlockFriendRequests) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (principalPreference *PrincipalPreference) String() string {
	return principalPreference.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (principalPreference *PrincipalPreference) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalPreference{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, principalPreference.StructureVersion))
	b.WriteString(fmt.Sprintf("%sShowOnlinePresence: %t,\n", indentationValues, principalPreference.ShowOnlinePresence))
	b.WriteString(fmt.Sprintf("%sShowCurrentTitle: %t,\n", indentationValues, principalPreference.ShowCurrentTitle))
	b.WriteString(fmt.Sprintf("%sBlockFriendRequests: %t\n", indentationValues, principalPreference.BlockFriendRequests))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalPreference returns a new PrincipalPreference
func NewPrincipalPreference() *PrincipalPreference {
	return &PrincipalPreference{}
}
