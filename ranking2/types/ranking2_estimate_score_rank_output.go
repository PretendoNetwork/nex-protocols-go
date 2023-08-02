// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2EstimateScoreRankOutput holds data for the Ranking 2  protocol
type Ranking2EstimateScoreRankOutput struct {
	nex.Structure
	Rank         uint32
	Length       uint32
	Score        uint32
	Category     uint32
	Season       int32
	SamplingRate uint8
}

// ExtractFromStream extracts a Ranking2EstimateScoreRankOutput structure from a stream
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2EstimateScoreRankOutput.Rank, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Rank from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankOutput.Length, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Length from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankOutput.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Score from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankOutput.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Category from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankOutput.Season, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.Season from stream. %s", err.Error())
	}

	ranking2EstimateScoreRankOutput.SamplingRate, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2EstimateScoreRankOutput.SamplingRate from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2EstimateScoreRankOutput and returns a byte array
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(ranking2EstimateScoreRankOutput.Rank)
	stream.WriteUInt32LE(ranking2EstimateScoreRankOutput.Length)
	stream.WriteUInt32LE(ranking2EstimateScoreRankOutput.Score)
	stream.WriteUInt32LE(ranking2EstimateScoreRankOutput.Category)
	stream.WriteInt32LE(ranking2EstimateScoreRankOutput.Season)
	stream.WriteUInt8(ranking2EstimateScoreRankOutput.SamplingRate)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2EstimateScoreRankOutput
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) Copy() nex.StructureInterface {
	copied := NewRanking2EstimateScoreRankOutput()

	copied.Rank = ranking2EstimateScoreRankOutput.Rank
	copied.Length = ranking2EstimateScoreRankOutput.Length
	copied.Score = ranking2EstimateScoreRankOutput.Score
	copied.Category = ranking2EstimateScoreRankOutput.Category
	copied.Season = ranking2EstimateScoreRankOutput.Season
	copied.SamplingRate = ranking2EstimateScoreRankOutput.SamplingRate
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2EstimateScoreRankOutput *Ranking2EstimateScoreRankOutput) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2EstimateScoreRankOutput)

	if ranking2EstimateScoreRankOutput.Rank != other.Rank {
		return false
	}

	if ranking2EstimateScoreRankOutput.Length != other.Length {
		return false
	}

	if ranking2EstimateScoreRankOutput.Score != other.Score {
		return false
	}

	if ranking2EstimateScoreRankOutput.Category != other.Category {
		return false
	}

	if ranking2EstimateScoreRankOutput.Season != other.Season {
		return false
	}

	if ranking2EstimateScoreRankOutput.SamplingRate != other.SamplingRate {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2EstimateScoreRankOutput.StructureVersion()))
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
