// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingStats holds parameters for ordering rankings
type RankingStats struct {
	nex.Structure
	StatsList []float64
}

// ExtractFromStream extracts a RankingStats structure from a stream
func (rankingStats *RankingStats) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingStats.StatsList, err = stream.ReadListFloat64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingStats.StatsList from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingStats and returns a byte array
func (rankingStats *RankingStats) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListFloat64LE(rankingStats.StatsList)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingStats
func (rankingStats *RankingStats) Copy() nex.StructureInterface {
	copied := NewRankingStats()

	copied.SetStructureVersion(rankingStats.StructureVersion())

	copied.StatsList = make([]float64, len(rankingStats.StatsList))

	copy(copied.StatsList, rankingStats.StatsList)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingStats *RankingStats) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingStats)

	if rankingStats.StructureVersion() != other.StructureVersion() {
		return false
	}

	if len(rankingStats.StatsList) != len(other.StatsList) {
		return false
	}

	for i := 0; i < len(rankingStats.StatsList); i++ {
		if rankingStats.StatsList[i] != other.StatsList[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (rankingStats *RankingStats) String() string {
	return rankingStats.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingStats *RankingStats) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingStats{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingStats.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sStatsList: %v\n", indentationValues, rankingStats.StatsList))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingStats returns a new RankingStats
func NewRankingStats() *RankingStats {
	return &RankingStats{}
}
