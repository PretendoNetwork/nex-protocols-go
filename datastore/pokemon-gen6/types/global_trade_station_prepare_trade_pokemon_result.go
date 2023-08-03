// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationPrepareTradePokemonResult holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationPrepareTradePokemonResult struct {
	nex.Structure
	Result          *GlobalTradeStationDownloadPokemonResult
	PrepareTradeKey *GlobalTradeStationRecordKey
}

// ExtractFromStream extracts a GlobalTradeStationPrepareTradePokemonResult structure from a stream
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	result, err := stream.ReadStructure(NewGlobalTradeStationDownloadPokemonResult())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.Result from stream. %s", err.Error())
	}

	globalTradeStationPrepareTradePokemonResult.Result = result.(*GlobalTradeStationDownloadPokemonResult)

	prepareTradeKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonResult.PrepareTradeKey from stream. %s", err.Error())
	}

	globalTradeStationPrepareTradePokemonResult.PrepareTradeKey = prepareTradeKey.(*GlobalTradeStationRecordKey)

	return nil
}

// Bytes encodes the GlobalTradeStationPrepareTradePokemonResult and returns a byte array
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationPrepareTradePokemonResult.Result)
	stream.WriteStructure(globalTradeStationPrepareTradePokemonResult.PrepareTradeKey)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonResult
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationPrepareTradePokemonResult()

	copied.Result = globalTradeStationPrepareTradePokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.PrepareTradeKey = globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationPrepareTradePokemonResult *GlobalTradeStationPrepareTradePokemonResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationPrepareTradePokemonResult)

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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationPrepareTradePokemonResult.StructureVersion()))

	if globalTradeStationPrepareTradePokemonResult.Result != nil {
		b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationPrepareTradePokemonResult.Result.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: nil\n", indentationValues))
	}

	if globalTradeStationPrepareTradePokemonResult.PrepareTradeKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonResult.PrepareTradeKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareTradeKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonResult returns a new GlobalTradeStationPrepareTradePokemonResult
func NewGlobalTradeStationPrepareTradePokemonResult() *GlobalTradeStationPrepareTradePokemonResult {
	return &GlobalTradeStationPrepareTradePokemonResult{}
}
