// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreRateObjectParam is sent in the RateObjects method
type DataStoreRateObjectParam struct {
	nex.Structure
	RatingValue    int32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreRateObjectParam structure from a stream
func (dataStoreRateObjectParam *DataStoreRateObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreRateObjectParam.RatingValue, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.RatingValue. %s", err.Error())
	}

	dataStoreRateObjectParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateObjectParam
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateObjectParam()

	copied.SetStructureVersion(dataStoreRateObjectParam.StructureVersion())

	copied.RatingValue = dataStoreRateObjectParam.RatingValue
	copied.AccessPassword = dataStoreRateObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateObjectParam)

	if dataStoreRateObjectParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreRateObjectParam.RatingValue != other.RatingValue {
		return false
	}

	if dataStoreRateObjectParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRateObjectParam *DataStoreRateObjectParam) String() string {
	return dataStoreRateObjectParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRateObjectParam *DataStoreRateObjectParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateObjectParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreRateObjectParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sRatingValue: %d,\n", indentationValues, dataStoreRateObjectParam.RatingValue))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %d\n", indentationValues, dataStoreRateObjectParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateObjectParam returns a new DataStoreRateObjectParam
func NewDataStoreRateObjectParam() *DataStoreRateObjectParam {
	return &DataStoreRateObjectParam{
		RatingValue:    0,
		AccessPassword: 0,
	}
}
