// Package types implements all the types used by the Ticket Granting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NintendoLoginData holds a nex auth token
type NintendoLoginData struct {
	types.Structure
	Token string
}

// ExtractFrom extracts the NintendoLoginData from the given readable
func (nintendoLoginData *NintendoLoginData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = nintendoLoginData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read NintendoLoginData header. %s", err.Error())
	}

	err = nintendoLoginData.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoLoginData.Token. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoLoginData
func (nintendoLoginData *NintendoLoginData) Copy() types.RVType {
	copied := NewNintendoLoginData()

	copied.StructureVersion = nintendoLoginData.StructureVersion

	copied.Token = nintendoLoginData.Token

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoLoginData *NintendoLoginData) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoLoginData); !ok {
		return false
	}

	other := o.(*NintendoLoginData)

	if nintendoLoginData.StructureVersion != other.StructureVersion {
		return false
	}

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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, nintendoLoginData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sToken: %s\n", indentationValues, nintendoLoginData.Token))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoLoginData returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	return &NintendoLoginData{}
}
