// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationSearchPokemonResult holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationSearchPokemonResult struct {
	types.Structure
	TotalCount     *types.PrimitiveU32
	Result         *types.List[*GlobalTradeStationData]
	TotalCountType *types.PrimitiveU8
}

// ExtractFrom extracts the GlobalTradeStationSearchPokemonResult from the given readable
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationSearchPokemonResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationSearchPokemonResult header. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonResult.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCount from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonResult.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.Result from stream. %s", err.Error())
	}

	err = globalTradeStationSearchPokemonResult.TotalCountType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCountType from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationSearchPokemonResult to the given writable
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationSearchPokemonResult.TotalCount.WriteTo(contentWritable)
	globalTradeStationSearchPokemonResult.Result.WriteTo(contentWritable)
	globalTradeStationSearchPokemonResult.TotalCountType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationSearchPokemonResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonResult
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationSearchPokemonResult()

	copied.StructureVersion = globalTradeStationSearchPokemonResult.StructureVersion

	copied.TotalCount = globalTradeStationSearchPokemonResult.TotalCount.Copy().(*types.PrimitiveU32)
	copied.Result = globalTradeStationSearchPokemonResult.Result.Copy().(*types.List[*GlobalTradeStationData])
	copied.TotalCountType = globalTradeStationSearchPokemonResult.TotalCountType.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationSearchPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationSearchPokemonResult)

	if globalTradeStationSearchPokemonResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationSearchPokemonResult.TotalCount.Equals(other.TotalCount) {
		return false
	}

	if !globalTradeStationSearchPokemonResult.Result.Equals(other.Result) {
		return false
	}

	return globalTradeStationSearchPokemonResult.TotalCountType.Equals(other.TotalCountType)
}

// String returns a string representation of the struct
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) String() string {
	return globalTradeStationSearchPokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationSearchPokemonResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTotalCount: %s,\n", indentationValues, globalTradeStationSearchPokemonResult.TotalCount))
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, globalTradeStationSearchPokemonResult.Result))
	b.WriteString(fmt.Sprintf("%sTotalCountType: %s,\n", indentationValues, globalTradeStationSearchPokemonResult.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonResult returns a new GlobalTradeStationSearchPokemonResult
func NewGlobalTradeStationSearchPokemonResult() *GlobalTradeStationSearchPokemonResult {
	globalTradeStationSearchPokemonResult := &GlobalTradeStationSearchPokemonResult{
		TotalCount: types.NewPrimitiveU32(0),
		Result: types.NewList[*GlobalTradeStationData](),
		TotalCountType: types.NewPrimitiveU8(0),
	}

	globalTradeStationSearchPokemonResult.Result.Type = NewGlobalTradeStationData()

	return globalTradeStationSearchPokemonResult
}
