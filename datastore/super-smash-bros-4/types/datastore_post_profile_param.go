// Package datastore_super_smash_bros_4_types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePostProfileParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
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

// String returns a string representation of the struct
func (dataStorePostProfileParam *DataStorePostProfileParam) String() string {
	return dataStorePostProfileParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePostProfileParam *DataStorePostProfileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostProfileParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePostProfileParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sProfile: %x\n", indentationValues, dataStorePostProfileParam.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostProfileParam returns a new DataStorePostProfileParam
func NewDataStorePostProfileParam() *DataStorePostProfileParam {
	return &DataStorePostProfileParam{}
}
