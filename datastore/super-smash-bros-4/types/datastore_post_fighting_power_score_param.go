// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePostFightingPowerScoreParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePostFightingPowerScoreParam struct {
	nex.Structure
	Mode             uint8
	Score            uint32
	IsWorldHighScore bool
}

// ExtractFromStream extracts a DataStorePostFightingPowerScoreParam structure from a stream
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePostFightingPowerScoreParam.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Mode. %s", err.Error())
	}

	dataStorePostFightingPowerScoreParam.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Score. %s", err.Error())
	}

	dataStorePostFightingPowerScoreParam.IsWorldHighScore, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.IsWorldHighScore. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePostFightingPowerScoreParam and returns a byte array
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePostFightingPowerScoreParam.Mode)
	stream.WriteUInt32LE(dataStorePostFightingPowerScoreParam.Score)
	stream.WriteBool(dataStorePostFightingPowerScoreParam.IsWorldHighScore)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePostFightingPowerScoreParam
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Copy() nex.StructureInterface {
	copied := NewDataStorePostFightingPowerScoreParam()

	copied.SetStructureVersion(dataStorePostFightingPowerScoreParam.StructureVersion())

	copied.Mode = dataStorePostFightingPowerScoreParam.Mode
	copied.Score = dataStorePostFightingPowerScoreParam.Score
	copied.IsWorldHighScore = dataStorePostFightingPowerScoreParam.IsWorldHighScore

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePostFightingPowerScoreParam)

	if dataStorePostFightingPowerScoreParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStorePostFightingPowerScoreParam.Mode != other.Mode {
		return false
	}

	if dataStorePostFightingPowerScoreParam.Score != other.Score {
		return false
	}

	if dataStorePostFightingPowerScoreParam.IsWorldHighScore != other.IsWorldHighScore {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) String() string {
	return dataStorePostFightingPowerScoreParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostFightingPowerScoreParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePostFightingPowerScoreParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sMode: %d,\n", indentationValues, dataStorePostFightingPowerScoreParam.Mode))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, dataStorePostFightingPowerScoreParam.Score))
	b.WriteString(fmt.Sprintf("%sIsWorldHighScore: %t\n", indentationValues, dataStorePostFightingPowerScoreParam.IsWorldHighScore))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostFightingPowerScoreParam returns a new DataStorePostFightingPowerScoreParam
func NewDataStorePostFightingPowerScoreParam() *DataStorePostFightingPowerScoreParam {
	return &DataStorePostFightingPowerScoreParam{}
}
