// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqGetInfo is sent in the PrepareGetObject method
type DataStoreReqGetInfo struct {
	nex.Structure
	URL            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCACert     []byte
	DataID         uint64 // NEX 3.5.0+
}

// Bytes encodes the DataStoreReqGetInfo and returns a byte array
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	stream.WriteString(dataStoreReqGetInfo.URL)
	nex.StreamWriteListStructure(stream, dataStoreReqGetInfo.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfo.Size)
	stream.WriteBuffer(dataStoreReqGetInfo.RootCACert)

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		stream.WriteUInt64LE(dataStoreReqGetInfo.DataID)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetInfo
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetInfo()

	copied.SetStructureVersion(dataStoreReqGetInfo.StructureVersion())

	copied.URL = dataStoreReqGetInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqGetInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqGetInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.Size = dataStoreReqGetInfo.Size
	copied.RootCACert = make([]byte, len(dataStoreReqGetInfo.RootCACert))

	copy(copied.RootCACert, dataStoreReqGetInfo.RootCACert)

	copied.DataID = dataStoreReqGetInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetInfo)

	if dataStoreReqGetInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

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

	if !bytes.Equal(dataStoreReqGetInfo.RootCACert, other.RootCACert) {
		return false
	}

	if dataStoreReqGetInfo.DataID != other.DataID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetInfo *DataStoreReqGetInfo) String() string {
	return dataStoreReqGetInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetInfo *DataStoreReqGetInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqGetInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, dataStoreReqGetInfo.URL))

	if len(dataStoreReqGetInfo.RequestHeaders) == 0 {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
			str := dataStoreReqGetInfo.RequestHeaders[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqGetInfo.RequestHeaders)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreReqGetInfo.Size))
	b.WriteString(fmt.Sprintf("%sRootCA: %x,\n", indentationValues, dataStoreReqGetInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%sDataID: %d\n", indentationValues, dataStoreReqGetInfo.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	return &DataStoreReqGetInfo{
		URL:            "",
		RequestHeaders: make([]*DataStoreKeyValue, 0),
		Size:           0,
		RootCACert:     make([]byte, 0),
		DataID:         0,
	}
}
