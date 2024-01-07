// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationSearchPokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationSearchPokemonParam struct {
	types.Structure
	PrepareUploadKey  *GlobalTradeStationRecordKey
	Conditions        *types.List[*types.PrimitiveU32]
	ResultOrderColumn *types.PrimitiveU8
	ResultOrder       *types.PrimitiveU8
	UploadedAfter     *types.DateTime
	UploadedBefore    *types.DateTime
	ResultRange       *types.ResultRange
}

// ExtractFrom extracts the GlobalTradeStationSearchPokemonParam from the given readable
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationSearchPokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationSearchPokemonParam header. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.Conditions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.Conditions from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.ResultOrderColumn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrderColumn from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.ResultOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrder from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.UploadedAfter.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedAfter from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.UploadedBefore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedBefore from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultRange from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationSearchPokemonParam to the given writable
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationSearchPokemonParam.PrepareUploadKey.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.Conditions.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.ResultOrderColumn.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.ResultOrder.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.UploadedAfter.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.UploadedBefore.WriteTo(contentWritable)
	globalTradeStationSearchPokemonParam.ResultRange.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationSearchPokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonParam
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationSearchPokemonParam()

	copied.StructureVersion = globalTradeStationSearchPokemonParam.StructureVersion

	copied.PrepareUploadKey = globalTradeStationSearchPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Conditions = globalTradeStationSearchPokemonParam.Conditions.Copy().(*types.List[*types.PrimitiveU32])
	copied.ResultOrderColumn = globalTradeStationSearchPokemonParam.ResultOrderColumn.Copy().(*types.PrimitiveU8)
	copied.ResultOrder = globalTradeStationSearchPokemonParam.ResultOrder.Copy().(*types.PrimitiveU8)
	copied.UploadedAfter = globalTradeStationSearchPokemonParam.UploadedAfter.Copy().(*types.DateTime)
	copied.UploadedBefore = globalTradeStationSearchPokemonParam.UploadedBefore.Copy().(*types.DateTime)
	copied.ResultRange = globalTradeStationSearchPokemonParam.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationSearchPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationSearchPokemonParam)

	if globalTradeStationSearchPokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationSearchPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.Conditions.Equals(other.Conditions) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.ResultOrderColumn.Equals(other.ResultOrderColumn) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.ResultOrder.Equals(other.ResultOrder) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.UploadedAfter.Equals(other.UploadedAfter) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.UploadedBefore.Equals(other.UploadedBefore) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) String() string {
	return globalTradeStationSearchPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationSearchPokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationSearchPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sConditions: %s,\n", indentationValues, globalTradeStationSearchPokemonParam.Conditions))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %s,\n", indentationValues, globalTradeStationSearchPokemonParam.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %s,\n", indentationValues, globalTradeStationSearchPokemonParam.ResultOrder))
	b.WriteString(fmt.Sprintf("%sUploadedAfter: %s\n", indentationValues, globalTradeStationSearchPokemonParam.UploadedAfter.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUploadedBefore: %s\n", indentationValues, globalTradeStationSearchPokemonParam.UploadedBefore.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, globalTradeStationSearchPokemonParam.ResultRange.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonParam returns a new GlobalTradeStationSearchPokemonParam
func NewGlobalTradeStationSearchPokemonParam() *GlobalTradeStationSearchPokemonParam {
	globalTradeStationSearchPokemonParam := &GlobalTradeStationSearchPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		Conditions: types.NewList[*types.PrimitiveU32](),
		ResultOrderColumn: types.NewPrimitiveU8(0),
		ResultOrder: types.NewPrimitiveU8(0),
		UploadedAfter: types.NewDateTime(0),
		UploadedBefore: types.NewDateTime(0),
		ResultRange: types.NewResultRange(),
	}

	globalTradeStationSearchPokemonParam.Conditions.Type = types.NewPrimitiveU32(0)

	return globalTradeStationSearchPokemonParam
}
