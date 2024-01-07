// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationDownloadMyPokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDownloadMyPokemonResult struct {
	types.Structure
	Result   *GlobalTradeStationDownloadPokemonResult
	IsTraded *types.PrimitiveBool
}

// ExtractFrom extracts the GlobalTradeStationDownloadMyPokemonResult from the given readable
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationDownloadMyPokemonResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationDownloadMyPokemonResult header. %s", err.Error())
	}

	err = globalTradeStationDownloadMyPokemonResult.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.Result from stream. %s", err.Error())
	}

	err = globalTradeStationDownloadMyPokemonResult.IsTraded.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.IsTraded from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationDownloadMyPokemonResult to the given writable
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationDownloadMyPokemonResult.Result.WriteTo(contentWritable)
	globalTradeStationDownloadMyPokemonResult.IsTraded.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationDownloadMyPokemonResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonResult
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadMyPokemonResult()

	copied.StructureVersion = globalTradeStationDownloadMyPokemonResult.StructureVersion

	copied.Result = globalTradeStationDownloadMyPokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.IsTraded = globalTradeStationDownloadMyPokemonResult.IsTraded.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadMyPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadMyPokemonResult)

	if globalTradeStationDownloadMyPokemonResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationDownloadMyPokemonResult.Result.Equals(other.Result) {
		return false
	}

	if !globalTradeStationDownloadMyPokemonResult.IsTraded.Equals(other.IsTraded) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) String() string {
	return globalTradeStationDownloadMyPokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadMyPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationDownloadMyPokemonResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationDownloadMyPokemonResult.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIsTraded: %s,\n", indentationValues, globalTradeStationDownloadMyPokemonResult.IsTraded))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonResult returns a new GlobalTradeStationDownloadMyPokemonResult
func NewGlobalTradeStationDownloadMyPokemonResult() *GlobalTradeStationDownloadMyPokemonResult {
	return &GlobalTradeStationDownloadMyPokemonResult{
		Result: NewGlobalTradeStationDownloadPokemonResult(),
		IsTraded: types.NewPrimitiveBool(false),
	}
}
