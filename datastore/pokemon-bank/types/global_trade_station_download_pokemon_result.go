// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationDownloadPokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDownloadPokemonResult struct {
	types.Structure
	DataID      *types.PrimitiveU64
	IndexData   *types.QBuffer
	PokemonData *types.QBuffer
}

// ExtractFrom extracts the GlobalTradeStationDownloadPokemonResult from the given readable
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationDownloadPokemonResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationDownloadPokemonResult header. %s", err.Error())
	}

	err = globalTradeStationDownloadPokemonResult.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.DataID from stream. %s", err.Error())
	}

	err = globalTradeStationDownloadPokemonResult.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.IndexData from stream. %s", err.Error())
	}

	err = globalTradeStationDownloadPokemonResult.PokemonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.PokemonData from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationDownloadPokemonResult to the given writable
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationDownloadPokemonResult.DataID.WriteTo(contentWritable)
	globalTradeStationDownloadPokemonResult.IndexData.WriteTo(contentWritable)
	globalTradeStationDownloadPokemonResult.PokemonData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationDownloadPokemonResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationDownloadPokemonResult
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadPokemonResult()

	copied.StructureVersion = globalTradeStationDownloadPokemonResult.StructureVersion

	copied.DataID = globalTradeStationDownloadPokemonResult.DataID.Copy().(*types.PrimitiveU64)
	copied.IndexData = globalTradeStationDownloadPokemonResult.IndexData.Copy().(*types.QBuffer)
	copied.PokemonData = globalTradeStationDownloadPokemonResult.PokemonData.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadPokemonResult)

	if globalTradeStationDownloadPokemonResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationDownloadPokemonResult.DataID.Equals(other.DataID) {
		return false
	}

	if !globalTradeStationDownloadPokemonResult.IndexData.Equals(other.IndexData) {
		return false
	}

	if !globalTradeStationDownloadPokemonResult.PokemonData.Equals(other.PokemonData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) String() string {
	return globalTradeStationDownloadPokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationDownloadPokemonResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, globalTradeStationDownloadPokemonResult.DataID))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, globalTradeStationDownloadPokemonResult.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %s,\n", indentationValues, globalTradeStationDownloadPokemonResult.PokemonData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadPokemonResult returns a new GlobalTradeStationDownloadPokemonResult
func NewGlobalTradeStationDownloadPokemonResult() *GlobalTradeStationDownloadPokemonResult {
	return &GlobalTradeStationDownloadPokemonResult{
		DataID: types.NewPrimitiveU64(0),
		IndexData: types.NewQBuffer(nil),
		PokemonData: types.NewQBuffer(nil),
	}
}
