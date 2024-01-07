// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2EstimateScoreRankOutput holds data for the Ranking 2  protocol
type Ranking2EstimateScoreRankOutput struct {
	types.Structure
	Rank         *types.PrimitiveU32
	Length       *types.PrimitiveU32
	Score        *types.PrimitiveU32
	Category     *types.PrimitiveU32
	Season       *types.PrimitiveS32
	SamplingRate *types.PrimitiveU8
}

// ExtractFrom extracts the Ranking2EstimateScoreRankOutput from the given readable
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2EstimateScoreRankOutput.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2EstimateScoreRankOutput header. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Rank from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Length from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Score from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Category from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Season from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankOutput.SamplingRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.SamplingRate from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2EstimateScoreRankOutput to the given writable
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2EstimateScoreRankOutput.Rank.WriteTo(contentWritable)
	ranking2EstimateScoreRankOutput.Length.WriteTo(contentWritable)
	ranking2EstimateScoreRankOutput.Score.WriteTo(contentWritable)
	ranking2EstimateScoreRankOutput.Category.WriteTo(contentWritable)
	ranking2EstimateScoreRankOutput.Season.WriteTo(contentWritable)
	ranking2EstimateScoreRankOutput.SamplingRate.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2EstimateScoreRankOutput.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankOutput
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) Copy() types.RVType {
	copied := NewRanking2EstimateScoreRankOutput()

	copied.StructureVersion = ranking2EstimateScoreRankOutput.StructureVersion

	copied.Rank = ranking2EstimateScoreRankOutput.Rank
	copied.Length = ranking2EstimateScoreRankOutput.Length
	copied.Score = ranking2EstimateScoreRankOutput.Score
	copied.Category = ranking2EstimateScoreRankOutput.Category
	copied.Season = ranking2EstimateScoreRankOutput.Season
	copied.SamplingRate = ranking2EstimateScoreRankOutput.SamplingRate
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2EstimateScoreRankOutput); !ok {
		return false
	}

	other := o.(*Ranking2EstimateScoreRankOutput)

	if ranking2EstimateScoreRankOutput.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2EstimateScoreRankOutput.Rank.Equals(other.Rank) {
		return false
	}

	if !ranking2EstimateScoreRankOutput.Length.Equals(other.Length) {
		return false
	}

	if !ranking2EstimateScoreRankOutput.Score.Equals(other.Score) {
		return false
	}

	if !ranking2EstimateScoreRankOutput.Category.Equals(other.Category) {
		return false
	}

	if !ranking2EstimateScoreRankOutput.Season.Equals(other.Season) {
		return false
	}

	if !ranking2EstimateScoreRankOutput.SamplingRate.Equals(other.SamplingRate) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) String() string {
	return ranking2EstimateScoreRankOutput.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2EstimateScoreRankOutput{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.StructureVersion))
	b.WriteString(fmt.Sprintf("%sRank: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.Rank))
	b.WriteString(fmt.Sprintf("%sLength: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.Length))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.Score))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.Category))
	b.WriteString(fmt.Sprintf("%sSeason: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.Season))
	b.WriteString(fmt.Sprintf("%sSamplingRate: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.SamplingRate))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2EstimateScoreRankOutput returns a new Ranking2EstimateScoreRankOutput
func NewRanking2EstimateScoreRankOutput() *Ranking2EstimateScoreRankOutput {
	return &Ranking2EstimateScoreRankOutput{}
}
