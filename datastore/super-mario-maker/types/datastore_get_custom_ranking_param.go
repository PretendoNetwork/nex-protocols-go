// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCustomRankingParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreGetCustomRankingParam struct {
	nex.Structure
	ApplicationID uint32
	Condition     *DataStoreCustomRankingRatingCondition
	ResultOption  uint8
	ResultRange   *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreGetCustomRankingParam structure from a stream
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetCustomRankingParam.ApplicationID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ApplicationID from stream. %s", err.Error())
	}

	condition, err := stream.ReadStructure(NewDataStoreCustomRankingRatingCondition())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.Condition from stream. %s", err.Error())
	}

	dataStoreGetCustomRankingParam.Condition = condition.(*DataStoreCustomRankingRatingCondition)

	dataStoreGetCustomRankingParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultOption from stream. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingParam.ResultRange from stream. %s", err.Error())
	}

	dataStoreGetCustomRankingParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreGetCustomRankingParam and returns a byte array
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreGetCustomRankingParam.ApplicationID)
	stream.WriteStructure(dataStoreGetCustomRankingParam.Condition)
	stream.WriteUInt8(dataStoreGetCustomRankingParam.ResultOption)
	stream.WriteStructure(dataStoreGetCustomRankingParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetCustomRankingParam
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCustomRankingParam()

	copied.ApplicationID = dataStoreGetCustomRankingParam.ApplicationID
	copied.Condition = dataStoreGetCustomRankingParam.Condition.Copy().(*DataStoreCustomRankingRatingCondition)
	copied.ResultOption = dataStoreGetCustomRankingParam.ResultOption
	copied.ResultRange = dataStoreGetCustomRankingParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCustomRankingParam)

	if dataStoreGetCustomRankingParam.ApplicationID != other.ApplicationID {
		return false
	}

	if !dataStoreGetCustomRankingParam.Condition.Equals(other.Condition) {
		return false
	}

	if dataStoreGetCustomRankingParam.ResultOption != other.ResultOption {
		return false
	}

	if !dataStoreGetCustomRankingParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) String() string {
	return dataStoreGetCustomRankingParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCustomRankingParam *DataStoreGetCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetCustomRankingParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sApplicationID: %d,\n", indentationValues, dataStoreGetCustomRankingParam.ApplicationID))

	if dataStoreGetCustomRankingParam.Condition != nil {
		b.WriteString(fmt.Sprintf("%sCondition: %s\n", indentationValues, dataStoreGetCustomRankingParam.Condition.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCondition: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sResultOption: %d,\n", indentationValues, dataStoreGetCustomRankingParam.ResultOption))

	if dataStoreGetCustomRankingParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, dataStoreGetCustomRankingParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingParam returns a new DataStoreGetCustomRankingParam
func NewDataStoreGetCustomRankingParam() *DataStoreGetCustomRankingParam {
	return &DataStoreGetCustomRankingParam{}
}
