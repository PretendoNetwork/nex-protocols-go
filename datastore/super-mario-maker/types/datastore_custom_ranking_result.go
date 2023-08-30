// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCustomRankingResult holds data for the DataStore (Super Mario Maker) protocol
type DataStoreCustomRankingResult struct {
	nex.Structure
	Order    uint32
	Score    uint32
	MetaInfo *datastore_types.DataStoreMetaInfo
}

// ExtractFromStream extracts a DataStoreCustomRankingResult structure from a stream
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCustomRankingResult.Order, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Order from stream. %s", err.Error())
	}

	dataStoreCustomRankingResult.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Score from stream. %s", err.Error())
	}

	metaInfo, err := stream.ReadStructure(datastore_types.NewDataStoreMetaInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.MetaInfo from stream. %s", err.Error())
	}

	dataStoreCustomRankingResult.MetaInfo = metaInfo.(*datastore_types.DataStoreMetaInfo)

	return nil
}

// Bytes encodes the DataStoreCustomRankingResult and returns a byte array
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Order)
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Score)
	stream.WriteStructure(dataStoreCustomRankingResult.MetaInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCustomRankingResult
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Copy() nex.StructureInterface {
	copied := NewDataStoreCustomRankingResult()

	copied.SetStructureVersion(dataStoreCustomRankingResult.StructureVersion())

	copied.Order = dataStoreCustomRankingResult.Order
	copied.Score = dataStoreCustomRankingResult.Score
	copied.MetaInfo = dataStoreCustomRankingResult.MetaInfo.Copy().(*datastore_types.DataStoreMetaInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCustomRankingResult)

	if dataStoreCustomRankingResult.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreCustomRankingResult.Order != other.Order {
		return false
	}

	if dataStoreCustomRankingResult.Score != other.Score {
		return false
	}

	if !dataStoreCustomRankingResult.MetaInfo.Equals(other.MetaInfo) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) String() string {
	return dataStoreCustomRankingResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCustomRankingResult{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCustomRankingResult.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOrder: %d,\n", indentationValues, dataStoreCustomRankingResult.Order))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, dataStoreCustomRankingResult.Score))

	if dataStoreCustomRankingResult.MetaInfo != nil {
		b.WriteString(fmt.Sprintf("%sMetaInfo: %s\n", indentationValues, dataStoreCustomRankingResult.MetaInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMetaInfo: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{}
}
