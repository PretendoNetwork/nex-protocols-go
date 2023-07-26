// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeBlockListParam holds parameters for a matchmake session
type MatchmakeBlockListParam struct {
	nex.Structure
	OptionFlag uint32
}

// ExtractFromStream extracts a MatchmakeBlockListParam structure from a stream
func (matchmakeBlockListParam *MatchmakeBlockListParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeBlockListParam.OptionFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeBlockListParam.OptionFlag. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeBlockListParam
func (matchmakeBlockListParam *MatchmakeBlockListParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeBlockListParam()

	copied.OptionFlag = matchmakeBlockListParam.OptionFlag

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeBlockListParam *MatchmakeBlockListParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeBlockListParam)

	return matchmakeBlockListParam.OptionFlag == other.OptionFlag
}

// String returns a string representation of the struct
func (matchmakeBlockListParam *MatchmakeBlockListParam) String() string {
	return matchmakeBlockListParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeBlockListParam *MatchmakeBlockListParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeBlockListParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeBlockListParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOptionFlag: %d\n", indentationValues, matchmakeBlockListParam.OptionFlag))

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeBlockListParam returns a new MatchmakeBlockListParam
func NewMatchmakeBlockListParam() *MatchmakeBlockListParam {
	return &MatchmakeBlockListParam{}
}
