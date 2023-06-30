// Package datastore_types implements all the types used by the DataStore protocol
package datastore_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreReqPostInfo is sent in the PreparePostObject method
type DataStoreReqPostInfo struct {
	nex.Structure
	DataID         uint64
	URL            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCACert     []byte
}

// Bytes encodes the DataStoreReqPostInfo and returns a byte array
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreReqPostInfo.DataID)
	stream.WriteString(dataStoreReqPostInfo.URL)
	stream.WriteListStructure(dataStoreReqPostInfo.RequestHeaders)
	stream.WriteListStructure(dataStoreReqPostInfo.FormFields)
	stream.WriteBuffer(dataStoreReqPostInfo.RootCACert)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqPostInfo
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqPostInfo()

	copied.DataID = dataStoreReqPostInfo.DataID
	copied.URL = dataStoreReqPostInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqPostInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqPostInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqPostInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqPostInfo.FormFields))

	for i := 0; i < len(dataStoreReqPostInfo.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqPostInfo.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCACert = make([]byte, len(dataStoreReqPostInfo.RootCACert))

	copy(copied.RootCACert, dataStoreReqPostInfo.RootCACert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqPostInfo)

	if dataStoreReqPostInfo.DataID != other.DataID {
		return false
	}

	if dataStoreReqPostInfo.URL != other.URL {
		return false
	}

	if len(dataStoreReqPostInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfo.RequestHeaders); i++ {
		if dataStoreReqPostInfo.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqPostInfo.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfo.FormFields); i++ {
		if dataStoreReqPostInfo.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	if !bytes.Equal(dataStoreReqPostInfo.RootCACert, other.RootCACert) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReqPostInfo *DataStoreReqPostInfo) String() string {
	return dataStoreReqPostInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReqPostInfo *DataStoreReqPostInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReqPostInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreReqPostInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreReqPostInfo.DataID))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, dataStoreReqPostInfo.URL))

	if len(dataStoreReqPostInfo.RequestHeaders) == 0 {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRequestHeaders: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqPostInfo.RequestHeaders); i++ {
			str := dataStoreReqPostInfo.RequestHeaders[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqPostInfo.RequestHeaders)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if len(dataStoreReqPostInfo.FormFields) == 0 {
		b.WriteString(fmt.Sprintf("%sFormFields: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sFormFields: [\n", indentationValues))

		for i := 0; i < len(dataStoreReqPostInfo.FormFields); i++ {
			str := dataStoreReqPostInfo.FormFields[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreReqPostInfo.FormFields)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sRootCACert: %x\n", indentationValues, dataStoreReqPostInfo.RootCACert))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReqPostInfo returns a new DataStoreReqPostInfo
func NewDataStoreReqPostInfo() *DataStoreReqPostInfo {
	return &DataStoreReqPostInfo{}
}
