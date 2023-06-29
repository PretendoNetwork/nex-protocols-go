package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRatingInfo is a data structure used by the DataStore protocol
type DataStoreRatingInfo struct {
	nex.Structure
	TotalValue   int64
	Count        uint32
	InitialValue int64
}

// ExtractFromStream extracts a DataStoreRatingInfo structure from a stream
func (dataStoreRatingInfo *DataStoreRatingInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRatingInfo.TotalValue, err = stream.ReadInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.TotalValue. %s", err.Error())
	}

	dataStoreRatingInfo.Count, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.Count. %s", err.Error())
	}

	dataStoreRatingInfo.InitialValue, err = stream.ReadInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRatingInfo.InitialValue. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreRatingInfo and returns a byte array
func (dataStoreRatingInfo *DataStoreRatingInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uint64(dataStoreRatingInfo.TotalValue))
	stream.WriteUInt32LE(dataStoreRatingInfo.Count)
	stream.WriteUInt64LE(uint64(dataStoreRatingInfo.InitialValue))

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRatingInfo
func (dataStoreRatingInfo *DataStoreRatingInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInfo()

	copied.TotalValue = dataStoreRatingInfo.TotalValue
	copied.Count = dataStoreRatingInfo.Count
	copied.InitialValue = dataStoreRatingInfo.InitialValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfo *DataStoreRatingInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInfo)

	if dataStoreRatingInfo.TotalValue != other.TotalValue {
		return false
	}

	if dataStoreRatingInfo.Count != other.Count {
		return false
	}

	if dataStoreRatingInfo.InitialValue != other.InitialValue {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRatingInfo *DataStoreRatingInfo) String() string {
	return dataStoreRatingInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRatingInfo *DataStoreRatingInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRatingInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRatingInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTotalValue: %d,\n", indentationValues, dataStoreRatingInfo.TotalValue))
	b.WriteString(fmt.Sprintf("%sCount: %d,\n", indentationValues, dataStoreRatingInfo.Count))
	b.WriteString(fmt.Sprintf("%sInitialValue: %d\n", indentationValues, dataStoreRatingInfo.InitialValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRatingInfo returns a new DataStoreRatingInfo
func NewDataStoreRatingInfo() *DataStoreRatingInfo {
	return &DataStoreRatingInfo{}
}
