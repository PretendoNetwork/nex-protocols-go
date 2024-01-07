// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationDeletePokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationDeletePokemonParam struct {
	types.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	DeleteFlag       *types.PrimitiveU8
}

// ExtractFrom extracts the GlobalTradeStationDeletePokemonParam from the given readable
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationDeletePokemonParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationDeletePokemonParam header. %s", err.Error())
	}

	err = globalTradeStationDeletePokemonParam.PrepareUploadKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	err = globalTradeStationDeletePokemonParam.DeleteFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.DeleteFlag from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationDeletePokemonParam to the given writable
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationDeletePokemonParam.PrepareUploadKey.WriteTo(contentWritable)
	globalTradeStationDeletePokemonParam.DeleteFlag.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationDeletePokemonParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationDeletePokemonParam
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) Copy() types.RVType {
	copied := NewGlobalTradeStationDeletePokemonParam()

	copied.StructureVersion = globalTradeStationDeletePokemonParam.StructureVersion

	copied.PrepareUploadKey = globalTradeStationDeletePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.DeleteFlag = globalTradeStationDeletePokemonParam.DeleteFlag.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDeletePokemonParam); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDeletePokemonParam)

	if globalTradeStationDeletePokemonParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationDeletePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if !globalTradeStationDeletePokemonParam.DeleteFlag.Equals(other.DeleteFlag) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) String() string {
	return globalTradeStationDeletePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDeletePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationDeletePokemonParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDeletePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDeleteFlag: %s,\n", indentationValues, globalTradeStationDeletePokemonParam.DeleteFlag))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDeletePokemonParam returns a new GlobalTradeStationDeletePokemonParam
func NewGlobalTradeStationDeletePokemonParam() *GlobalTradeStationDeletePokemonParam {
	return &GlobalTradeStationDeletePokemonParam{
		PrepareUploadKey: NewGlobalTradeStationRecordKey(),
		DeleteFlag: types.NewPrimitiveU8(0),
	}
}
