package datastore_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreReqGetInfoV1 struct {
	nex.Structure
	Url            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCaCert     []byte
}

// ExtractFromStream extracts a DataStoreReqGetInfoV1 structure from a stream
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqGetInfoV1.Url, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.Url. %s", err.Error())
	}

	requestHeaders, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RequestHeaders. %s", err.Error())
	}

	dataStoreReqGetInfoV1.RequestHeaders = requestHeaders.([]*DataStoreKeyValue)
	dataStoreReqGetInfoV1.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.Size. %s", err.Error())
	}

	dataStoreReqGetInfoV1.RootCaCert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqGetInfoV1.RootCaCert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqGetInfoV1 and returns a byte array
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetInfoV1.Url)
	stream.WriteListStructure(dataStoreReqGetInfoV1.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfoV1.Size)
	stream.WriteBuffer(dataStoreReqGetInfoV1.RootCaCert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetInfoV1
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetInfoV1()

	copied.Url = dataStoreReqGetInfoV1.Url
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqGetInfoV1.RequestHeaders))

	for i := 0; i < len(dataStoreReqGetInfoV1.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqGetInfoV1.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.Size = dataStoreReqGetInfoV1.Size

	copied.RootCaCert = make([]byte, len(dataStoreReqGetInfoV1.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqGetInfoV1.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetInfoV1)

	if dataStoreReqGetInfoV1.Url != other.Url {
		return false
	}

	if len(dataStoreReqGetInfoV1.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqGetInfoV1.RequestHeaders); i++ {
		if dataStoreReqGetInfoV1.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if dataStoreReqGetInfoV1.Size != other.Size {
		return false
	}

	if !bytes.Equal(dataStoreReqGetInfoV1.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) String() string {
	return dataStoreReqGetInfoV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqGetInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqGetInfoV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUrl: %q,\n", indentationValues, dataStoreReqGetInfoV1.Url))

	if len(dataStoreReqGetInfoV1.RequestHeaders) == 0 {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqGetInfoV1.RequestHeaders); i++ {
			str := dataStoreReqGetInfoV1.RequestHeaders[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqGetInfoV1.RequestHeaders)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreReqGetInfoV1.Size))
	b.WriteString(fmt.Sprintf("%sRootCaCert: %x\n", indentationValues, dataStoreReqGetInfoV1.RootCaCert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqGetInfoV1 returns a new DataStoreReqGetInfoV1
func NewDataStoreReqGetInfoV1() *DataStoreReqGetInfoV1 {
	return &DataStoreReqGetInfoV1{}
}
