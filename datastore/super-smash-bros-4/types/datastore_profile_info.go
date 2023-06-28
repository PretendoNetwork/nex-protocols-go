package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreProfileInfo struct {
	nex.Structure
	Pid     uint32
	Profile []byte
}

// ExtractFromStream extracts a DataStoreProfileInfo structure from a stream
func (dataStoreProfileInfo *DataStoreProfileInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreProfileInfo.Pid, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.Pid. %s", err.Error())
	}

	dataStoreProfileInfo.Profile, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.Profile. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreProfileInfo and returns a byte array
func (dataStoreProfileInfo *DataStoreProfileInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreProfileInfo.Pid)
	stream.WriteQBuffer(dataStoreProfileInfo.Profile)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreProfileInfo
func (dataStoreProfileInfo *DataStoreProfileInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreProfileInfo()

	copied.Pid = dataStoreProfileInfo.Pid
	copied.Profile = make([]byte, len(dataStoreProfileInfo.Profile))

	copy(copied.Profile, dataStoreProfileInfo.Profile)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreProfileInfo *DataStoreProfileInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreProfileInfo)

	if dataStoreProfileInfo.Pid != other.Pid {
		return false
	}

	if !bytes.Equal(dataStoreProfileInfo.Profile, other.Profile) {
		return false
	}

	return true
}

// NewDataStoreProfileInfo returns a new DataStoreProfileInfo
func NewDataStoreProfileInfo() *DataStoreProfileInfo {
	return &DataStoreProfileInfo{}
}
