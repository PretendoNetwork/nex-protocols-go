// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqUpdateInfo is a data structure used by the DataStore protocol
type DataStoreReqUpdateInfo struct {
	nex.Structure
	Version        uint32
	URL            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCACert     []byte
}

// ExtractFromStream extracts a DataStoreReqUpdateInfo structure from a stream
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		version, err := stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreReqUpdateInfo.Version = version
	} else {
		version, err := stream.ReadUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreReqUpdateInfo.Version = uint32(version)
	}

	dataStoreReqUpdateInfo.URL, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.URL. %s", err.Error())
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

	dataStoreReqUpdateInfo.RootCACert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.RootCACert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqUpdateInfo and returns a byte array
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt32LE(dataStoreReqUpdateInfo.Version)
	} else {
		stream.WriteUInt16LE(uint16(dataStoreReqUpdateInfo.Version))
	}

	stream.WriteString(dataStoreReqUpdateInfo.URL)
	stream.WriteListStructure(dataStoreReqUpdateInfo.RequestHeaders)
	stream.WriteListStructure(dataStoreReqUpdateInfo.FormFields)
	stream.WriteBuffer(dataStoreReqUpdateInfo.RootCACert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqUpdateInfo
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqUpdateInfo()

	copied.SetStructureVersion(dataStoreReqUpdateInfo.StructureVersion())

	copied.Version = dataStoreReqUpdateInfo.Version
	copied.URL = dataStoreReqUpdateInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqUpdateInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.FormFields))

	for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqUpdateInfo.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCACert = make([]byte, len(dataStoreReqUpdateInfo.RootCACert))

	copy(copied.RootCACert, dataStoreReqUpdateInfo.RootCACert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqUpdateInfo)

	if dataStoreReqUpdateInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreReqUpdateInfo.Version != other.Version {
		return false
	}

	if dataStoreReqUpdateInfo.URL != other.URL {
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

	return bytes.Equal(dataStoreReqUpdateInfo.RootCACert, other.RootCACert)
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
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, dataStoreReqUpdateInfo.URL))

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

	b.WriteString(fmt.Sprintf("%sRootCACert: %x\n", indentationValues, dataStoreReqUpdateInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqUpdateInfo returns a new DataStoreReqUpdateInfo
func NewDataStoreReqUpdateInfo() *DataStoreReqUpdateInfo {
	return &DataStoreReqUpdateInfo{
		Version:        0,
		URL:            "",
		RequestHeaders: make([]*DataStoreKeyValue, 0),
		FormFields:     make([]*DataStoreKeyValue, 0),
		RootCACert:     make([]byte, 0),
	}
}
