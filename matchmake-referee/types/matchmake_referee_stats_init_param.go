// Package matchmake_referee_types implements all the types used by the Matchmake Referee protocol
package matchmake_referee_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeRefereeStatsInitParam contains the results of a round
type MatchmakeRefereeStatsInitParam struct {
	nex.Structure
	*nex.Data
	Category           uint32
	InitialRatingValue uint32
}

// Bytes encodes the MatchmakeRefereeStatsInitParam and returns a byte array
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(matchmakeRefereeStatsInitParam.Category)
	stream.WriteUInt32LE(matchmakeRefereeStatsInitParam.InitialRatingValue)

	return stream.Bytes()
}

// ExtractFromStream extracts a MatchmakeRefereeStatsInitParam structure from a stream
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeRefereeStatsInitParam.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.Category. %s", err.Error())
	}

	matchmakeRefereeStatsInitParam.InitialRatingValue, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeStatsInitParam.InitialRatingValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeStatsInitParam
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeRefereeStatsInitParam()

	copied.Data = matchmakeRefereeStatsInitParam.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.Category = matchmakeRefereeStatsInitParam.Category
	copied.InitialRatingValue = matchmakeRefereeStatsInitParam.InitialRatingValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeRefereeStatsInitParam *MatchmakeRefereeStatsInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeRefereeStatsInitParam)

	if !matchmakeRefereeStatsInitParam.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeRefereeStatsInitParam.Category != other.Category {
		return false
	}

	if matchmakeRefereeStatsInitParam.InitialRatingValue != other.InitialRatingValue {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.Category))
	b.WriteString(fmt.Sprintf("%sInitialRatingValue: %d,\n", indentationValues, matchmakeRefereeStatsInitParam.InitialRatingValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeStatsInitParam returns a new MatchmakeRefereeStatsInitParam
func NewMatchmakeRefereeStatsInitParam() *MatchmakeRefereeStatsInitParam {
	return &MatchmakeRefereeStatsInitParam{}
}
