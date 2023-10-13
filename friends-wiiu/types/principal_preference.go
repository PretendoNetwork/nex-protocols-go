// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PrincipalPreference contains unknown data
type PrincipalPreference struct {
	nex.Structure
	*nex.Data
	ShowOnlinePresence  bool
	ShowCurrentTitle    bool
	BlockFriendRequests bool
}

// Bytes encodes the PrincipalPreference and returns a byte array
func (principalPreference *PrincipalPreference) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteBool(principalPreference.ShowOnlinePresence)
	stream.WriteBool(principalPreference.ShowCurrentTitle)
	stream.WriteBool(principalPreference.BlockFriendRequests)

	return stream.Bytes()
}

// ExtractFromStream extracts a PrincipalPreference structure from a stream
func (principalPreference *PrincipalPreference) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalPreference.ShowOnlinePresence, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowOnlinePresence. %s", err.Error())
	}

	principalPreference.ShowCurrentTitle, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowCurrentTitle. %s", err.Error())
	}

	principalPreference.BlockFriendRequests, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.BlockFriendRequests. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalPreference
func (principalPreference *PrincipalPreference) Copy() nex.StructureInterface {
	copied := NewPrincipalPreference()

	copied.SetStructureVersion(principalPreference.StructureVersion())

	if principalPreference.ParentType() != nil {
		copied.Data = principalPreference.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.ShowOnlinePresence = principalPreference.ShowOnlinePresence
	copied.ShowCurrentTitle = principalPreference.ShowCurrentTitle
	copied.BlockFriendRequests = principalPreference.BlockFriendRequests

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalPreference *PrincipalPreference) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalPreference)

	if principalPreference.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !principalPreference.ParentType().Equals(other.ParentType()) {
		return false
	}

	if principalPreference.ShowOnlinePresence != other.ShowOnlinePresence {
		return false
	}

	if principalPreference.ShowCurrentTitle != other.ShowCurrentTitle {
		return false
	}

	if principalPreference.BlockFriendRequests != other.BlockFriendRequests {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, principalPreference.StructureVersion()))
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
