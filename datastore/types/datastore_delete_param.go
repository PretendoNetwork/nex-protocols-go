package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreDeleteParam struct {
	nex.Structure
	DataID         uint64
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStoreDeleteParam structure from a stream
func (dataStoreDeleteParam *DataStoreDeleteParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreDeleteParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.DataID. %s", err.Error())
	}

	dataStoreDeleteParam.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreDeleteParam.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreDeleteParam and returns a byte array
func (dataStoreDeleteParam *DataStoreDeleteParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreDeleteParam.DataID)
	stream.WriteUInt64LE(dataStoreDeleteParam.UpdatePassword)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreDeleteParam *DataStoreDeleteParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParamV1()

	copied.DataID = dataStoreDeleteParam.DataID
	copied.UpdatePassword = dataStoreDeleteParam.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreDeleteParam *DataStoreDeleteParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParamV1)

	if dataStoreDeleteParam.DataID != other.DataID {
		return false
	}

	if dataStoreDeleteParam.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
}

// NewDataStoreDeleteParam returns a new DataStoreDeleteParam
func NewDataStoreDeleteParam() *DataStoreDeleteParam {
	return &DataStoreDeleteParam{}
}
