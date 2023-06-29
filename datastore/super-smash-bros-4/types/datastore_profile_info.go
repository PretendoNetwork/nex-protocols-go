package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreProfileInfo is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreProfileInfo struct {
	nex.Structure
	PID     uint32
	Profile []byte
}

// ExtractFromStream extracts a DataStoreProfileInfo structure from a stream
func (dataStoreProfileInfo *DataStoreProfileInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreProfileInfo.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.PID. %s", err.Error())
	}

	dataStoreProfileInfo.Profile, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.Profile. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreProfileInfo and returns a byte array
func (dataStoreProfileInfo *DataStoreProfileInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreProfileInfo.PID)
	stream.WriteQBuffer(dataStoreProfileInfo.Profile)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreProfileInfo
func (dataStoreProfileInfo *DataStoreProfileInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreProfileInfo()

	copied.PID = dataStoreProfileInfo.PID
	copied.Profile = make([]byte, len(dataStoreProfileInfo.Profile))

	copy(copied.Profile, dataStoreProfileInfo.Profile)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreProfileInfo *DataStoreProfileInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreProfileInfo)

	if dataStoreProfileInfo.PID != other.PID {
		return false
	}

	if !bytes.Equal(dataStoreProfileInfo.Profile, other.Profile) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreProfileInfo *DataStoreProfileInfo) String() string {
	return dataStoreProfileInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreProfileInfo *DataStoreProfileInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreProfileInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreProfileInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, dataStoreProfileInfo.PID))
	b.WriteString(fmt.Sprintf("%sProfile: %x\n", indentationValues, dataStoreProfileInfo.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreProfileInfo returns a new DataStoreProfileInfo
func NewDataStoreProfileInfo() *DataStoreProfileInfo {
	return &DataStoreProfileInfo{}
}
