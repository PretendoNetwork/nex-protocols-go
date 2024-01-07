// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePostProfileParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePostProfileParam struct {
	types.Structure
	Profile *types.QBuffer
}

// ExtractFrom extracts the DataStorePostProfileParam from the given readable
func (dataStorePostProfileParam *DataStorePostProfileParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePostProfileParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePostProfileParam header. %s", err.Error())
	}

	err = dataStorePostProfileParam.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostProfileParam.Profile. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePostProfileParam to the given writable
func (dataStorePostProfileParam *DataStorePostProfileParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePostProfileParam.Profile.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePostProfileParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePostProfileParam
func (dataStorePostProfileParam *DataStorePostProfileParam) Copy() types.RVType {
	copied := NewDataStorePostProfileParam()

	copied.StructureVersion = dataStorePostProfileParam.StructureVersion

	copied.Profile = dataStorePostProfileParam.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostProfileParam *DataStorePostProfileParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePostProfileParam); !ok {
		return false
	}

	other := o.(*DataStorePostProfileParam)

	if dataStorePostProfileParam.StructureVersion != other.StructureVersion {
		return false
	}

	return dataStorePostProfileParam.Profile.Equals(other.Profile)
}

// String returns a string representation of the struct
func (dataStorePostProfileParam *DataStorePostProfileParam) String() string {
	return dataStorePostProfileParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePostProfileParam *DataStorePostProfileParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostProfileParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePostProfileParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sProfile: %s\n", indentationValues, dataStorePostProfileParam.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostProfileParam returns a new DataStorePostProfileParam
func NewDataStorePostProfileParam() *DataStorePostProfileParam {
	return &DataStorePostProfileParam{
		Profile: types.NewQBuffer(nil),
	}
}
