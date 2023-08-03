// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationSearchPokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationSearchPokemonResult struct {
	nex.Structure
	TotalCount     uint32
	Result         []*GlobalTradeStationData
	TotalCountType uint8
}

// ExtractFromStream extracts a GlobalTradeStationSearchPokemonResult structure from a stream
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationSearchPokemonResult.TotalCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCount from stream. %s", err.Error())
	}

	result, err := stream.ReadListStructure(NewGlobalTradeStationData())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.Result from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonResult.Result = result.([]*GlobalTradeStationData)

	globalTradeStationSearchPokemonResult.TotalCountType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonResult.TotalCountType from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationSearchPokemonResult and returns a byte array
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(globalTradeStationSearchPokemonResult.TotalCount)
	stream.WriteListStructure(globalTradeStationSearchPokemonResult.Result)
	stream.WriteUInt8(globalTradeStationSearchPokemonResult.TotalCountType)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonResult
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationSearchPokemonResult()

	copied.TotalCount = globalTradeStationSearchPokemonResult.TotalCount
	copied.Result = make([]*GlobalTradeStationData, len(globalTradeStationSearchPokemonResult.Result))

	for i := 0; i < len(globalTradeStationSearchPokemonResult.Result); i++ {
		copied.Result[i] = globalTradeStationSearchPokemonResult.Result[i].Copy().(*GlobalTradeStationData)
	}

	copied.TotalCountType = globalTradeStationSearchPokemonResult.TotalCountType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationSearchPokemonResult)

	if globalTradeStationSearchPokemonResult.TotalCount != other.TotalCount {
		return false
	}

	if len(globalTradeStationSearchPokemonResult.Result) != len(other.Result) {
		return false
	}

	for i := 0; i < len(globalTradeStationSearchPokemonResult.Result); i++ {
		if !globalTradeStationSearchPokemonResult.Result[i].Equals(other.Result[i]) {
			return false
		}
	}

	return globalTradeStationSearchPokemonResult.TotalCountType == other.TotalCountType
}

// String returns a string representation of the struct
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) String() string {
	return globalTradeStationSearchPokemonResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationSearchPokemonResult *GlobalTradeStationSearchPokemonResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonResult{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationSearchPokemonResult.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTotalCount: %d,\n", indentationValues, globalTradeStationSearchPokemonResult.TotalCount))

	if len(globalTradeStationSearchPokemonResult.Result) == 0 {
		b.WriteString(fmt.Sprintf("%sResult: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: [\n", indentationValues))

		for i := 0; i < len(globalTradeStationSearchPokemonResult.Result); i++ {
			str := globalTradeStationSearchPokemonResult.Result[i].FormatToString(indentationLevel + 2)
			if i == len(globalTradeStationSearchPokemonResult.Result)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTotalCountType: %d,\n", indentationValues, globalTradeStationSearchPokemonResult.TotalCountType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonResult returns a new GlobalTradeStationSearchPokemonResult
func NewGlobalTradeStationSearchPokemonResult() *GlobalTradeStationSearchPokemonResult {
	return &GlobalTradeStationSearchPokemonResult{}
}
