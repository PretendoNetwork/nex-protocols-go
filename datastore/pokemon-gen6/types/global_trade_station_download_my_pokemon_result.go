// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationDownloadMyPokemonResult is a type within the DataStore protocol
type GlobalTradeStationDownloadMyPokemonResult struct {
	types.Structure
	Result   GlobalTradeStationDownloadPokemonResult
	IsTraded types.Bool
}

// WriteTo writes the GlobalTradeStationDownloadMyPokemonResult to the given writable
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsdmpr.Result.WriteTo(contentWritable)
	gtsdmpr.IsTraded.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsdmpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationDownloadMyPokemonResult from the given readable
func (gtsdmpr *GlobalTradeStationDownloadMyPokemonResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtsdmpr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult header. %s", err.Error())
	}

	err = gtsdmpr.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.Result. %s", err.Error())
	}

	err = gtsdmpr.IsTraded.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.IsTraded. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonResult
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) Copy() types.RVType {
	copied := NewGlobalTradeStationDownloadMyPokemonResult()

	copied.StructureVersion = gtsdmpr.StructureVersion
	copied.Result = gtsdmpr.Result.Copy().(GlobalTradeStationDownloadPokemonResult)
	copied.IsTraded = gtsdmpr.IsTraded.Copy().(types.Bool)

	return copied
}

// Equals checks if the given GlobalTradeStationDownloadMyPokemonResult contains the same data as the current GlobalTradeStationDownloadMyPokemonResult
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationDownloadMyPokemonResult); !ok {
		return false
	}

	other := o.(*GlobalTradeStationDownloadMyPokemonResult)

	if gtsdmpr.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsdmpr.Result.Equals(other.Result) {
		return false
	}

	return gtsdmpr.IsTraded.Equals(other.IsTraded)
}

// CopyRef copies the current value of the GlobalTradeStationDownloadMyPokemonResult
// and returns a pointer to the new copy
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) CopyRef() types.RVTypePtr {
	copied := gtsdmpr.Copy().(GlobalTradeStationDownloadMyPokemonResult)
	return &copied
}

// Deref takes a pointer to the GlobalTradeStationDownloadMyPokemonResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gtsdmpr *GlobalTradeStationDownloadMyPokemonResult) Deref() types.RVType {
	return *gtsdmpr
}

// String returns the string representation of the GlobalTradeStationDownloadMyPokemonResult
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) String() string {
	return gtsdmpr.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationDownloadMyPokemonResult using the provided indentation level
func (gtsdmpr GlobalTradeStationDownloadMyPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadMyPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, gtsdmpr.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIsTraded: %s,\n", indentationValues, gtsdmpr.IsTraded))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonResult returns a new GlobalTradeStationDownloadMyPokemonResult
func NewGlobalTradeStationDownloadMyPokemonResult() GlobalTradeStationDownloadMyPokemonResult {
	return GlobalTradeStationDownloadMyPokemonResult{
		Result:   NewGlobalTradeStationDownloadPokemonResult(),
		IsTraded: types.NewBool(false),
	}

}
