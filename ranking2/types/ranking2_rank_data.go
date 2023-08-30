// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2RankData holds data for the Ranking 2  protocol
type Ranking2RankData struct {
	nex.Structure
	Misc        uint64
	NexUniqueID uint64
	PrincipalID uint32
	Rank        uint32
	Score       uint32
	CommonData  *Ranking2CommonData
}

// ExtractFromStream extracts a Ranking2RankData structure from a stream
func (ranking2RankData *Ranking2RankData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2RankData.Misc, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Misc from stream. %s", err.Error())
	}

	ranking2RankData.NexUniqueID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.NexUniqueID from stream. %s", err.Error())
	}

	ranking2RankData.PrincipalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.PrincipalID from stream. %s", err.Error())
	}

	ranking2RankData.Rank, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Rank from stream. %s", err.Error())
	}

	ranking2RankData.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Score from stream. %s", err.Error())
	}

	commonData, err := stream.ReadStructure(NewRanking2CommonData())
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.CommonData from stream. %s", err.Error())
	}

	ranking2RankData.CommonData = commonData.(*Ranking2CommonData)

	return nil
}

// Bytes encodes the Ranking2RankData and returns a byte array
func (ranking2RankData *Ranking2RankData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(ranking2RankData.Misc)
	stream.WriteUInt64LE(ranking2RankData.NexUniqueID)
	stream.WriteUInt32LE(ranking2RankData.PrincipalID)
	stream.WriteUInt32LE(ranking2RankData.Rank)
	stream.WriteUInt32LE(ranking2RankData.Score)
	stream.WriteStructure(ranking2RankData.CommonData)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2RankData
func (ranking2RankData *Ranking2RankData) Copy() nex.StructureInterface {
	copied := NewRanking2RankData()

	copied.SetStructureVersion(ranking2RankData.StructureVersion())

	copied.Misc = ranking2RankData.Misc
	copied.NexUniqueID = ranking2RankData.NexUniqueID
	copied.PrincipalID = ranking2RankData.PrincipalID
	copied.Rank = ranking2RankData.Rank
	copied.Score = ranking2RankData.Score
	copied.CommonData = ranking2RankData.CommonData.Copy().(*Ranking2CommonData)
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2RankData *Ranking2RankData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2RankData)

	if ranking2RankData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if ranking2RankData.Misc != other.Misc {
		return false
	}

	if ranking2RankData.NexUniqueID != other.NexUniqueID {
		return false
	}

	if ranking2RankData.PrincipalID != other.PrincipalID {
		return false
	}

	if ranking2RankData.Rank != other.Rank {
		return false
	}

	if ranking2RankData.Score != other.Score {
		return false
	}

	if !ranking2RankData.CommonData.Equals(other.CommonData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2RankData *Ranking2RankData) String() string {
	return ranking2RankData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2RankData *Ranking2RankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2RankData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2RankData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sMisc: %d,\n", indentationValues, ranking2RankData.Misc))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %d,\n", indentationValues, ranking2RankData.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %d,\n", indentationValues, ranking2RankData.PrincipalID))
	b.WriteString(fmt.Sprintf("%sRank: %d,\n", indentationValues, ranking2RankData.Rank))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2RankData.Score))

	if ranking2RankData.CommonData != nil {
		b.WriteString(fmt.Sprintf("%sCommonData: %s\n", indentationValues, ranking2RankData.CommonData.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCommonData: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2RankData returns a new Ranking2RankData
func NewRanking2RankData() *Ranking2RankData {
	return &Ranking2RankData{}
}
