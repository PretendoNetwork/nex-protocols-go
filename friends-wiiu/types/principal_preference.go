// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PrincipalPreference is a type within the FriendsWiiU protocol
type PrincipalPreference struct {
	types.Structure
	types.Data
	ShowOnlinePresence  types.Bool
	ShowCurrentTitle    types.Bool
	BlockFriendRequests types.Bool
}

// ObjectID returns the object identifier of the type
func (pp PrincipalPreference) ObjectID() types.RVType {
	return pp.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (pp PrincipalPreference) DataObjectID() types.RVType {
	return types.NewString("PrincipalPreference")
}

// WriteTo writes the PrincipalPreference to the given writable
func (pp PrincipalPreference) WriteTo(writable types.Writable) {
	pp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	pp.ShowOnlinePresence.WriteTo(contentWritable)
	pp.ShowCurrentTitle.WriteTo(contentWritable)
	pp.BlockFriendRequests.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrincipalPreference from the given readable
func (pp *PrincipalPreference) ExtractFrom(readable types.Readable) error {
	var err error

	err = pp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.Data. %s", err.Error())
	}

	err = pp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference header. %s", err.Error())
	}

	err = pp.ShowOnlinePresence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowOnlinePresence. %s", err.Error())
	}

	err = pp.ShowCurrentTitle.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowCurrentTitle. %s", err.Error())
	}

	err = pp.BlockFriendRequests.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.BlockFriendRequests. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalPreference
func (pp PrincipalPreference) Copy() types.RVType {
	copied := NewPrincipalPreference()

	copied.StructureVersion = pp.StructureVersion
	copied.Data = pp.Data.Copy().(types.Data)
	copied.ShowOnlinePresence = pp.ShowOnlinePresence.Copy().(types.Bool)
	copied.ShowCurrentTitle = pp.ShowCurrentTitle.Copy().(types.Bool)
	copied.BlockFriendRequests = pp.BlockFriendRequests.Copy().(types.Bool)

	return copied
}

// Equals checks if the given PrincipalPreference contains the same data as the current PrincipalPreference
func (pp PrincipalPreference) Equals(o types.RVType) bool {
	if _, ok := o.(PrincipalPreference); !ok {
		return false
	}

	other := o.(PrincipalPreference)

	if pp.StructureVersion != other.StructureVersion {
		return false
	}

	if !pp.Data.Equals(other.Data) {
		return false
	}

	if !pp.ShowOnlinePresence.Equals(other.ShowOnlinePresence) {
		return false
	}

	if !pp.ShowCurrentTitle.Equals(other.ShowCurrentTitle) {
		return false
	}

	return pp.BlockFriendRequests.Equals(other.BlockFriendRequests)
}

// CopyRef copies the current value of the PrincipalPreference
// and returns a pointer to the new copy
func (pp PrincipalPreference) CopyRef() types.RVTypePtr {
	copied := pp.Copy().(PrincipalPreference)
	return &copied
}

// Deref takes a pointer to the PrincipalPreference
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pp *PrincipalPreference) Deref() types.RVType {
	return *pp
}

// String returns the string representation of the PrincipalPreference
func (pp PrincipalPreference) String() string {
	return pp.FormatToString(0)
}

// FormatToString pretty-prints the PrincipalPreference using the provided indentation level
func (pp PrincipalPreference) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalPreference{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, pp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sShowOnlinePresence: %s,\n", indentationValues, pp.ShowOnlinePresence))
	b.WriteString(fmt.Sprintf("%sShowCurrentTitle: %s,\n", indentationValues, pp.ShowCurrentTitle))
	b.WriteString(fmt.Sprintf("%sBlockFriendRequests: %s,\n", indentationValues, pp.BlockFriendRequests))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalPreference returns a new PrincipalPreference
func NewPrincipalPreference() PrincipalPreference {
	return PrincipalPreference{
		Data:                types.NewData(),
		ShowOnlinePresence:  types.NewBool(false),
		ShowCurrentTitle:    types.NewBool(false),
		BlockFriendRequests: types.NewBool(false),
	}

}
