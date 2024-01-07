// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2EstimateScoreRankInput holds data for the Ranking 2  protocol
type Ranking2EstimateScoreRankInput struct {
	types.Structure
	Category           *types.PrimitiveU32
	NumSeasonsToGoBack *types.PrimitiveU8
	Score              *types.PrimitiveU32
}

// ExtractFrom extracts the Ranking2EstimateScoreRankInput from the given readable
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2EstimateScoreRankInput.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2EstimateScoreRankInput header. %s", err.Error())
	}

	err = ranking2EstimateScoreRankInput.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Category from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankInput.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	err = ranking2EstimateScoreRankInput.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Score from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2EstimateScoreRankInput to the given writable
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2EstimateScoreRankInput.Category.WriteTo(contentWritable)
	ranking2EstimateScoreRankInput.NumSeasonsToGoBack.WriteTo(contentWritable)
	ranking2EstimateScoreRankInput.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2EstimateScoreRankInput.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankInput
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) Copy() types.RVType {
	copied := NewRanking2EstimateScoreRankInput()

	copied.StructureVersion = ranking2EstimateScoreRankInput.StructureVersion

	copied.Category = ranking2EstimateScoreRankInput.Category
	copied.NumSeasonsToGoBack = ranking2EstimateScoreRankInput.NumSeasonsToGoBack
	copied.Score = ranking2EstimateScoreRankInput.Score
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2EstimateScoreRankInput); !ok {
		return false
	}

	other := o.(*Ranking2EstimateScoreRankInput)

	if ranking2EstimateScoreRankInput.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2EstimateScoreRankInput.Category.Equals(other.Category) {
		return false
	}

	if !ranking2EstimateScoreRankInput.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack) {
		return false
	}

	if !ranking2EstimateScoreRankInput.Score.Equals(other.Score) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) String() string {
	return ranking2EstimateScoreRankInput.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2EstimateScoreRankInput{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2EstimateScoreRankInput.StructureVersion))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2EstimateScoreRankInput.Category))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %d,\n", indentationValues, ranking2EstimateScoreRankInput.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2EstimateScoreRankInput.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2EstimateScoreRankInput returns a new Ranking2EstimateScoreRankInput
func NewRanking2EstimateScoreRankInput() *Ranking2EstimateScoreRankInput {
	return &Ranking2EstimateScoreRankInput{}
}
