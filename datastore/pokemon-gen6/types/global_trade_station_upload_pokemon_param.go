// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationUploadPokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationUploadPokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           *types.PrimitiveU16
	IndexData        *types.QBuffer
	PokemonData      *types.QBuffer
	Signature        *types.QBuffer
}

// ExtractFrom extracts the GlobalTradeStationUploadPokemonParam from the given readable
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationUploadPokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationUploadPokemonParam header. %s", err.Error())
	}

	err = globalTradeStationUploadPokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	err = globalTradeStationUploadPokemonParam.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Period from stream. %s", err.Error())
	}

	err = globalTradeStationUploadPokemonParam.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.IndexData from stream. %s", err.Error())
	}

	err = globalTradeStationUploadPokemonParam.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PokemonData from stream. %s", err.Error())
	}

	err = globalTradeStationUploadPokemonParam.Signature.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Signature from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationUploadPokemonParam to the given writable
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationUploadPokemonParam.PrepareUploadKey.WriteTo(contentWritable)
	globalTradeStationUploadPokemonParam.Period.WriteTo(contentWritable)
	globalTradeStationUploadPokemonParam.IndexData.WriteTo(contentWritable)
	globalTradeStationUploadPokemonParam.PokemonData.WriteTo(contentWritable)
	globalTradeStationUploadPokemonParam.Signature.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationUploadPokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationUploadPokemonParam
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationUploadPokemonParam()

	copied.StructureVersion = globalTradeStationUploadPokemonParam.StructureVersion

	copied.PrepareUploadKey = globalTradeStationUploadPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = globalTradeStationUploadPokemonParam.Period.Copy().(*types.PrimitiveU16)
	copied.IndexData = globalTradeStationUploadPokemonParam.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = globalTradeStationUploadPokemonParam.PokemonData.Copy().(*types.QBuffer)
	copied.Signature = globalTradeStationUploadPokemonParam.Signature.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationUploadPokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationUploadPokemonParam)

	if globalTradeStationUploadPokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationUploadPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !globalTradeStationUploadPokemonParam.Period.Equals(other.Period) {
		return false
	}

	if !globalTradeStationUploadPokemonParam.IndexData.Equals(other.IndexData) {
		return false
	}

	if !globalTradeStationUploadPokemonParam.PokemonData.Equals(other.PokemonData) {
		return false
	}

	if !globalTradeStationUploadPokemonParam.Signature.Equals(other.Signature) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) String() string {
	return globalTradeStationUploadPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationUploadPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationUploadPokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationUploadPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, globalTradeStationUploadPokemonParam.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, globalTradeStationUploadPokemonParam.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, globalTradeStationUploadPokemonParam.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %s,\n", indentationValues, globalTradeStationUploadPokemonParam.Signature))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationUploadPokemonParam returns a new GlobalTradeStationUploadPokemonParam
func NewGlobalTradeStationUploadPokemonParam() *GlobalTradeStationUploadPokemonParam {
	return &GlobalTradeStationUploadPokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		Period: types.NewPrimitiveU16(0),
		IndexData: types.NewQBuffer(nil),
		PokemonData: types.NewQBuffer(nil),
		Signature: types.NewQBuffer(nil),
	}
}
