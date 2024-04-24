// Package types implements all the types used by the DataStoreACHappyHomeDesigner protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreSearchHouseParam is a type within the DataStoreACHappyHomeDesigner protocol
type DataStoreSearchHouseParam struct {
	types.Structure

	DataType           *types.PrimitiveU16
	ResultOrderColumns *types.Buffer
	ResultRange        *types.ResultRange
	ResultOption       *types.PrimitiveU8
	Region             *types.PrimitiveU8
	Country            *types.PrimitiveU8
}

// WriteTo writes the DataStoreSearchHouseParam to the given variable
func (dsshp *DataStoreSearchHouseParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsshp.DataType.WriteTo(contentWritable)
	dsshp.ResultOrderColumns.WriteTo(contentWritable)
	dsshp.ResultRange.WriteTo(contentWritable)
	dsshp.ResultOption.WriteTo(contentWritable)
	dsshp.Region.WriteTo(contentWritable)
	dsshp.Country.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsshp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreSearchHouseParam from the given readable
func (dsshp *DataStoreSearchHouseParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsshp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam header. %s", err.Error())
	}

	err = dsshp.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.DataType. %s", err.Error())
	}

	err = dsshp.ResultOrderColumns.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.ResultOrderColumns. %s", err.Error())
	}

	err = dsshp.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.ResultRange. %s", err.Error())
	}

	err = dsshp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.ResultOption. %s", err.Error())
	}

	err = dsshp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.Region. %s", err.Error())
	}

	err = dsshp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSearchHouseParam.Country. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchHouseParam
func (dsshp *DataStoreSearchHouseParam) Copy() types.RVType {
	copied := NewDataStoreSearchHouseParam()

	copied.DataType = dsshp.DataType
	copied.ResultOrderColumns = dsshp.ResultOrderColumns
	copied.ResultRange = dsshp.ResultRange
	copied.ResultOption = dsshp.ResultOption
	copied.Region = dsshp.Region
	copied.Country = dsshp.Country

	return copied
}

// Equals checks if the given DataStoreSearchHouseParam contains the same data as the current DataStoreSearchHouseParam
func (dsshp *DataStoreSearchHouseParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSearchHouseParam); !ok {
		return false
	}

	other := o.(*DataStoreSearchHouseParam)

	if !dsshp.DataType.Equals(other.DataType) {
		return false
	}

	if !dsshp.ResultOrderColumns.Equals(other.ResultOrderColumns) {
		return false
	}

	if !dsshp.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if !dsshp.ResultOption.Equals(other.ResultOption) {
		return false
	}

	if !dsshp.Region.Equals(other.Region) {
		return false
	}

	return dsshp.Country.Equals(other.Country)
}

// String returns the string representation of the DataStoreSearchHouseParam
func (dsshp *DataStoreSearchHouseParam) String() string {
	return dsshp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreSearchHouseParam using the provided indentation level

func (dsshp *DataStoreSearchHouseParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSearchHouseParam{\n")
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dsshp.DataType))
	b.WriteString(fmt.Sprintf("%sResultOrderColumns: %s,\n", indentationValues, dsshp.ResultOrderColumns))
	b.WriteString(fmt.Sprintf("%sResultRange: %s,\n", indentationValues, dsshp.ResultRange))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, dsshp.ResultOption))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, dsshp.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, dsshp.Country))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSearchHouseParam returns a new DataStoreSearchHouseParam
func NewDataStoreSearchHouseParam() *DataStoreSearchHouseParam {
	dsshp := &DataStoreSearchHouseParam{
		DataType:           types.NewPrimitiveU16(0),
		ResultOrderColumns: types.NewBuffer(nil),
		ResultRange:        types.NewResultRange(),
		ResultOption:       types.NewPrimitiveU8(0),
		Region:             types.NewPrimitiveU8(0),
		Country:            types.NewPrimitiveU8(0),
	}

	return dsshp
}
