// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeRefereeRound contains the results of a round
type MatchmakeRefereeRound struct {
	nex.Structure
	*nex.Data
	RoundID                        uint64
	GID                            uint32
	State                          uint32
	PersonalDataCategory           uint32
	NormalizedPersonalRoundResults []*MatchmakeRefereePersonalRoundResult
}

// Bytes encodes the MatchmakeRefereeRound and returns a byte array
func (matchmakeRefereeRound *MatchmakeRefereeRound) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(matchmakeRefereeRound.RoundID)
	stream.WriteUInt32LE(matchmakeRefereeRound.GID)
	stream.WriteUInt32LE(matchmakeRefereeRound.State)
	stream.WriteUInt32LE(matchmakeRefereeRound.PersonalDataCategory)
	stream.WriteListStructure(matchmakeRefereeRound.NormalizedPersonalRoundResults)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereeRound structure from a stream
func (matchmakeRefereeRound *MatchmakeRefereeRound) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereeRound.RoundID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.RoundID. %s", err.Error())
	}

	matchmakeRefereeRound.GID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.GID. %s", err.Error())
	}

	matchmakeRefereeRound.State, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.State. %s", err.Error())
	}

	matchmakeRefereeRound.PersonalDataCategory, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.PersonalDataCategory. %s", err.Error())
	}

	resultList, err := stream.ReadListStructure(NewMatchmakeRefereePersonalRoundResult())
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.NormalizedPersonalRoundResults. %s", err.Error())
	}

	matchmakeRefereeRound.NormalizedPersonalRoundResults = resultList.([]*MatchmakeRefereePersonalRoundResult)

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeRound
func (matchmakeRefereeRound *MatchmakeRefereeRound) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereeRound()

	copied.SetStructureVersion(matchmakeRefereeRound.StructureVersion())

	copied.Data = matchmakeRefereeRound.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.RoundID = matchmakeRefereeRound.RoundID
	copied.GID = matchmakeRefereeRound.GID
	copied.State = matchmakeRefereeRound.State
	copied.PersonalDataCategory = matchmakeRefereeRound.PersonalDataCategory

	copied.NormalizedPersonalRoundResults = make([]*MatchmakeRefereePersonalRoundResult, len(matchmakeRefereeRound.NormalizedPersonalRoundResults))
	for i := 0; i < len(matchmakeRefereeRound.NormalizedPersonalRoundResults); i++ {
		copied.NormalizedPersonalRoundResults[i] = matchmakeRefereeRound.NormalizedPersonalRoundResults[i].Copy().(*MatchmakeRefereePersonalRoundResult)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeRound *MatchmakeRefereeRound) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereeRound)

	if matchmakeRefereeRound.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !matchmakeRefereeRound.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereeRound.RoundID != other.RoundID {
		return false
	}

	if matchmakeRefereeRound.GID != other.GID {
		return false
	}

	if matchmakeRefereeRound.State != other.State {
		return false
	}

	if matchmakeRefereeRound.PersonalDataCategory != other.PersonalDataCategory {
		return false
	}

	if matchmakeRefereeRound.NormalizedPersonalRoundResults != nil && other.NormalizedPersonalRoundResults != nil {
		if len(matchmakeRefereeRound.NormalizedPersonalRoundResults) != len(other.NormalizedPersonalRoundResults) {
			return false
		}

		for i := 0; i < len(matchmakeRefereeRound.NormalizedPersonalRoundResults); i++ {
			if !matchmakeRefereeRound.NormalizedPersonalRoundResults[i].Equals(other.NormalizedPersonalRoundResults[i]) {
				return false
			}
		}
	} else if matchmakeRefereeRound.NormalizedPersonalRoundResults == nil {
		return false
	} else if other.NormalizedPersonalRoundResults == nil {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeRound *MatchmakeRefereeRound) String() string {
	return matchmakeRefereeRound.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeRound *MatchmakeRefereeRound) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeRound{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereeRound.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sRoundID: %d,\n", indentationValues, matchmakeRefereeRound.RoundID))
	b.WriteString(fmt.Sprintf("%sGID: %d,\n", indentationValues, matchmakeRefereeRound.GID))
	b.WriteString(fmt.Sprintf("%sState: %d,\n", indentationValues, matchmakeRefereeRound.State))
	b.WriteString(fmt.Sprintf("%sPersonalDataCategory: %d,\n", indentationValues, matchmakeRefereeRound.PersonalDataCategory))
	if len(matchmakeRefereeRound.NormalizedPersonalRoundResults) == 0 {
		b.WriteString(fmt.Sprintf("%sPersonalRoundResults: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPersonalRoundResults: [\n", indentationValues))

		for i := 0; i < len(matchmakeRefereeRound.NormalizedPersonalRoundResults); i++ {
			str := matchmakeRefereeRound.NormalizedPersonalRoundResults[i].FormatToString(indentationLevel + 2)
			if i == len(matchmakeRefereeRound.NormalizedPersonalRoundResults)-1 {
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

// NewMatchmakeRefereeRound returns a new MatchmakeRefereeRound
func NewMatchmakeRefereeRound() *MatchmakeRefereeRound {
	return &MatchmakeRefereeRound{}
}
