// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingStats holds parameters for ordering rankings
type RankingStats struct {
	types.Structure
	StatsList *types.List[*types.PrimitiveF64]
}

// ExtractFrom extracts the RankingStats from the given readable
func (rankingStats *RankingStats) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingStats.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingStats header. %s", err.Error())
	}

	err = rankingStats.StatsList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingStats.StatsList from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingStats to the given writable
func (rankingStats *RankingStats) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingStats.StatsList.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingStats.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingStats
func (rankingStats *RankingStats) Copy() types.RVType {
	copied := NewRankingStats()

	copied.StructureVersion = rankingStats.StructureVersion

	copied.StatsList = make(*types.List[*types.PrimitiveF64], len(rankingStats.StatsList))

	copy(copied.StatsList, rankingStats.StatsList)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingStats *RankingStats) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingStats); !ok {
		return false
	}

	other := o.(*RankingStats)

	if rankingStats.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingStats.StructureVersion))
	b.WriteString(fmt.Sprintf("%sStatsList: %v\n", indentationValues, rankingStats.StatsList))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingStats returns a new RankingStats
func NewRankingStats() *RankingStats {
	return &RankingStats{}
}
