// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationTradePokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationTradePokemonResult struct {
	types.Structure
	Result   *GlobalTradeStationDownloadPokemonResult
	MyDataID *types.PrimitiveU64
}

// ExtractFrom extracts the GlobalTradeStationTradePokemonResult from the given readable
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationTradePokemonResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationTradePokemonResult header. %s", err.Error())
	}

	err = globalTradeStationTradePokemonResult.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.Result from stream. %s", err.Error())
	}

	err = globalTradeStationTradePokemonResult.MyDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.MyDataID from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationTradePokemonResult to the given writable
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationTradePokemonResult.Result.WriteTo(contentWritable)
	globalTradeStationTradePokemonResult.MyDataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationTradePokemonResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonResult
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationTradePokemonResult()

	copied.StructureVersion = globalTradeStationTradePokemonResult.StructureVersion

	copied.Result = globalTradeStationTradePokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.MyDataID = globalTradeStationTradePokemonResult.MyDataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationTradePokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationTradePokemonResult)

	if globalTradeStationTradePokemonResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationTradePokemonResult.Result.Equals(other.Result) {
		return false
	}

	if !globalTradeStationTradePokemonResult.MyDataID.Equals(other.MyDataID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) String() string {
	return globalTradeStationTradePokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradePokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationTradePokemonResult.StructureVersion))
	b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationTradePokemonResult.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMyDataID: %s,\n", indentationValues, globalTradeStationTradePokemonResult.MyDataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonResult returns a new GlobalTradeStationTradePokemonResult
func NewGlobalTradeStationTradePokemonResult() *GlobalTradeStationTradePokemonResult {
	return &GlobalTradeStationTradePokemonResult{
		Result: NewGlobalTradeStationDownloadPokemonResult(),
		MyDataID: types.NewPrimitiveU64(0),
	}
}
