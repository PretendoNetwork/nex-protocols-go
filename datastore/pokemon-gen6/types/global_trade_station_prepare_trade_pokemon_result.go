// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationPrepareTradePokemonResult holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationPrepareTradePokemonResult struct {
	types.Structure
	Result          *GlobalTradeStationDownloadPokemonResult
	PrepareTradeKey *GlobalTradeStationRecordKey
}

// ExtractFrom extracts the GlobalTradeStationPrepareTradePokemonResult from the given readable
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationPrepareTradePokemonResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationPrepareTradePokemonResult header. %s", err.Error())
	}

	err = globalTradeStationPrepareTradePokemonResult.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.Result from stream. %s", err.Error())
	}

	err = globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.PrepareTradeKey from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationPrepareTradePokemonResult to the given writable
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationPrepareTradePokemonResult.Result.WriteTo(contentWritable)
	globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationPrepareTradePokemonResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonResult
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationPrepareTradePokemonResult()

	copied.StructureVersion = globalTradeStationPrepareTradePokemonResult.StructureVersion

	copied.Result = globalTradeStationPrepareTradePokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.PrepareTradeKey = globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationPrepareTradePokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationPrepareTradePokemonResult)

	if globalTradeStationPrepareTradePokemonResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationPrepareTradePokemonResult.Result.Equals(other.Result) {
		return false
	}

	if !globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.Equals(other.PrepareTradeKey) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) String() string {
	return globalTradeStationPrepareTradePokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationPrepareTradePokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationPrepareTradePokemonResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationPrepareTradePokemonResult.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonResult returns a new GlobalTradeStationPrepareTradePokemonResult
func NewGlobalTradeStationPrepareTradePokemonResult() *GlobalTradeStationPrepareTradePokemonResult {
	return &GlobalTradeStationPrepareTradePokemonResult{
		Result: NewGlobalTradeStationDownloadPokemonResult(),
		PrepareTradeKey: NewGlobalTradeStationRecordKey(),
	}
}
