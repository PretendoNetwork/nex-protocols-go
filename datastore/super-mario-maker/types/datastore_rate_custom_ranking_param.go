// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreRateCustomRankingParam holds data for the DataStore (Super Mario Maker) protocol
type DataStoreRateCustomRankingParam struct {
	types.Structure
	DataID        *types.PrimitiveU64
	ApplicationID *types.PrimitiveU32
	Score         *types.PrimitiveU32
	Period        *types.PrimitiveU16
}

// ExtractFrom extracts the DataStoreRateCustomRankingParam from the given readable
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreRateCustomRankingParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreRateCustomRankingParam header. %s", err.Error())
	}

	err = dataStoreRateCustomRankingParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID from stream. %s", err.Error())
	}

	err = dataStoreRateCustomRankingParam.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationID from stream. %s", err.Error())
	}

	err = dataStoreRateCustomRankingParam.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score from stream. %s", err.Error())
	}

	err = dataStoreRateCustomRankingParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreRateCustomRankingParam to the given writable
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreRateCustomRankingParam.DataID.WriteTo(contentWritable)
	dataStoreRateCustomRankingParam.ApplicationID.WriteTo(contentWritable)
	dataStoreRateCustomRankingParam.Score.WriteTo(contentWritable)
	dataStoreRateCustomRankingParam.Period.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreRateCustomRankingParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Copy() types.RVType {
	copied := NewDataStoreRateCustomRankingParam()

	copied.StructureVersion = dataStoreRateCustomRankingParam.StructureVersion

	copied.DataID = dataStoreRateCustomRankingParam.DataID.Copy().(*types.PrimitiveU64)
	copied.ApplicationID = dataStoreRateCustomRankingParam.ApplicationID.Copy().(*types.PrimitiveU32)
	copied.Score = dataStoreRateCustomRankingParam.Score.Copy().(*types.PrimitiveU32)
	copied.Period = dataStoreRateCustomRankingParam.Period.Copy().(*types.PrimitiveU16)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRateCustomRankingParam); !ok {
		return false
	}

	other := o.(*DataStoreRateCustomRankingParam)

	if dataStoreRateCustomRankingParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreRateCustomRankingParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreRateCustomRankingParam.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dataStoreRateCustomRankingParam.Score.Equals(other.Score) {
		return false
	}

	if !dataStoreRateCustomRankingParam.Period.Equals(other.Period) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) String() string {
	return dataStoreRateCustomRankingParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreRateCustomRankingParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreRateCustomRankingParam.DataID))
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dataStoreRateCustomRankingParam.ApplicationID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStoreRateCustomRankingParam.Score))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dataStoreRateCustomRankingParam.Period))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	return &DataStoreRateCustomRankingParam{
		DataID:        types.NewPrimitiveU64(0),
		ApplicationID: types.NewPrimitiveU32(0),
		Score:         types.NewPrimitiveU32(0),
		Period:        types.NewPrimitiveU16(0),
	}
}
