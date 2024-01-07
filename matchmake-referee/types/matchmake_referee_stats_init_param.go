// Package types implements all the types used by the Matchmake Referee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MatchmakeRefereeStatsInitParam contains the results of a round
type MatchmakeRefereeStatsInitParam struct {
	types.Structure
	*types.Data
	Category           *types.PrimitiveU32
	InitialRatingValue *types.PrimitiveU32
}

// WriteTo writes the MatchmakeRefereeStatsInitParam to the given writable
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakeRefereeStatsInitParam.Category.WriteTo(contentWritable)
	matchmakeRefereeStatsInitParam.InitialRatingValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	matchmakeRefereeStatsInitParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeStatsInitParam from the given readable
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = matchmakeRefereeStatsInitParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MatchmakeRefereeStatsInitParam header. %s", err.Error())
	}

	err = matchmakeRefereeStatsInitParam.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.Category. %s", err.Error())
	}

	err = matchmakeRefereeStatsInitParam.InitialRatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.InitialRatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsInitParam
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) Copy() types.RVType {
	copied := NewMatchmakeRefereeStatsInitParam()

	copied.StructureVersion = matchmakeRefereeStatsInitParam.StructureVersion

	copied.Data = matchmakeRefereeStatsInitParam.Data.Copy().(*types.Data)

	copied.Category = matchmakeRefereeStatsInitParam.Category
	copied.InitialRatingValue = matchmakeRefereeStatsInitParam.InitialRatingValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeRefereeStatsInitParam); !ok {
		return false
	}

	other := o.(*MatchmakeRefereeStatsInitParam)

	if matchmakeRefereeStatsInitParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeRefereeStatsInitParam.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeRefereeStatsInitParam.Category.Equals(other.Category) {
		return false
	}

	if !matchmakeRefereeStatsInitParam.InitialRatingValue.Equals(other.InitialRatingValue) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) String() string {
	return matchmakeRefereeStatsInitParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeStatsInitParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.Category))
	b.WriteString(fmt.Sprintf("%sInitialRatingValue: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.InitialRatingValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsInitParam returns a new MatchmakeRefereeStatsInitParam
func NewMatchmakeRefereeStatsInitParam() *MatchmakeRefereeStatsInitParam {
	return &MatchmakeRefereeStatsInitParam{}
}
