// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// CompetitionRankingInfoGetParam holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingInfoGetParam struct {
	nex.Structure
	Unknown uint8
	Result  *nex.ResultRange
}

// ExtractFromStream extracts a CompetitionRankingInfoGetParam structure from a stream
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	competitionRankingInfoGetParam.Unknown, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Unknown from stream. %s", err.Error())
	}

	competitionRankingInfoGetParam.Result, err = nex.StreamReadStructure(stream, nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Result from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the CompetitionRankingInfoGetParam and returns a byte array
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(competitionRankingInfoGetParam.Unknown)
	stream.WriteStructure(competitionRankingInfoGetParam.Result)

	return stream.Bytes()
}

// Copy returns a new copied instance of CompetitionRankingInfoGetParam
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) Copy() nex.StructureInterface {
	copied := NewCompetitionRankingInfoGetParam()

	copied.SetStructureVersion(competitionRankingInfoGetParam.StructureVersion())

	copied.Unknown = competitionRankingInfoGetParam.Unknown
	copied.Result = competitionRankingInfoGetParam.Result.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*CompetitionRankingInfoGetParam)

	if competitionRankingInfoGetParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if competitionRankingInfoGetParam.Unknown != other.Unknown {
		return false
	}

	if !competitionRankingInfoGetParam.Result.Equals(other.Result) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) String() string {
	return competitionRankingInfoGetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingInfoGetParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, competitionRankingInfoGetParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, competitionRankingInfoGetParam.Unknown))

	if competitionRankingInfoGetParam.Result != nil {
		b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, competitionRankingInfoGetParam.Result.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingInfoGetParam returns a new CompetitionRankingInfoGetParam
func NewCompetitionRankingInfoGetParam() *CompetitionRankingInfoGetParam {
	return &CompetitionRankingInfoGetParam{}
}
