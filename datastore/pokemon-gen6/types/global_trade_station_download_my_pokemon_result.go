// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationDownloadMyPokemonResult holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationDownloadMyPokemonResult struct {
	nex.Structure
	Result   *GlobalTradeStationDownloadPokemonResult
	IsTraded bool
}

// ExtractFromStream extracts a GlobalTradeStationDownloadMyPokemonResult structure from a stream
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	result, err := stream.ReadStructure(NewGlobalTradeStationDownloadPokemonResult())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.Result from stream. %s", err.Error())
	}

	globalTradeStationDownloadMyPokemonResult.Result = result.(*GlobalTradeStationDownloadPokemonResult)

	globalTradeStationDownloadMyPokemonResult.IsTraded, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonResult.IsTraded from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationDownloadMyPokemonResult and returns a byte array
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationDownloadMyPokemonResult.Result)
	stream.WriteBool(globalTradeStationDownloadMyPokemonResult.IsTraded)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonResult
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationDownloadMyPokemonResult()

	copied.Result = globalTradeStationDownloadMyPokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.IsTraded = globalTradeStationDownloadMyPokemonResult.IsTraded

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadMyPokemonResult *GlobalTradeStationDownloadMyPokemonResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationDownloadMyPokemonResult)

	if !globalTradeStationDownloadMyPokemonResult.Result.Equals(other.Result) {
		return false
	}

	if globalTradeStationDownloadMyPokemonResult.IsTraded != other.IsTraded {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationDownloadMyPokemonResult.StructureVersion()))

	if globalTradeStationDownloadMyPokemonResult.Result != nil {
		b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationDownloadMyPokemonResult.Result.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sIsTraded: %t,\n", indentationValues, globalTradeStationDownloadMyPokemonResult.IsTraded))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonResult returns a new GlobalTradeStationDownloadMyPokemonResult
func NewGlobalTradeStationDownloadMyPokemonResult() *GlobalTradeStationDownloadMyPokemonResult {
	return &GlobalTradeStationDownloadMyPokemonResult{}
}
