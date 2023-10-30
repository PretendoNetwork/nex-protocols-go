// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetMetaByOwnerIDParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetMetaByOwnerIDParam struct {
	nex.Structure
	OwnerIDs     []uint32
	DataTypes    []uint16
	ResultOption uint8
	ResultRange  *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreGetMetaByOwnerIDParam structure from a stream
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetMetaByOwnerIDParam.OwnerIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.OwnerIDs from stream. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.DataTypes, err = stream.ReadListUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.DataTypes from stream. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultOption from stream. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetMetaByOwnerIDParam.ResultRange from stream. %s", err.Error())
	}

	dataStoreGetMetaByOwnerIDParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreGetMetaByOwnerIDParam and returns a byte array
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetMetaByOwnerIDParam.OwnerIDs)
	stream.WriteListUInt16LE(dataStoreGetMetaByOwnerIDParam.DataTypes)
	stream.WriteUInt8(dataStoreGetMetaByOwnerIDParam.ResultOption)
	stream.WriteStructure(dataStoreGetMetaByOwnerIDParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetMetaByOwnerIDParam
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetMetaByOwnerIDParam()

	copied.SetStructureVersion(dataStoreGetMetaByOwnerIDParam.StructureVersion())

	copied.OwnerIDs = make([]uint32, len(dataStoreGetMetaByOwnerIDParam.OwnerIDs))

	copy(copied.OwnerIDs, dataStoreGetMetaByOwnerIDParam.OwnerIDs)

	copied.DataTypes = make([]uint16, len(dataStoreGetMetaByOwnerIDParam.DataTypes))

	copy(copied.DataTypes, dataStoreGetMetaByOwnerIDParam.DataTypes)

	copied.ResultOption = dataStoreGetMetaByOwnerIDParam.ResultOption
	copied.ResultRange = dataStoreGetMetaByOwnerIDParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetMetaByOwnerIDParam)

	if dataStoreGetMetaByOwnerIDParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if len(dataStoreGetMetaByOwnerIDParam.OwnerIDs) != len(other.OwnerIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetMetaByOwnerIDParam.OwnerIDs); i++ {
		if dataStoreGetMetaByOwnerIDParam.OwnerIDs[i] != other.OwnerIDs[i] {
			return false
		}
	}

	if len(dataStoreGetMetaByOwnerIDParam.DataTypes) != len(other.DataTypes) {
		return false
	}

	for i := 0; i < len(dataStoreGetMetaByOwnerIDParam.DataTypes); i++ {
		if dataStoreGetMetaByOwnerIDParam.DataTypes[i] != other.DataTypes[i] {
			return false
		}
	}

	if dataStoreGetMetaByOwnerIDParam.ResultOption != other.ResultOption {
		return false
	}

	if !dataStoreGetMetaByOwnerIDParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) String() string {
	return dataStoreGetMetaByOwnerIDParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetMetaByOwnerIDParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOwnerIDs: %v,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.OwnerIDs))
	b.WriteString(fmt.Sprintf("%sDataTypes: %v,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.DataTypes))
	b.WriteString(fmt.Sprintf("%sResultOption: %d,\n", indentationValues, dataStoreGetMetaByOwnerIDParam.ResultOption))

	if dataStoreGetMetaByOwnerIDParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreGetMetaByOwnerIDParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	return &DataStoreGetMetaByOwnerIDParam{
		OwnerIDs:     make([]uint32, 0),
		DataTypes:    make([]uint16, 0),
		ResultOption: 0,
		ResultRange:  nex.NewResultRange(),
	}
}
