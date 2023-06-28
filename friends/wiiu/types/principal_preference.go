package friends_wiiu_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// PrincipalPreference contains unknown data
type PrincipalPreference struct {
	nex.Structure
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

	copied.ShowOnlinePresence = principalPreference.ShowOnlinePresence
	copied.ShowCurrentTitle = principalPreference.ShowCurrentTitle
	copied.BlockFriendRequests = principalPreference.BlockFriendRequests

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalPreference *PrincipalPreference) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalPreference)

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

// NewPrincipalPreference returns a new PrincipalPreference
func NewPrincipalPreference() *PrincipalPreference {
	return &PrincipalPreference{}
}
