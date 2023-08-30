// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationUploadPokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationUploadPokemonParam struct {
	nex.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           uint16
	IndexData        []byte
	PokemonData      []byte
	Signature        []byte
}

// ExtractFromStream extracts a GlobalTradeStationUploadPokemonParam structure from a stream
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationUploadPokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	globalTradeStationUploadPokemonParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Period from stream. %s", err.Error())
	}

	globalTradeStationUploadPokemonParam.IndexData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.IndexData from stream. %s", err.Error())
	}

	globalTradeStationUploadPokemonParam.PokemonData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.PokemonData from stream. %s", err.Error())
	}

	globalTradeStationUploadPokemonParam.Signature, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationUploadPokemonParam.Signature from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationUploadPokemonParam and returns a byte array
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationUploadPokemonParam.PrepareUploadKey)
	stream.WriteUInt16LE(globalTradeStationUploadPokemonParam.Period)
	stream.WriteQBuffer(globalTradeStationUploadPokemonParam.IndexData)
	stream.WriteQBuffer(globalTradeStationUploadPokemonParam.PokemonData)
	stream.WriteQBuffer(globalTradeStationUploadPokemonParam.Signature)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationUploadPokemonParam
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationUploadPokemonParam()

	copied.SetStructureVersion(globalTradeStationUploadPokemonParam.StructureVersion())

	copied.PrepareUploadKey = globalTradeStationUploadPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = globalTradeStationUploadPokemonParam.Period
	copied.IndexData = globalTradeStationUploadPokemonParam.IndexData
	copied.PokemonData = globalTradeStationUploadPokemonParam.PokemonData
	copied.Signature = globalTradeStationUploadPokemonParam.Signature

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationUploadPokemonParam)

	if globalTradeStationUploadPokemonParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !globalTradeStationUploadPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if globalTradeStationUploadPokemonParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(globalTradeStationUploadPokemonParam.IndexData, other.IndexData) {
		return false
	}

	if !bytes.Equal(globalTradeStationUploadPokemonParam.PokemonData, other.PokemonData) {
		return false
	}

	if !bytes.Equal(globalTradeStationUploadPokemonParam.Signature, other.Signature) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) String() string {
	return globalTradeStationUploadPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationUploadPokemonParam *GlobalTradeStationUploadPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationUploadPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationUploadPokemonParam.StructureVersion()))

	if globalTradeStationUploadPokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationUploadPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, globalTradeStationUploadPokemonParam.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %x,\n", indentationValues, globalTradeStationUploadPokemonParam.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %x,\n", indentationValues, globalTradeStationUploadPokemonParam.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %x,\n", indentationValues, globalTradeStationUploadPokemonParam.Signature))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationUploadPokemonParam returns a new GlobalTradeStationUploadPokemonParam
func NewGlobalTradeStationUploadPokemonParam() *GlobalTradeStationUploadPokemonParam {
	return &GlobalTradeStationUploadPokemonParam{}
}
