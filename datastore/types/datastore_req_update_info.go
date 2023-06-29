package datastore_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreReqUpdateInfo struct {
	nex.Structure
	Version        uint32
	Url            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCaCert     []byte
}

// ExtractFromStream extracts a DataStoreReqUpdateInfo structure from a stream
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqUpdateInfo.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
	}

	dataStoreReqUpdateInfo.Url, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Url. %s", err.Error())
	}

	requestHeaders, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RequestHeaders. %s", err.Error())
	}

	dataStoreReqUpdateInfo.RequestHeaders = requestHeaders.([]*DataStoreKeyValue)

	formFields, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.FormFields. %s", err.Error())
	}

	dataStoreReqUpdateInfo.FormFields = formFields.([]*DataStoreKeyValue)

	dataStoreReqUpdateInfo.RootCaCert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RootCaCert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqUpdateInfo and returns a byte array
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqUpdateInfo.Version)
	stream.WriteString(dataStoreReqUpdateInfo.Url)
	stream.WriteListStructure(dataStoreReqUpdateInfo.RequestHeaders)
	stream.WriteListStructure(dataStoreReqUpdateInfo.FormFields)
	stream.WriteBuffer(dataStoreReqUpdateInfo.RootCaCert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqUpdateInfo
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqUpdateInfo()

	copied.Version = dataStoreReqUpdateInfo.Version
	copied.Url = dataStoreReqUpdateInfo.Url
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqUpdateInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.FormFields))

	for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqUpdateInfo.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCaCert = make([]byte, len(dataStoreReqUpdateInfo.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqUpdateInfo.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqUpdateInfo)

	if dataStoreReqUpdateInfo.Version != other.Version {
		return false
	}

	if dataStoreReqUpdateInfo.Url != other.Url {
		return false
	}

	if len(dataStoreReqUpdateInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
		if dataStoreReqUpdateInfo.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqUpdateInfo.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
		if dataStoreReqUpdateInfo.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	if !bytes.Equal(dataStoreReqUpdateInfo.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) String() string {
	return dataStoreReqUpdateInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqUpdateInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqUpdateInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUrl: %q,\n", indentationValues, dataStoreReqUpdateInfo.Url))

	if len(dataStoreReqUpdateInfo.RequestHeaders) == 0 {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
			str := dataStoreReqUpdateInfo.RequestHeaders[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqUpdateInfo.RequestHeaders)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if len(dataStoreReqUpdateInfo.FormFields) == 0 {
		b.WriteString(fmt.Sprintf("%sFormFields: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sFormFields: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
			str := dataStoreReqUpdateInfo.FormFields[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqUpdateInfo.FormFields)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sRootCaCert: %x\n", indentationValues, dataStoreReqUpdateInfo.RootCaCert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqUpdateInfo returns a new DataStoreReqUpdateInfo
func NewDataStoreReqUpdateInfo() *DataStoreReqUpdateInfo {
	return &DataStoreReqUpdateInfo{}
}
