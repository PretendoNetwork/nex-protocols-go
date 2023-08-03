// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// CompetitionRankingInfo holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingInfo struct {
	nex.Structure
	Unknown uint32
	Unknown2 uint32
	Unknown3 []uint32
}

// ExtractFromStream extracts a CompetitionRankingInfo structure from a stream
func (competitionRankingInfo *CompetitionRankingInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	competitionRankingInfo.Unknown, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown from stream. %s", err.Error())
	}

	competitionRankingInfo.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown2 from stream. %s", err.Error())
	}

	competitionRankingInfo.Unknown3, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfo.Unknown3 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the CompetitionRankingInfo and returns a byte array
func (competitionRankingInfo *CompetitionRankingInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(competitionRankingInfo.Unknown)
	stream.WriteUInt32LE(competitionRankingInfo.Unknown2)
	stream.WriteListUInt32LE(competitionRankingInfo.Unknown3)

	return stream.Bytes()
}

// Copy returns a new copied instance of CompetitionRankingInfo
func (competitionRankingInfo *CompetitionRankingInfo) Copy() nex.StructureInterface {
	copied := NewCompetitionRankingInfo()

	copied.Unknown = competitionRankingInfo.Unknown
	copied.Unknown2 = competitionRankingInfo.Unknown2
	copied.Unknown3 = make([]uint32, len(competitionRankingInfo.Unknown3))

	copy(copied.Unknown3, competitionRankingInfo.Unknown3)


	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (competitionRankingInfo *CompetitionRankingInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*CompetitionRankingInfo)

	if competitionRankingInfo.Unknown != other.Unknown {
		return false
	}

	if competitionRankingInfo.Unknown2 != other.Unknown2 {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, competitionRankingInfo.StructureVersion()))
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
