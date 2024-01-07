// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// CompetitionRankingInfo holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingInfo struct {
	types.Structure
	Unknown  *types.PrimitiveU32
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the CompetitionRankingInfo from the given readable
func (competitionRankingInfo *CompetitionRankingInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = competitionRankingInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read CompetitionRankingInfo header. %s", err.Error())
	}

	err = competitionRankingInfo.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown from stream. %s", err.Error())
	}

	err = competitionRankingInfo.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown2 from stream. %s", err.Error())
	}

	err = competitionRankingInfo.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown3 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the CompetitionRankingInfo to the given writable
func (competitionRankingInfo *CompetitionRankingInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	competitionRankingInfo.Unknown.WriteTo(contentWritable)
	competitionRankingInfo.Unknown2.WriteTo(contentWritable)
	competitionRankingInfo.Unknown3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	competitionRankingInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of CompetitionRankingInfo
func (competitionRankingInfo *CompetitionRankingInfo) Copy() types.RVType {
	copied := NewCompetitionRankingInfo()

	copied.StructureVersion = competitionRankingInfo.StructureVersion

	copied.Unknown = competitionRankingInfo.Unknown
	copied.Unknown2 = competitionRankingInfo.Unknown2
	copied.Unknown3 = make(*types.List[*types.PrimitiveU32], len(competitionRankingInfo.Unknown3))

	copy(copied.Unknown3, competitionRankingInfo.Unknown3)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (competitionRankingInfo *CompetitionRankingInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*CompetitionRankingInfo); !ok {
		return false
	}

	other := o.(*CompetitionRankingInfo)

	if competitionRankingInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !competitionRankingInfo.Unknown.Equals(other.Unknown) {
		return false
	}

	if !competitionRankingInfo.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if len(competitionRankingInfo.Unknown3) != len(other.Unknown3) {
		return false
	}

	for i := 0; i < len(competitionRankingInfo.Unknown3); i++ {
		if competitionRankingInfo.Unknown3[i] != other.Unknown3[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (competitionRankingInfo *CompetitionRankingInfo) String() string {
	return competitionRankingInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (competitionRankingInfo *CompetitionRankingInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, competitionRankingInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, competitionRankingInfo.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, competitionRankingInfo.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %v,\n", indentationValues, competitionRankingInfo.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingInfo returns a new CompetitionRankingInfo
func NewCompetitionRankingInfo() *CompetitionRankingInfo {
	return &CompetitionRankingInfo{}
}
