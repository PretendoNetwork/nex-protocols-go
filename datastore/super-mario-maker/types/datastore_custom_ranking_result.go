// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// DataStoreCustomRankingResult is a type within the DataStore protocol
type DataStoreCustomRankingResult struct {
	types.Structure
	Order    types.UInt32
	Score    types.UInt32
	MetaInfo datastore_types.DataStoreMetaInfo
}

// WriteTo writes the DataStoreCustomRankingResult to the given writable
func (dscrr DataStoreCustomRankingResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscrr.Order.WriteTo(contentWritable)
	dscrr.Score.WriteTo(contentWritable)
	dscrr.MetaInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscrr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCustomRankingResult from the given readable
func (dscrr *DataStoreCustomRankingResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscrr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult header. %s", err.Error())
	}

	err = dscrr.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Order. %s", err.Error())
	}

	err = dscrr.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Score. %s", err.Error())
	}

	err = dscrr.MetaInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.MetaInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCustomRankingResult
func (dscrr DataStoreCustomRankingResult) Copy() types.RVType {
	copied := NewDataStoreCustomRankingResult()

	copied.StructureVersion = dscrr.StructureVersion
	copied.Order = dscrr.Order.Copy().(types.UInt32)
	copied.Score = dscrr.Score.Copy().(types.UInt32)
	copied.MetaInfo = dscrr.MetaInfo.Copy().(datastore_types.DataStoreMetaInfo)

	return copied
}

// Equals checks if the given DataStoreCustomRankingResult contains the same data as the current DataStoreCustomRankingResult
func (dscrr DataStoreCustomRankingResult) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreCustomRankingResult); !ok {
		return false
	}

	other := o.(DataStoreCustomRankingResult)

	if dscrr.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscrr.Order.Equals(other.Order) {
		return false
	}

	if !dscrr.Score.Equals(other.Score) {
		return false
	}

	return dscrr.MetaInfo.Equals(other.MetaInfo)
}

// CopyRef copies the current value of the DataStoreCustomRankingResult
// and returns a pointer to the new copy
func (dscrr DataStoreCustomRankingResult) CopyRef() types.RVTypePtr {
	copied := dscrr.Copy().(DataStoreCustomRankingResult)
	return &copied
}

// Deref takes a pointer to the DataStoreCustomRankingResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dscrr *DataStoreCustomRankingResult) Deref() types.RVType {
	return *dscrr
}

// String returns the string representation of the DataStoreCustomRankingResult
func (dscrr DataStoreCustomRankingResult) String() string {
	return dscrr.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCustomRankingResult using the provided indentation level
func (dscrr DataStoreCustomRankingResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCustomRankingResult{\n")
	b.WriteString(fmt.Sprintf("%sOrder: %s,\n", indentationValues, dscrr.Order))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dscrr.Score))
	b.WriteString(fmt.Sprintf("%sMetaInfo: %s,\n", indentationValues, dscrr.MetaInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() DataStoreCustomRankingResult {
	return DataStoreCustomRankingResult{
		Order:    types.NewUInt32(0),
		Score:    types.NewUInt32(0),
		MetaInfo: datastore_types.NewDataStoreMetaInfo(),
	}

}
