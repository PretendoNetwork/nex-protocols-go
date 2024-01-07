// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRateObjectParam is sent in the RateObjects method
type DataStoreRateObjectParam struct {
	types.Structure
	RatingValue    *types.PrimitiveS32
	AccessPassword *types.PrimitiveU64
}

// WriteTo writes the DataStoreRateObjectParam to the given writable
func (dataStoreRateObjectParam *DataStoreRateObjectParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRateObjectParam.RatingValue.WriteTo(contentWritable)
	dataStoreRateObjectParam.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRateObjectParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRateObjectParam from the given readable
func (dataStoreRateObjectParam *DataStoreRateObjectParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRateObjectParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRateObjectParam header. %s", err.Error())
	}

	err = dataStoreRateObjectParam.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.RatingValue. %s", err.Error())
	}

	err = dataStoreRateObjectParam.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateObjectParam
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Copy() types.RVType {
	copied := NewDataStoreRateObjectParam()

	copied.StructureVersion = dataStoreRateObjectParam.StructureVersion

	copied.RatingValue = dataStoreRateObjectParam.RatingValue.Copy().(*types.PrimitiveS32)
	copied.AccessPassword = dataStoreRateObjectParam.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRateObjectParam); !ok {
		return false
	}

	other := o.(*DataStoreRateObjectParam)

	if dataStoreRateObjectParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRateObjectParam.RatingValue.Equals(other.RatingValue) {
		return false
	}

	if !dataStoreRateObjectParam.AccessPassword.Equals(other.AccessPassword) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRateObjectParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sRatingValue: %s,\n", indentationValues, dataStoreRateObjectParam.RatingValue))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s\n", indentationValues, dataStoreRateObjectParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateObjectParam returns a new DataStoreRateObjectParam
func NewDataStoreRateObjectParam() *DataStoreRateObjectParam {
	return &DataStoreRateObjectParam{
		RatingValue:    types.NewPrimitiveS32(0),
		AccessPassword: types.NewPrimitiveU64(0),
	}
}
