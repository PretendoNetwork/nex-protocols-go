// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// CompetitionRankingInfoGetParam holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingInfoGetParam struct {
	types.Structure
	Unknown *types.PrimitiveU8
	Result  *types.ResultRange
}

// ExtractFrom extracts the CompetitionRankingInfoGetParam from the given readable
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = competitionRankingInfoGetParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read CompetitionRankingInfoGetParam header. %s", err.Error())
	}

	err = competitionRankingInfoGetParam.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Unknown from stream. %s", err.Error())
	}

	err = competitionRankingInfoGetParam.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Result from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the CompetitionRankingInfoGetParam to the given writable
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	competitionRankingInfoGetParam.Unknown.WriteTo(contentWritable)
	competitionRankingInfoGetParam.Result.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	competitionRankingInfoGetParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of CompetitionRankingInfoGetParam
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) Copy() types.RVType {
	copied := NewCompetitionRankingInfoGetParam()

	copied.StructureVersion = competitionRankingInfoGetParam.StructureVersion

	copied.Unknown = competitionRankingInfoGetParam.Unknown
	copied.Result = competitionRankingInfoGetParam.Result.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (competitionRankingInfoGetParam *CompetitionRankingInfoGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*CompetitionRankingInfoGetParam); !ok {
		return false
	}

	other := o.(*CompetitionRankingInfoGetParam)

	if competitionRankingInfoGetParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !competitionRankingInfoGetParam.Unknown.Equals(other.Unknown) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, competitionRankingInfoGetParam.StructureVersion))
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
