// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationPrepareTradePokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationPrepareTradePokemonParam struct {
	types.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFrom extracts the GlobalTradeStationPrepareTradePokemonParam from the given readable
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationPrepareTradePokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationPrepareTradePokemonParam header. %s", err.Error())
	}

	err = globalTradeStationPrepareTradePokemonParam.TradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.TradeKey from stream. %s", err.Error())
	}

	err = globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationPrepareTradePokemonParam to the given writable
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationPrepareTradePokemonParam.TradeKey.WriteTo(contentWritable)
	globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationPrepareTradePokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonParam
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationPrepareTradePokemonParam()

	copied.StructureVersion = globalTradeStationPrepareTradePokemonParam.StructureVersion

	copied.TradeKey = globalTradeStationPrepareTradePokemonParam.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareUploadKey = globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationPrepareTradePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationPrepareTradePokemonParam)

	if globalTradeStationPrepareTradePokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationPrepareTradePokemonParam.TradeKey.Equals(other.TradeKey) {
		return false
	}

	if !globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) String() string {
	return globalTradeStationPrepareTradePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationPrepareTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationPrepareTradePokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTradeKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonParam.TradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonParam returns a new GlobalTradeStationPrepareTradePokemonParam
func NewGlobalTradeStationPrepareTradePokemonParam() *GlobalTradeStationPrepareTradePokemonParam {
	return &GlobalTradeStationPrepareTradePokemonParam{
		TradeKey: NewGlobalTradeStationTradeKey(),
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
	}
}
