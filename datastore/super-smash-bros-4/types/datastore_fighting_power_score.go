package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreFightingPowerScore struct {
	nex.Structure
	Score uint32
	Rank  uint32
}

// ExtractFromStream extracts a DataStoreFightingPowerScore structure from a stream
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreFightingPowerScore.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Score. %s", err.Error())
	}

	dataStoreFightingPowerScore.Rank, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerScore.Rank. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreFightingPowerScore and returns a byte array
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreFightingPowerScore.Score)
	stream.WriteUInt32LE(dataStoreFightingPowerScore.Rank)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFightingPowerScore
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Copy() nex.StructureInterface {
	copied := NewDataStoreFightingPowerScore()

	copied.Score = dataStoreFightingPowerScore.Score
	copied.Rank = dataStoreFightingPowerScore.Rank

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFightingPowerScore)

	if dataStoreFightingPowerScore.Score != other.Score {
		return false
	}

	if dataStoreFightingPowerScore.Rank != other.Rank {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreFightingPowerScore.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, dataStoreFightingPowerScore.Score))
	b.WriteString(fmt.Sprintf("%sRank: %d\n", indentationValues, dataStoreFightingPowerScore.Rank))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerScore returns a new DataStoreFightingPowerScore
func NewDataStoreFightingPowerScore() *DataStoreFightingPowerScore {
	return &DataStoreFightingPowerScore{}
}
