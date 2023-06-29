package authentication_types

import (
	"fmt"
	"strings"

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

// String returns a string representation of the struct
func (nintendoLoginData *NintendoLoginData) String() string {
	return nintendoLoginData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (nintendoLoginData *NintendoLoginData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoLoginData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, nintendoLoginData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sToken: %s\n", indentationValues, nintendoLoginData.Token))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoLoginData returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	return &NintendoLoginData{}
}
