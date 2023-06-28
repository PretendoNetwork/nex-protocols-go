package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStorePostProfileParam struct {
	nex.Structure
	Profile []byte
}

// ExtractFromStream extracts a DataStorePostProfileParam structure from a stream
func (dataStorePostProfileParam *DataStorePostProfileParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePostProfileParam.Profile, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostProfileParam.Profile. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePostProfileParam and returns a byte array
func (dataStorePostProfileParam *DataStorePostProfileParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(dataStorePostProfileParam.Profile)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePostProfileParam
func (dataStorePostProfileParam *DataStorePostProfileParam) Copy() nex.StructureInterface {
	copied := NewDataStorePostProfileParam()

	copied.Profile = make([]byte, len(dataStorePostProfileParam.Profile))

	copy(copied.Profile, dataStorePostProfileParam.Profile)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostProfileParam *DataStorePostProfileParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePostProfileParam)

	if !bytes.Equal(dataStorePostProfileParam.Profile, other.Profile) {
		return false
	}

	return true
}

// NewDataStorePostProfileParam returns a new DataStorePostProfileParam
func NewDataStorePostProfileParam() *DataStorePostProfileParam {
	return &DataStorePostProfileParam{}
}
