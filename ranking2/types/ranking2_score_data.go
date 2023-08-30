// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2ScoreData holds data for the Ranking 2  protocol
type Ranking2ScoreData struct {
	nex.Structure
	Misc     uint64
	Category uint32
	Score    uint32
}

// ExtractFromStream extracts a Ranking2ScoreData structure from a stream
func (ranking2ScoreData *Ranking2ScoreData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2ScoreData.Misc, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Misc from stream. %s", err.Error())
	}

	ranking2ScoreData.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Category from stream. %s", err.Error())
	}

	ranking2ScoreData.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Score from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2ScoreData and returns a byte array
func (ranking2ScoreData *Ranking2ScoreData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(ranking2ScoreData.Misc)
	stream.WriteUInt32LE(ranking2ScoreData.Category)
	stream.WriteUInt32LE(ranking2ScoreData.Score)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2ScoreData
func (ranking2ScoreData *Ranking2ScoreData) Copy() nex.StructureInterface {
	copied := NewRanking2ScoreData()

	copied.SetStructureVersion(ranking2ScoreData.StructureVersion())

	copied.Misc = ranking2ScoreData.Misc
	copied.Category = ranking2ScoreData.Category
	copied.Score = ranking2ScoreData.Score
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ScoreData *Ranking2ScoreData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2ScoreData)

	if ranking2ScoreData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if ranking2ScoreData.Misc != other.Misc {
		return false
	}

	if ranking2ScoreData.Category != other.Category {
		return false
	}

	if ranking2ScoreData.Score != other.Score {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2ScoreData *Ranking2ScoreData) String() string {
	return ranking2ScoreData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2ScoreData *Ranking2ScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ScoreData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2ScoreData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sMisc: %d,\n", indentationValues, ranking2ScoreData.Misc))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2ScoreData.Category))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2ScoreData.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ScoreData returns a new Ranking2ScoreData
func NewRanking2ScoreData() *Ranking2ScoreData {
	return &Ranking2ScoreData{}
}
