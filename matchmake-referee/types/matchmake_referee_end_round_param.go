// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeEndRoundParam contains the results of a round
type MatchmakeRefereeEndRoundParam struct {
	types.Structure
	*types.Data
	RoundID              *types.PrimitiveU64
	PersonalRoundResults []*MatchmakeRefereePersonalRoundResult
}

// WriteTo writes the MatchmakeRefereeEndRoundParam to the given writable
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeEndRoundParam.RoundID.WriteTo(contentWritable)
	matchmakeRefereeEndRoundParam.PersonalRoundResults.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeEndRoundParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeEndRoundParam from the given readable
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeEndRoundParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeEndRoundParam header. %s", err.Error())
	}

	err = matchmakeRefereeEndRoundParam.RoundID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam.RoundID. %s", err.Error())
	}

	resultList, err := nex.StreamReadListStructure(stream, NewMatchmakeRefereePersonalRoundResult())
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeEndRoundParam.PersonalRoundResults. %s", err.Error())
	}

	matchmakeRefereeEndRoundParam.PersonalRoundResults = resultList

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeEndRoundParam
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeEndRoundParam()

	copied.StructureVersion = matchmakeRefereeEndRoundParam.StructureVersion

	copied.Data = matchmakeRefereeEndRoundParam.Data.Copy().(*types.Data)

	copied.RoundID = matchmakeRefereeEndRoundParam.RoundID

	copied.PersonalRoundResults = make([]*MatchmakeRefereePersonalRoundResult, len(matchmakeRefereeEndRoundParam.PersonalRoundResults))
	for i := 0; i < len(matchmakeRefereeEndRoundParam.PersonalRoundResults); i++ {
		copied.PersonalRoundResults[i] = matchmakeRefereeEndRoundParam.PersonalRoundResults[i].Copy().(*MatchmakeRefereePersonalRoundResult)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeEndRoundParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeEndRoundParam)

	if matchmakeRefereeEndRoundParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeEndRoundParam.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeEndRoundParam.RoundID.Equals(other.RoundID) {
		return false
	}

	if matchmakeRefereeEndRoundParam.PersonalRoundResults != nil && other.PersonalRoundResults != nil {
		if len(matchmakeRefereeEndRoundParam.PersonalRoundResults) != len(other.PersonalRoundResults) {
			return false
		}

		for i := 0; i < len(matchmakeRefereeEndRoundParam.PersonalRoundResults); i++ {
			if !matchmakeRefereeEndRoundParam.PersonalRoundResults[i].Equals(other.PersonalRoundResults[i]) {
				return false
			}
		}
	} else if matchmakeRefereeEndRoundParam.PersonalRoundResults == nil {
		return false
	} else if other.PersonalRoundResults == nil {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) String() string {
	return matchmakeRefereeEndRoundParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeEndRoundParam *MatchmakeRefereeEndRoundParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeEndRoundParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeEndRoundParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sRoundID: %d,\n", indentationValues, matchmakeRefereeEndRoundParam.RoundID))
	if len(matchmakeRefereeEndRoundParam.PersonalRoundResults) == 0 {
		b.WriteString(fmt.Sprintf("%sPersonalRoundResults: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPersonalRoundResults: [\n", indentationValues))

		for i := 0; i < len(matchmakeRefereeEndRoundParam.PersonalRoundResults); i++ {
			str := matchmakeRefereeEndRoundParam.PersonalRoundResults[i].FormatToString(indentationLevel + 2)
			if i == len(matchmakeRefereeEndRoundParam.PersonalRoundResults)-1 {
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

// NewMatchmakeRefereeEndRoundParam returns a new MatchmakeRefereeEndRoundParam
func NewMatchmakeRefereeEndRoundParam() *MatchmakeRefereeEndRoundParam {
	return &MatchmakeRefereeEndRoundParam{}
}
