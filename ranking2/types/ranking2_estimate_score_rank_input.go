// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2EstimateScoreRankInput holds data for the Ranking 2  protocol
type Ranking2EstimateScoreRankInput struct {
	nex.Structure
	Category           uint32
	NumSeasonsToGoBack uint8
	Score              uint32
}

// ExtractFromStream extracts a Ranking2EstimateScoreRankInput structure from a stream
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2EstimateScoreRankInput.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Category from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankInput.NumSeasonsToGoBack, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankInput.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankInput.Score from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2EstimateScoreRankInput and returns a byte array
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(ranking2EstimateScoreRankInput.Category)
	stream.WriteUInt8(ranking2EstimateScoreRankInput.NumSeasonsToGoBack)
	stream.WriteUInt32LE(ranking2EstimateScoreRankInput.Score)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankInput
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) Copy() nex.StructureInterface {
	copied := NewRanking2EstimateScoreRankInput()

	copied.Category = ranking2EstimateScoreRankInput.Category
	copied.NumSeasonsToGoBack = ranking2EstimateScoreRankInput.NumSeasonsToGoBack
	copied.Score = ranking2EstimateScoreRankInput.Score
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2EstimateScoreRankInput *Ranking2EstimateScoreRankInput) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2EstimateScoreRankInput)

	if ranking2EstimateScoreRankInput.Category != other.Category {
		return false
	}

	if ranking2EstimateScoreRankInput.NumSeasonsToGoBack != other.NumSeasonsToGoBack {
		return false
	}

	if ranking2EstimateScoreRankInput.Score != other.Score {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2EstimateScoreRankInput.StructureVersion()))
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
