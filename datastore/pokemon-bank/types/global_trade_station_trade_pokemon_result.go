// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationTradePokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationTradePokemonResult struct {
	nex.Structure
	Result   *GlobalTradeStationDownloadPokemonResult
	MyDataID uint64
}

// ExtractFromStream extracts a GlobalTradeStationTradePokemonResult structure from a stream
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationTradePokemonResult.Result, err = nex.StreamReadStructure(stream, NewGlobalTradeStationDownloadPokemonResult())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.Result from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonResult.MyDataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonResult.MyDataID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationTradePokemonResult and returns a byte array
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationTradePokemonResult.Result)
	stream.WriteUInt64LE(globalTradeStationTradePokemonResult.MyDataID)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonResult
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationTradePokemonResult()

	copied.SetStructureVersion(globalTradeStationTradePokemonResult.StructureVersion())

	copied.Result = globalTradeStationTradePokemonResult.Result.Copy().(*GlobalTradeStationDownloadPokemonResult)
	copied.MyDataID = globalTradeStationTradePokemonResult.MyDataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradePokemonResult *GlobalTradeStationTradePokemonResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationTradePokemonResult)

	if globalTradeStationTradePokemonResult.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !globalTradeStationTradePokemonResult.Result.Equals(other.Result) {
		return false
	}

	if globalTradeStationTradePokemonResult.MyDataID != other.MyDataID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationTradePokemonResult.StructureVersion()))

	if globalTradeStationTradePokemonResult.Result != nil {
		b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, globalTradeStationTradePokemonResult.Result.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sMyDataID: %d,\n", indentationValues, globalTradeStationTradePokemonResult.MyDataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonResult returns a new GlobalTradeStationTradePokemonResult
func NewGlobalTradeStationTradePokemonResult() *GlobalTradeStationTradePokemonResult {
	return &GlobalTradeStationTradePokemonResult{}
}
