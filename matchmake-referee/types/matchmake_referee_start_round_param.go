// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"golang.org/x/exp/slices"
)

// MatchmakeRefereeStartRoundParam contains the results of a round
type MatchmakeRefereeStartRoundParam struct {
	nex.Structure
	*nex.Data
	PersonalDataCategory uint32
	GID                  uint32
	PIDs                 []uint32
}

// Bytes encodes the MatchmakeRefereeStartRoundParam and returns a byte array
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(matchmakeRefereeStartRoundParam.PersonalDataCategory)
	stream.WriteUInt32LE(matchmakeRefereeStartRoundParam.GID)
	stream.WriteListUInt32LE(matchmakeRefereeStartRoundParam.PIDs)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereeStartRoundParam structure from a stream
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereeStartRoundParam.PersonalDataCategory, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PersonalDataCategory. %s", err.Error())
	}

	matchmakeRefereeStartRoundParam.GID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.GID. %s", err.Error())
	}

	matchmakeRefereeStartRoundParam.PIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStartRoundParam
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereeStartRoundParam()

	copied.Data = matchmakeRefereeStartRoundParam.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.PersonalDataCategory = matchmakeRefereeStartRoundParam.PersonalDataCategory
	copied.GID = matchmakeRefereeStartRoundParam.GID
	copied.PIDs = make([]uint32, len(matchmakeRefereeStartRoundParam.PIDs))
	copy(copied.PIDs, matchmakeRefereeStartRoundParam.PIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereeStartRoundParam)

	if !matchmakeRefereeStartRoundParam.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereeStartRoundParam.PersonalDataCategory != other.PersonalDataCategory {
		return false
	}

	if matchmakeRefereeStartRoundParam.GID != other.GID {
		return false
	}

	if !slices.Equal(matchmakeRefereeStartRoundParam.PIDs, other.PIDs) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) String() string {
	return matchmakeRefereeStartRoundParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStartRoundParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereeStartRoundParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, matchmakeRefereeStartRoundParam.PersonalDataCategory))
	b.WriteString(fmt.Sprintf("%sPersonalRoundResultFlag: %d,\n", indentationValues, matchmakeRefereeStartRoundParam.GID))
	if len(matchmakeRefereeStartRoundParam.PIDs) == 0 {
		b.WriteString(fmt.Sprintf("%sPIDs: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPIDs: [\n", indentationValues))

		for i := 0; i < len(matchmakeRefereeStartRoundParam.PIDs); i++ {
			str := fmt.Sprintf("%d", matchmakeRefereeStartRoundParam.PIDs[i])
			if i == len(matchmakeRefereeStartRoundParam.PIDs)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStartRoundParam returns a new MatchmakeRefereeStartRoundParam
func NewMatchmakeRefereeStartRoundParam() *MatchmakeRefereeStartRoundParam {
	return &MatchmakeRefereeStartRoundParam{}
}
