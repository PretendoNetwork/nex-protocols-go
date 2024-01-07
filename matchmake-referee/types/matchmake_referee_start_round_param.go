// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeStartRoundParam contains the results of a round
type MatchmakeRefereeStartRoundParam struct {
	types.Structure
	*types.Data
	PersonalDataCategory *types.PrimitiveU32
	GID                  *types.PrimitiveU32
	PIDs                 *types.List[*types.PID]
}

// WriteTo writes the MatchmakeRefereeStartRoundParam to the given writable
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeStartRoundParam.PersonalDataCategory.WriteTo(contentWritable)
	matchmakeRefereeStartRoundParam.GID.WriteTo(contentWritable)
	matchmakeRefereeStartRoundParam.PIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeStartRoundParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStartRoundParam from the given readable
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeStartRoundParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeStartRoundParam header. %s", err.Error())
	}

	err = matchmakeRefereeStartRoundParam.PersonalDataCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PersonalDataCategory. %s", err.Error())
	}

	err = matchmakeRefereeStartRoundParam.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.GID. %s", err.Error())
	}

	err = matchmakeRefereeStartRoundParam.PIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStartRoundParam.PIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStartRoundParam
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeStartRoundParam()

	copied.StructureVersion = matchmakeRefereeStartRoundParam.StructureVersion

	copied.Data = matchmakeRefereeStartRoundParam.Data.Copy().(*types.Data)

	copied.PersonalDataCategory = matchmakeRefereeStartRoundParam.PersonalDataCategory
	copied.GID = matchmakeRefereeStartRoundParam.GID
	copied.PIDs = make(*types.List[*types.PID], len(matchmakeRefereeStartRoundParam.PIDs))

	for i := 0; i < len(matchmakeRefereeStartRoundParam.PIDs); i++ {
		copied.PIDs[i] = matchmakeRefereeStartRoundParam.PIDs[i].Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStartRoundParam *MatchmakeRefereeStartRoundParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStartRoundParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStartRoundParam)

	if matchmakeRefereeStartRoundParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeStartRoundParam.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeStartRoundParam.PersonalDataCategory.Equals(other.PersonalDataCategory) {
		return false
	}

	if !matchmakeRefereeStartRoundParam.GID.Equals(other.GID) {
		return false
	}

	if len(matchmakeRefereeStartRoundParam.PIDs) != len(other.PIDs) {
		return false
	}

	for i := 0; i < len(matchmakeRefereeStartRoundParam.PIDs); i++ {
		if !matchmakeRefereeStartRoundParam.PIDs[i].Equals(other.PIDs[i]) {
			return false
		}
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeStartRoundParam.StructureVersion))
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
