// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2EstimateScoreRankOutput is a type within the Ranking2 protocol
type Ranking2EstimateScoreRankOutput struct {
	types.Structure
	Rank         types.UInt32
	Length       types.UInt32
	Score        types.UInt32
	Category     types.UInt32
	Season       types.Int32
	SamplingRate types.UInt8
}

// WriteTo writes the Ranking2EstimateScoreRankOutput to the given writable
func (resro Ranking2EstimateScoreRankOutput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	resro.Rank.WriteTo(contentWritable)
	resro.Length.WriteTo(contentWritable)
	resro.Score.WriteTo(contentWritable)
	resro.Category.WriteTo(contentWritable)
	resro.Season.WriteTo(contentWritable)
	resro.SamplingRate.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	resro.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2EstimateScoreRankOutput from the given readable
func (resro *Ranking2EstimateScoreRankOutput) ExtractFrom(readable types.Readable) error {
	var err error

	err = resro.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput header. %s", err.Error())
	}

	err = resro.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Rank. %s", err.Error())
	}

	err = resro.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Length. %s", err.Error())
	}

	err = resro.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Score. %s", err.Error())
	}

	err = resro.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Category. %s", err.Error())
	}

	err = resro.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Season. %s", err.Error())
	}

	err = resro.SamplingRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.SamplingRate. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankOutput
func (resro Ranking2EstimateScoreRankOutput) Copy() types.RVType {
	copied := NewRanking2EstimateScoreRankOutput()

	copied.StructureVersion = resro.StructureVersion
	copied.Rank = resro.Rank.Copy().(types.UInt32)
	copied.Length = resro.Length.Copy().(types.UInt32)
	copied.Score = resro.Score.Copy().(types.UInt32)
	copied.Category = resro.Category.Copy().(types.UInt32)
	copied.Season = resro.Season.Copy().(types.Int32)
	copied.SamplingRate = resro.SamplingRate.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given Ranking2EstimateScoreRankOutput contains the same data as the current Ranking2EstimateScoreRankOutput
func (resro Ranking2EstimateScoreRankOutput) Equals(o types.RVType) bool {
	if _, ok := o.(Ranking2EstimateScoreRankOutput); !ok {
		return false
	}

	other := o.(Ranking2EstimateScoreRankOutput)

	if resro.StructureVersion != other.StructureVersion {
		return false
	}

	if !resro.Rank.Equals(other.Rank) {
		return false
	}

	if !resro.Length.Equals(other.Length) {
		return false
	}

	if !resro.Score.Equals(other.Score) {
		return false
	}

	if !resro.Category.Equals(other.Category) {
		return false
	}

	if !resro.Season.Equals(other.Season) {
		return false
	}

	return resro.SamplingRate.Equals(other.SamplingRate)
}

// CopyRef copies the current value of the Ranking2EstimateScoreRankOutput
// and returns a pointer to the new copy
func (resro Ranking2EstimateScoreRankOutput) CopyRef() types.RVTypePtr {
	copied := resro.Copy().(Ranking2EstimateScoreRankOutput)
	return &copied
}

// Deref takes a pointer to the Ranking2EstimateScoreRankOutput
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (resro *Ranking2EstimateScoreRankOutput) Deref() types.RVType {
	return *resro
}

// String returns the string representation of the Ranking2EstimateScoreRankOutput
func (resro Ranking2EstimateScoreRankOutput) String() string {
	return resro.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2EstimateScoreRankOutput using the provided indentation level
func (resro Ranking2EstimateScoreRankOutput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2EstimateScoreRankOutput{\n")
	b.WriteString(fmt.Sprintf("%sRank: %s,\n", indentationValues, resro.Rank))
	b.WriteString(fmt.Sprintf("%sLength: %s,\n", indentationValues, resro.Length))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, resro.Score))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, resro.Category))
	b.WriteString(fmt.Sprintf("%sSeason: %s,\n", indentationValues, resro.Season))
	b.WriteString(fmt.Sprintf("%sSamplingRate: %s,\n", indentationValues, resro.SamplingRate))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2EstimateScoreRankOutput returns a new Ranking2EstimateScoreRankOutput
func NewRanking2EstimateScoreRankOutput() Ranking2EstimateScoreRankOutput {
	return Ranking2EstimateScoreRankOutput{
		Rank:         types.NewUInt32(0),
		Length:       types.NewUInt32(0),
		Score:        types.NewUInt32(0),
		Category:     types.NewUInt32(0),
		Season:       types.NewInt32(0),
		SamplingRate: types.NewUInt8(0),
	}

}
