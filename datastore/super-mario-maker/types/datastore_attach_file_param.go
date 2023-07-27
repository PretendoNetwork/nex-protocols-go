// Package types implements all the types used by the DataStore Super Mario Maker protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreAttachFileParam is sent in the PrepareAttachFile method
type DataStoreAttachFileParam struct {
	nex.Structure
	PostParam   *datastore_types.DataStorePreparePostParam
	ReferDataID uint64
	ContentType string
}

// ExtractFromStream extracts a DataStoreAttachFileParam structure from a stream
func (dataStoreAttachFileParam *DataStoreAttachFileParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	postParam, err := stream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.PostParam. %s", err.Error())
	}

	dataStoreAttachFileParam.PostParam = postParam.(*datastore_types.DataStorePreparePostParam)
	dataStoreAttachFileParam.ReferDataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ReferDataID. %s", err.Error())
	}

	dataStoreAttachFileParam.ContentType, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ContentType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreAttachFileParam
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Copy() nex.StructureInterface {
	copied := NewDataStoreAttachFileParam()

	copied.PostParam = dataStoreAttachFileParam.PostParam.Copy().(*datastore_types.DataStorePreparePostParam)
	copied.ReferDataID = dataStoreAttachFileParam.ReferDataID
	copied.ContentType = dataStoreAttachFileParam.ContentType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreAttachFileParam)

	if !dataStoreAttachFileParam.PostParam.Equals(other.PostParam) {
		return false
	}

	if dataStoreAttachFileParam.ReferDataID != other.ReferDataID {
		return false
	}

	if dataStoreAttachFileParam.ContentType != other.ContentType {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreAttachFileParam *DataStoreAttachFileParam) String() string {
	return dataStoreAttachFileParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreAttachFileParam *DataStoreAttachFileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreAttachFileParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreAttachFileParam.StructureVersion()))

	if dataStoreAttachFileParam.PostParam != nil {
		b.WriteString(fmt.Sprintf("%sPostParam: %s,\n", indentationValues, dataStoreAttachFileParam.PostParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPostParam: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sReferDataID: %d,\n", indentationValues, dataStoreAttachFileParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sContentType: %q\n", indentationValues, dataStoreAttachFileParam.ContentType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreAttachFileParam returns a new DataStoreAttachFileParam
func NewDataStoreAttachFileParam() *DataStoreAttachFileParam {
	return &DataStoreAttachFileParam{}
}
