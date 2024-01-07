// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeRound contains the results of a round
type MatchmakeRefereeRound struct {
	types.Structure
	*types.Data
	RoundID                        *types.PrimitiveU64
	GID                            *types.PrimitiveU32
	State                          *types.PrimitiveU32
	PersonalDataCategory           *types.PrimitiveU32
	NormalizedPersonalRoundResults []*MatchmakeRefereePersonalRoundResult
}

// WriteTo writes the MatchmakeRefereeRound to the given writable
func (matchmakeRefereeRound *MatchmakeRefereeRound) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeRound.RoundID.WriteTo(contentWritable)
	matchmakeRefereeRound.GID.WriteTo(contentWritable)
	matchmakeRefereeRound.State.WriteTo(contentWritable)
	matchmakeRefereeRound.PersonalDataCategory.WriteTo(contentWritable)
	matchmakeRefereeRound.NormalizedPersonalRoundResults.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeRound.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeRound from the given readable
func (matchmakeRefereeRound *MatchmakeRefereeRound) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeRound.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeRound header. %s", err.Error())
	}

	err = matchmakeRefereeRound.RoundID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.RoundID. %s", err.Error())
	}

	err = matchmakeRefereeRound.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.GID. %s", err.Error())
	}

	err = matchmakeRefereeRound.State.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.State. %s", err.Error())
	}

	err = matchmakeRefereeRound.PersonalDataCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.PersonalDataCategory. %s", err.Error())
	}

	resultList, err := nex.StreamReadListStructure(stream, NewMatchmakeRefereePersonalRoundResult())
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.NormalizedPersonalRoundResults. %s", err.Error())
	}

	matchmakeRefereeRound.NormalizedPersonalRoundResults = resultList

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeRound
func (matchmakeRefereeRound *MatchmakeRefereeRound) Copy() types.RVType {
	copied := NewMatchmakeRefereeRound()

	copied.StructureVersion = matchmakeRefereeRound.StructureVersion

	copied.Data = matchmakeRefereeRound.Data.Copy().(*types.Data)

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
func (matchmakeRefereeRound *MatchmakeRefereeRound) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeRound); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeRound)

	if matchmakeRefereeRound.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeRound.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeRound.RoundID.Equals(other.RoundID) {
		return false
	}

	if !matchmakeRefereeRound.GID.Equals(other.GID) {
		return false
	}

	if !matchmakeRefereeRound.State.Equals(other.State) {
		return false
	}

	if !matchmakeRefereeRound.PersonalDataCategory.Equals(other.PersonalDataCategory) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeRound.StructureVersion))
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
