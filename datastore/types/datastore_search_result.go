package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

	if dataStoreSearchResult.TotalCountType != other.TotalCountType {
		return false
	}

	return true
}

// NewDataStoreSearchResult returns a new DataStoreSearchResult
func NewDataStoreSearchResult() *DataStoreSearchResult {
	return &DataStoreSearchResult{}
}
