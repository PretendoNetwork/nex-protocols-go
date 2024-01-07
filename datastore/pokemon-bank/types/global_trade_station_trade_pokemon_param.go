// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationTradePokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationTradePokemonParam struct {
	types.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareTradeKey  *GlobalTradeStationRecordKey
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           *types.PrimitiveU16
	IndexData        *types.QBuffer
	PokemonData      *types.QBuffer
	Signature        *types.QBuffer
	NeedData         *types.PrimitiveBool
}

// ExtractFrom extracts the GlobalTradeStationTradePokemonParam from the given readable
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationTradePokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationTradePokemonParam header. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.TradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.TradeKey from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.PrepareTradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareTradeKey from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Period from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.IndexData from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PokemonData from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.Signature.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Signature from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonParam.NeedData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.NeedData from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationTradePokemonParam to the given writable
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationTradePokemonParam.TradeKey.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.PrepareTradeKey.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.PrepareUploadKey.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.Period.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.IndexData.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.PokemonData.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.Signature.WriteTo(contentWritable)
	globalTradeStationTradePokemonParam.NeedData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationTradePokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonParam
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationTradePokemonParam()

	copied.StructureVersion = globalTradeStationTradePokemonParam.StructureVersion

	copied.TradeKey = globalTradeStationTradePokemonParam.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareTradeKey = globalTradeStationTradePokemonParam.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)
	copied.PrepareUploadKey = globalTradeStationTradePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = globalTradeStationTradePokemonParam.Period.Copy().(*types.PrimitiveU16)
	copied.IndexData = globalTradeStationTradePokemonParam.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = globalTradeStationTradePokemonParam.PokemonData.Copy().(*types.QBuffer)
	copied.Signature = globalTradeStationTradePokemonParam.Signature.Copy().(*types.QBuffer)
	copied.NeedData = globalTradeStationTradePokemonParam.NeedData.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationTradePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationTradePokemonParam)

	if globalTradeStationTradePokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationTradePokemonParam.TradeKey.Equals(other.TradeKey) {
		return false
	}

	if !globalTradeStationTradePokemonParam.PrepareTradeKey.Equals(other.PrepareTradeKey) {
		return false
	}

	if !globalTradeStationTradePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !globalTradeStationTradePokemonParam.Period.Equals(other.Period) {
		return false
	}

	if !globalTradeStationTradePokemonParam.IndexData.Equals(other.IndexData) {
		return false
	}

	if !globalTradeStationTradePokemonParam.PokemonData.Equals(other.PokemonData) {
		return false
	}

	if !globalTradeStationTradePokemonParam.Signature.Equals(other.Signature) {
		return false
	}

	if !globalTradeStationTradePokemonParam.NeedData.Equals(other.NeedData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) String() string {
	return globalTradeStationTradePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationTradePokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTradeKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.TradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.PrepareTradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, globalTradeStationTradePokemonParam.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, globalTradeStationTradePokemonParam.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, globalTradeStationTradePokemonParam.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %s,\n", indentationValues, globalTradeStationTradePokemonParam.Signature))
	b.WriteString(fmt.Sprintf("%sNeedData: %s,\n", indentationValues, globalTradeStationTradePokemonParam.NeedData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonParam returns a new GlobalTradeStationTradePokemonParam
func NewGlobalTradeStationTradePokemonParam() *GlobalTradeStationTradePokemonParam {
	return &GlobalTradeStationTradePokemonParam{
		TradeKey: NewGlobalTradeStationTradeKey(),
		PrepareTradeKey: NewGlobalTradeStationRecordKey(),
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		Period: types.NewPrimitiveU16(0),
		IndexData: types.NewQBuffer(nil),
		PokemonData: types.NewQBuffer(nil),
		Signature: types.NewQBuffer(nil),
		NeedData: types.NewPrimitiveBool(false),
	}
}
