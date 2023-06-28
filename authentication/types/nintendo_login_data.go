package authentication_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// NintendoLoginData holds a nex auth token
type NintendoLoginData struct {
	nex.Structure
	Token string
}

// ExtractFromStream extracts a AuthenticationInfo structure from a stream
func (nintendoLoginData *NintendoLoginData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nintendoLoginData.Token, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoLoginData.Token. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoLoginData
func (nintendoLoginData *NintendoLoginData) Copy() nex.StructureInterface {
	copied := NewNintendoLoginData()

	copied.Token = nintendoLoginData.Token

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoLoginData *NintendoLoginData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoLoginData)

	return nintendoLoginData.Token == other.Token
}

// NewAuthenticationInfo returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	return &NintendoLoginData{}
}
