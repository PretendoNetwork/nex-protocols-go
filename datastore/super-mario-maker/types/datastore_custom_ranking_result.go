// Package types implements all the types used by the DataStore (Super Mario Maker) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

// DataStoreCustomRankingResult holds data for the DataStore (Super Mario Maker) protocol
type DataStoreCustomRankingResult struct {
	types.Structure
	Order    *types.PrimitiveU32
	Score    *types.PrimitiveU32
	MetaInfo *datastore_types.DataStoreMetaInfo
}

// ExtractFrom extracts the DataStoreCustomRankingResult from the given readable
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreCustomRankingResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCustomRankingResult header. %s", err.Error())
	}

	err = dataStoreCustomRankingResult.Order.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Order from stream. %s", err.Error())
	}

	err = dataStoreCustomRankingResult.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Score from stream. %s", err.Error())
	}

	err = dataStoreCustomRankingResult.MetaInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.MetaInfo from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreCustomRankingResult to the given writable
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreCustomRankingResult.Order.WriteTo(contentWritable)
	dataStoreCustomRankingResult.Score.WriteTo(contentWritable)
	dataStoreCustomRankingResult.MetaInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCustomRankingResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCustomRankingResult
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Copy() types.RVType {
	copied := NewDataStoreCustomRankingResult()

	copied.StructureVersion = dataStoreCustomRankingResult.StructureVersion

	copied.Order = dataStoreCustomRankingResult.Order.Copy().(*types.PrimitiveU32)
	copied.Score = dataStoreCustomRankingResult.Score.Copy().(*types.PrimitiveU32)
	copied.MetaInfo = dataStoreCustomRankingResult.MetaInfo.Copy().(*datastore_types.DataStoreMetaInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCustomRankingResult); !ok {
		return false
	}

	other := o.(*DataStoreCustomRankingResult)

	if dataStoreCustomRankingResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCustomRankingResult.Order.Equals(other.Order) {
		return false
	}

	if !dataStoreCustomRankingResult.Score.Equals(other.Score) {
		return false
	}

	if !dataStoreCustomRankingResult.MetaInfo.Equals(other.MetaInfo) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) String() string {
	return dataStoreCustomRankingResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCustomRankingResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCustomRankingResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOrder: %s,\n", indentationValues, dataStoreCustomRankingResult.Order))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStoreCustomRankingResult.Score))
	b.WriteString(fmt.Sprintf("%sMetaInfo: %s\n", indentationValues, dataStoreCustomRankingResult.MetaInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{
		Order:    types.NewPrimitiveU32(0),
		Score:    types.NewPrimitiveU32(0),
		MetaInfo: datastore_types.NewDataStoreMetaInfo(),
	}
}
