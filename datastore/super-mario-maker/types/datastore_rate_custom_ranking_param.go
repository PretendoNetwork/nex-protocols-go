// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRateCustomRankingParam is a type within the DataStore protocol
type DataStoreRateCustomRankingParam struct {
	types.Structure
	DataID        types.UInt64
	ApplicationID types.UInt32
	Score         types.UInt32
	Period        types.UInt16
}

// WriteTo writes the DataStoreRateCustomRankingParam to the given writable
func (dsrcrp DataStoreRateCustomRankingParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrcrp.DataID.WriteTo(contentWritable)
	dsrcrp.ApplicationID.WriteTo(contentWritable)
	dsrcrp.Score.WriteTo(contentWritable)
	dsrcrp.Period.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrcrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRateCustomRankingParam from the given readable
func (dsrcrp *DataStoreRateCustomRankingParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrcrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam header. %s", err.Error())
	}

	err = dsrcrp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID. %s", err.Error())
	}

	err = dsrcrp.ApplicationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationID. %s", err.Error())
	}

	err = dsrcrp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score. %s", err.Error())
	}

	err = dsrcrp.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dsrcrp DataStoreRateCustomRankingParam) Copy() types.RVType {
	copied := NewDataStoreRateCustomRankingParam()

	copied.StructureVersion = dsrcrp.StructureVersion
	copied.DataID = dsrcrp.DataID.Copy().(types.UInt64)
	copied.ApplicationID = dsrcrp.ApplicationID.Copy().(types.UInt32)
	copied.Score = dsrcrp.Score.Copy().(types.UInt32)
	copied.Period = dsrcrp.Period.Copy().(types.UInt16)

	return copied
}

// Equals checks if the given DataStoreRateCustomRankingParam contains the same data as the current DataStoreRateCustomRankingParam
func (dsrcrp DataStoreRateCustomRankingParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRateCustomRankingParam); !ok {
		return false
	}

	other := o.(*DataStoreRateCustomRankingParam)

	if dsrcrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrcrp.DataID.Equals(other.DataID) {
		return false
	}

	if !dsrcrp.ApplicationID.Equals(other.ApplicationID) {
		return false
	}

	if !dsrcrp.Score.Equals(other.Score) {
		return false
	}

	return dsrcrp.Period.Equals(other.Period)
}

// CopyRef copies the current value of the DataStoreRateCustomRankingParam
// and returns a pointer to the new copy
func (dsrcrp DataStoreRateCustomRankingParam) CopyRef() types.RVTypePtr {
	copied := dsrcrp.Copy().(DataStoreRateCustomRankingParam)
	return &copied
}

// Deref takes a pointer to the DataStoreRateCustomRankingParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsrcrp *DataStoreRateCustomRankingParam) Deref() types.RVType {
	return *dsrcrp
}

// String returns the string representation of the DataStoreRateCustomRankingParam
func (dsrcrp DataStoreRateCustomRankingParam) String() string {
	return dsrcrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRateCustomRankingParam using the provided indentation level
func (dsrcrp DataStoreRateCustomRankingParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateCustomRankingParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrcrp.DataID))
	b.WriteString(fmt.Sprintf("%sApplicationID: %s,\n", indentationValues, dsrcrp.ApplicationID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dsrcrp.Score))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dsrcrp.Period))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() DataStoreRateCustomRankingParam {
	return DataStoreRateCustomRankingParam{
		DataID:        types.NewUInt64(0),
		ApplicationID: types.NewUInt32(0),
		Score:         types.NewUInt32(0),
		Period:        types.NewUInt16(0),
	}

}
