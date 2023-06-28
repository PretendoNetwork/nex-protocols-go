package datastore_types

import (
	"bytes"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqGetInfo is sent in the PrepareGetObject method
type DataStoreReqGetInfo struct {
	nex.Structure
	URL            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCA         []byte
	DataID         uint64 // NEX 3.5.0+
}

// Bytes encodes the DataStoreReqGetInfo and returns a byte array
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	stream.WriteString(dataStoreReqGetInfo.URL)
	stream.WriteListStructure(dataStoreReqGetInfo.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfo.Size)
	stream.WriteBuffer(dataStoreReqGetInfo.RootCA)

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		stream.WriteUInt64LE(dataStoreReqGetInfo.DataID)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetInfo
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetInfo()

	copied.URL = dataStoreReqGetInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqGetInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqGetInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.Size = dataStoreReqGetInfo.Size
	copied.RootCA = make([]byte, len(dataStoreReqGetInfo.RootCA))

	copy(copied.RootCA, dataStoreReqGetInfo.RootCA)

	copied.DataID = dataStoreReqGetInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetInfo)

	if dataStoreReqGetInfo.URL != other.URL {
		return false
	}

	if len(dataStoreReqGetInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
		if !dataStoreReqGetInfo.RequestHeaders[i].Equals(other.RequestHeaders[i]) {
			return false
		}
	}

	if dataStoreReqGetInfo.Size != other.Size {
		return false
	}

	if !bytes.Equal(dataStoreReqGetInfo.RootCA, other.RootCA) {
		return false
	}

	if dataStoreReqGetInfo.DataID != other.DataID {
		return false
	}

	return true
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	return &DataStoreReqGetInfo{}
}
