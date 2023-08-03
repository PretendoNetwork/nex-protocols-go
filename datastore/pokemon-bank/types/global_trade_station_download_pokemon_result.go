// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationDownloadPokemonResult holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDownloadPokemonResult struct {
	nex.Structure
	DataID      uint64
	IndexData   []byte
	PokemonData []byte
}

// ExtractFromStream extracts a GlobalTradeStationDownloadPokemonResult structure from a stream
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationDownloadPokemonResult.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.DataID from stream. %s", err.Error())
	}

	globalTradeStationDownloadPokemonResult.IndexData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.IndexData from stream. %s", err.Error())
	}

	globalTradeStationDownloadPokemonResult.PokemonData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadPokemonResult.PokemonData from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationDownloadPokemonResult and returns a byte array
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(globalTradeStationDownloadPokemonResult.DataID)
	stream.WriteQBuffer(globalTradeStationDownloadPokemonResult.IndexData)
	stream.WriteQBuffer(globalTradeStationDownloadPokemonResult.PokemonData)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationDownloadPokemonResult
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationDownloadPokemonResult()

	copied.DataID = globalTradeStationDownloadPokemonResult.DataID
	copied.IndexData = globalTradeStationDownloadPokemonResult.IndexData
	copied.PokemonData = globalTradeStationDownloadPokemonResult.PokemonData

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadPokemonResult *GlobalTradeStationDownloadPokemonResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationDownloadPokemonResult)

	if globalTradeStationDownloadPokemonResult.DataID != other.DataID {
		return false
	}

	if !bytes.Equal(globalTradeStationDownloadPokemonResult.IndexData, other.IndexData) {
		return false
	}

	if !bytes.Equal(globalTradeStationDownloadPokemonResult.PokemonData, other.PokemonData) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationDownloadPokemonResult.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, globalTradeStationDownloadPokemonResult.DataID))
	b.WriteString(fmt.Sprintf("%sIndexData: %x,\n", indentationValues, globalTradeStationDownloadPokemonResult.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %x,\n", indentationValues, globalTradeStationDownloadPokemonResult.PokemonData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadPokemonResult returns a new GlobalTradeStationDownloadPokemonResult
func NewGlobalTradeStationDownloadPokemonResult() *GlobalTradeStationDownloadPokemonResult {
	return &GlobalTradeStationDownloadPokemonResult{}
}
