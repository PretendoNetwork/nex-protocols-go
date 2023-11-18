// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqPostInfoV1 is a data structure used by the DataStore protocol
type DataStoreReqPostInfoV1 struct {
	nex.Structure
	DataID         uint32
	URL            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCACert     []byte
}

// ExtractFromStream extracts a DataStoreReqPostInfoV1 structure from a stream
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreReqPostInfoV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.DataID. %s", err.Error())
	}

	dataStoreReqPostInfoV1.URL, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.URL. %s", err.Error())
	}

	requestHeaders, err := nex.StreamReadListStructure(stream, NewDataStoreKeyValue())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RequestHeaders. %s", err.Error())
	}

	dataStoreReqPostInfoV1.RequestHeaders = requestHeaders

	formFields, err := nex.StreamReadListStructure(stream, NewDataStoreKeyValue())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.FormFields. %s", err.Error())
	}

	dataStoreReqPostInfoV1.FormFields = formFields

	dataStoreReqPostInfoV1.RootCACert, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReqPostInfoV1.RootCACert. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreReqPostInfoV1 and returns a byte array
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqPostInfoV1.DataID)
	stream.WriteString(dataStoreReqPostInfoV1.URL)
	nex.StreamWriteListStructure(stream, dataStoreReqPostInfoV1.RequestHeaders)
	nex.StreamWriteListStructure(stream, dataStoreReqPostInfoV1.FormFields)
	stream.WriteBuffer(dataStoreReqPostInfoV1.RootCACert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqPostInfoV1
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreReqPostInfoV1()

	copied.SetStructureVersion(dataStoreReqPostInfoV1.StructureVersion())

	copied.DataID = dataStoreReqPostInfoV1.DataID
	copied.URL = dataStoreReqPostInfoV1.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqPostInfoV1.RequestHeaders))

	for i := 0; i < len(dataStoreReqPostInfoV1.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqPostInfoV1.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqPostInfoV1.FormFields))

	for i := 0; i < len(dataStoreReqPostInfoV1.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqPostInfoV1.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCACert = make([]byte, len(dataStoreReqPostInfoV1.RootCACert))

	copy(copied.RootCACert, dataStoreReqPostInfoV1.RootCACert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqPostInfoV1)

	if dataStoreReqPostInfoV1.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreReqPostInfoV1.DataID != other.DataID {
		return false
	}

	if dataStoreReqPostInfoV1.URL != other.URL {
		return false
	}

	if len(dataStoreReqPostInfoV1.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfoV1.RequestHeaders); i++ {
		if dataStoreReqPostInfoV1.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqPostInfoV1.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfoV1.FormFields); i++ {
		if dataStoreReqPostInfoV1.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	return bytes.Equal(dataStoreReqPostInfoV1.RootCACert, other.RootCACert)
}

// String returns a string representation of the struct
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) String() string {
	return dataStoreReqPostInfoV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfoV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqPostInfoV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreReqPostInfoV1.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, dataStoreReqPostInfoV1.URL))

	if len(dataStoreReqPostInfoV1.RequestHeaders) == 0 {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqPostInfoV1.RequestHeaders); i++ {
			str := dataStoreReqPostInfoV1.RequestHeaders[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqPostInfoV1.RequestHeaders)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if len(dataStoreReqPostInfoV1.FormFields) == 0 {
		b.WriteString(fmt.Sprintf("%sFormFields: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sFormFields: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqPostInfoV1.FormFields); i++ {
			str := dataStoreReqPostInfoV1.FormFields[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqPostInfoV1.FormFields)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sRootCACert: %x\n", indentationValues, dataStoreReqPostInfoV1.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfoV1 returns a new DataStoreReqPostInfoV1
func NewDataStoreReqPostInfoV1() *DataStoreReqPostInfoV1 {
	return &DataStoreReqPostInfoV1{
		DataID:         0,
		URL:            "",
		RequestHeaders: make([]*DataStoreKeyValue, 0),
		FormFields:     make([]*DataStoreKeyValue, 0),
		RootCACert:     make([]byte, 0),
	}
}
