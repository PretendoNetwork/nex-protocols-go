// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCompletePostReplayParam is a data structure used by the DataStore Super Mario Maker protocol
type DataStoreCompletePostReplayParam struct {
	types.Structure
	ReplayID      *types.PrimitiveU64
	CompleteParam *datastore_types.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostReplayParam
}

// ExtractFrom extracts the DataStoreCompletePostReplayParam from the given readable
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCompletePostReplayParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCompletePostReplayParam header. %s", err.Error())
	}

	err = dataStoreCompletePostReplayParam.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.ReplayID. %s", err.Error())
	}

	err = dataStoreCompletePostReplayParam.CompleteParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.CompleteParam. %s", err.Error())
	}

	err = dataStoreCompletePostReplayParam.PrepareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.PrepareParam. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreCompletePostReplayParam to the given writable
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCompletePostReplayParam.ReplayID.WriteTo(contentWritable)
	dataStoreCompletePostReplayParam.CompleteParam.WriteTo(contentWritable)
	dataStoreCompletePostReplayParam.PrepareParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCompletePostReplayParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCompletePostReplayParam
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostReplayParam()

	copied.StructureVersion = dataStoreCompletePostReplayParam.StructureVersion

	copied.ReplayID = dataStoreCompletePostReplayParam.ReplayID.Copy().(*types.PrimitiveU64)
	copied.CompleteParam = dataStoreCompletePostReplayParam.CompleteParam.Copy().(*datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostReplayParam.PrepareParam.Copy().(*DataStorePreparePostReplayParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompletePostReplayParam); !ok {
		return false
	}

	other := o.(*DataStoreCompletePostReplayParam)

	if dataStoreCompletePostReplayParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCompletePostReplayParam.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dataStoreCompletePostReplayParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostReplayParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) String() string {
	return dataStoreCompletePostReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCompletePostReplayParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dataStoreCompletePostReplayParam.ReplayID))
	b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dataStoreCompletePostReplayParam.CompleteParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareParam: %s\n", indentationValues, dataStoreCompletePostReplayParam.PrepareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostReplayParam returns a new DataStoreCompletePostReplayParam
func NewDataStoreCompletePostReplayParam() *DataStoreCompletePostReplayParam {
	return &DataStoreCompletePostReplayParam{
		ReplayID: types.NewPrimitiveU64(0),
		CompleteParam: datastore_types.NewDataStoreCompletePostParam(),
		PrepareParam: NewDataStorePreparePostReplayParam(),
	}
}
