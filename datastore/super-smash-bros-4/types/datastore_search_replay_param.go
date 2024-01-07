// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSearchReplayParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreSearchReplayParam struct {
	types.Structure
	Mode        *types.PrimitiveU8
	Style       *types.PrimitiveU8
	Fighter     *types.PrimitiveU8
	ResultRange *types.ResultRange
}

// ExtractFrom extracts the DataStoreSearchReplayParam from the given readable
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreSearchReplayParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSearchReplayParam header. %s", err.Error())
	}

	err = dataStoreSearchReplayParam.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Mode. %s", err.Error())
	}
	err = dataStoreSearchReplayParam.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Style. %s", err.Error())
	}
	err = dataStoreSearchReplayParam.Fighter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Fighter. %s", err.Error())
	}

	err = dataStoreSearchReplayParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.ResultRange. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreSearchReplayParam to the given writable
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreSearchReplayParam.Mode.WriteTo(contentWritable)
	dataStoreSearchReplayParam.Style.WriteTo(contentWritable)
	dataStoreSearchReplayParam.Fighter.WriteTo(contentWritable)
	dataStoreSearchReplayParam.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreSearchReplayParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreSearchReplayParam
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Copy() types.RVType {
	copied := NewDataStoreSearchReplayParam()

	copied.StructureVersion = dataStoreSearchReplayParam.StructureVersion

	copied.Mode = dataStoreSearchReplayParam.Mode
	copied.Style = dataStoreSearchReplayParam.Style
	copied.Fighter = dataStoreSearchReplayParam.Fighter
	copied.ResultRange = dataStoreSearchReplayParam.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchReplayParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchReplayParam)

	if dataStoreSearchReplayParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreSearchReplayParam.Mode.Equals(other.Mode) {
		return false
	}

	if !dataStoreSearchReplayParam.Style.Equals(other.Style) {
		return false
	}

	if !dataStoreSearchReplayParam.Fighter.Equals(other.Fighter) {
		return false
	}

	if !dataStoreSearchReplayParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) String() string {
	return dataStoreSearchReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreSearchReplayParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMode: %d,\n", indentationValues, dataStoreSearchReplayParam.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %d,\n", indentationValues, dataStoreSearchReplayParam.Style))
	b.WriteString(fmt.Sprintf("%sFighter: %d,\n", indentationValues, dataStoreSearchReplayParam.Fighter))

	if dataStoreSearchReplayParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreSearchReplayParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchReplayParam returns a new DataStoreSearchReplayParam
func NewDataStoreSearchReplayParam() *DataStoreSearchReplayParam {
	return &DataStoreSearchReplayParam{}
}
