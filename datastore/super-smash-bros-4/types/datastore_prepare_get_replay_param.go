// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePrepareGetReplayParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePrepareGetReplayParam struct {
	types.Structure
	ReplayID  *types.PrimitiveU64
	ExtraData *types.List[*types.String]
}

// ExtractFrom extracts the DataStorePrepareGetReplayParam from the given readable
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePrepareGetReplayParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePrepareGetReplayParam header. %s", err.Error())
	}

	err = dataStorePrepareGetReplayParam.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ReplayID. %s", err.Error())
	}

	err = dataStorePrepareGetReplayParam.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePrepareGetReplayParam to the given writable
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePrepareGetReplayParam.ReplayID.WriteTo(contentWritable)
	dataStorePrepareGetReplayParam.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePrepareGetReplayParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePrepareGetReplayParam
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Copy() types.RVType {
	copied := NewDataStorePrepareGetReplayParam()

	copied.StructureVersion = dataStorePrepareGetReplayParam.StructureVersion

	copied.ReplayID = dataStorePrepareGetReplayParam.ReplayID.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dataStorePrepareGetReplayParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetReplayParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetReplayParam)

	if dataStorePrepareGetReplayParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePrepareGetReplayParam.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dataStorePrepareGetReplayParam.ExtraData.Equals(other.ExtraData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) String() string {
	return dataStorePrepareGetReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePrepareGetReplayParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dataStorePrepareGetReplayParam.ReplayID))
	b.WriteString(fmt.Sprintf("%sExtraData: %s\n", indentationValues, dataStorePrepareGetReplayParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetReplayParam returns a new DataStorePrepareGetReplayParam
func NewDataStorePrepareGetReplayParam() *DataStorePrepareGetReplayParam {
	dataStorePrepareGetReplayParam := &DataStorePrepareGetReplayParam{
		ReplayID: types.NewPrimitiveU64(0),
		ExtraData: types.NewList[*types.String](),
	}

	dataStorePrepareGetReplayParam.ExtraData.Type = types.NewString("")

	return dataStorePrepareGetReplayParam
}
