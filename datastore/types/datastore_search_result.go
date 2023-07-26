// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSearchResult is a data structure used by the DataStore protocol
type DataStoreSearchResult struct {
	nex.Structure
	TotalCount     uint32
	Result         []*DataStoreMetaInfo
	TotalCountType uint8
}

// ExtractFromStream extracts a DataStoreSearchResult structure from a stream
func (dataStoreSearchResult *DataStoreSearchResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSearchResult.TotalCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCount. %s", err.Error())
	}

	result, err := stream.ReadListStructure(NewDataStoreMetaInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.Result. %s", err.Error())
	}

	dataStoreSearchResult.Result = result.([]*DataStoreMetaInfo)
	dataStoreSearchResult.TotalCountType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchResult.TotalCountType. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreSearchResult and returns a byte array
func (dataStoreSearchResult *DataStoreSearchResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreSearchResult.TotalCount)
	stream.WriteListStructure(dataStoreSearchResult.Result)
	stream.WriteUInt8(dataStoreSearchResult.TotalCountType)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchResult
func (dataStoreSearchResult *DataStoreSearchResult) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchResult()

	copied.TotalCount = dataStoreSearchResult.TotalCount
	copied.Result = make([]*DataStoreMetaInfo, len(dataStoreSearchResult.Result))

	for i := 0; i < len(dataStoreSearchResult.Result); i++ {
		copied.Result[i] = dataStoreSearchResult.Result[i].Copy().(*DataStoreMetaInfo)
	}

	copied.TotalCountType = dataStoreSearchResult.TotalCountType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchResult *DataStoreSearchResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchResult)

	if dataStoreSearchResult.TotalCount != other.TotalCount {
		return false
	}

	if len(dataStoreSearchResult.Result) != len(other.Result) {
		return false
	}

	for i := 0; i < len(dataStoreSearchResult.Result); i++ {
		if dataStoreSearchResult.Result[i] != other.Result[i] {
			return false
		}
	}

	return dataStoreSearchResult.TotalCountType == other.TotalCountType
}

// String returns a string representation of the struct
func (dataStoreSearchResult *DataStoreSearchResult) String() string {
	return dataStoreSearchResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchResult *DataStoreSearchResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchResult{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSearchResult.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTotalCount: %d,\n", indentationValues, dataStoreSearchResult.TotalCount))

	if len(dataStoreSearchResult.Result) == 0 {
		b.WriteString(fmt.Sprintf("%sResult: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: [\n", indentationValues))

		for i := 0; i < len(dataStoreSearchResult.Result); i++ {
			str := dataStoreSearchResult.Result[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreSearchResult.Result)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTotalCountType: %d\n", indentationValues, dataStoreSearchResult.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchResult returns a new DataStoreSearchResult
func NewDataStoreSearchResult() *DataStoreSearchResult {
	return &DataStoreSearchResult{}
}
