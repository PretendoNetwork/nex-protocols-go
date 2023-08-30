// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRateCustomRankingParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreRateCustomRankingParam struct {
	nex.Structure
	DataID        uint64
	ApplicationID uint32
	Score         uint32
	Period        uint16
}

// ExtractFromStream extracts a DataStoreRateCustomRankingParam structure from a stream
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRateCustomRankingParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID from stream. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.ApplicationID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationID from stream. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score from stream. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreRateCustomRankingParam and returns a byte array
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreRateCustomRankingParam.DataID)
	stream.WriteUInt32LE(dataStoreRateCustomRankingParam.ApplicationID)
	stream.WriteUInt32LE(dataStoreRateCustomRankingParam.Score)
	stream.WriteUInt16LE(dataStoreRateCustomRankingParam.Period)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateCustomRankingParam()

	copied.SetStructureVersion(dataStoreRateCustomRankingParam.StructureVersion())

	copied.DataID = dataStoreRateCustomRankingParam.DataID
	copied.ApplicationID = dataStoreRateCustomRankingParam.ApplicationID
	copied.Score = dataStoreRateCustomRankingParam.Score
	copied.Period = dataStoreRateCustomRankingParam.Period

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateCustomRankingParam)

	if dataStoreRateCustomRankingParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreRateCustomRankingParam.DataID != other.DataID {
		return false
	}

	if dataStoreRateCustomRankingParam.ApplicationID != other.ApplicationID {
		return false
	}

	if dataStoreRateCustomRankingParam.Score != other.Score {
		return false
	}

	if dataStoreRateCustomRankingParam.Period != other.Period {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) String() string {
	return dataStoreRateCustomRankingParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRateCustomRankingParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreRateCustomRankingParam.DataID))
	b.WriteString(fmt.Sprintf("%sApplicationID: %d,\n", indentationValues, dataStoreRateCustomRankingParam.ApplicationID))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, dataStoreRateCustomRankingParam.Score))
	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStoreRateCustomRankingParam.Period))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	return &DataStoreRateCustomRankingParam{}
}
