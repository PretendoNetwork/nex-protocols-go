// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreFightingPowerScore is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreFightingPowerScore struct {
	types.Structure
	Score *types.PrimitiveU32
	Rank  *types.PrimitiveU32
}

// ExtractFrom extracts the DataStoreFightingPowerScore from the given readable
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreFightingPowerScore.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreFightingPowerScore header. %s", err.Error())
	}

	err = dataStoreFightingPowerScore.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Score. %s", err.Error())
	}

	err = dataStoreFightingPowerScore.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Rank. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreFightingPowerScore to the given writable
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreFightingPowerScore.Score.WriteTo(contentWritable)
	dataStoreFightingPowerScore.Rank.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreFightingPowerScore.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreFightingPowerScore
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Copy() types.RVType {
	copied := NewDataStoreFightingPowerScore()

	copied.StructureVersion = dataStoreFightingPowerScore.StructureVersion

	copied.Score = dataStoreFightingPowerScore.Score.Copy().(*types.PrimitiveU32)
	copied.Rank = dataStoreFightingPowerScore.Rank.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFightingPowerScore); !ok {
		return false
	}

	other := o.(*DataStoreFightingPowerScore)

	if dataStoreFightingPowerScore.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreFightingPowerScore.Score.Equals(other.Score) {
		return false
	}

	if !dataStoreFightingPowerScore.Rank.Equals(other.Rank) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) String() string {
	return dataStoreFightingPowerScore.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFightingPowerScore{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreFightingPowerScore.StructureVersion))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStoreFightingPowerScore.Score))
	b.WriteString(fmt.Sprintf("%sRank: %s\n", indentationValues, dataStoreFightingPowerScore.Rank))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerScore returns a new DataStoreFightingPowerScore
func NewDataStoreFightingPowerScore() *DataStoreFightingPowerScore {
	return &DataStoreFightingPowerScore{
		Score: types.NewPrimitiveU32(0),
		Rank: types.NewPrimitiveU32(0),
	}
}
