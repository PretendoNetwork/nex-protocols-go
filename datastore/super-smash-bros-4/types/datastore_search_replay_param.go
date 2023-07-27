// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSearchReplayParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreSearchReplayParam struct {
	nex.Structure
	Mode        uint8
	Style       uint8
	Fighter     uint8
	ResultRange *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreSearchReplayParam structure from a stream
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreSearchReplayParam.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Mode. %s", err.Error())
	}
	dataStoreSearchReplayParam.Style, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Style. %s", err.Error())
	}
	dataStoreSearchReplayParam.Fighter, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.Fighter. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchReplayParam.ResultRange. %s", err.Error())
	}

	dataStoreSearchReplayParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreSearchReplayParam and returns a byte array
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreSearchReplayParam.Mode)
	stream.WriteUInt8(dataStoreSearchReplayParam.Style)
	stream.WriteUInt8(dataStoreSearchReplayParam.Fighter)
	stream.WriteStructure(dataStoreSearchReplayParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchReplayParam
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchReplayParam()

	copied.Mode = dataStoreSearchReplayParam.Mode
	copied.Style = dataStoreSearchReplayParam.Style
	copied.Fighter = dataStoreSearchReplayParam.Fighter
	copied.ResultRange = dataStoreSearchReplayParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchReplayParam)

	if dataStoreSearchReplayParam.Mode != other.Mode {
		return false
	}

	if dataStoreSearchReplayParam.Style != other.Style {
		return false
	}

	if dataStoreSearchReplayParam.Fighter != other.Fighter {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSearchReplayParam.StructureVersion()))
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
