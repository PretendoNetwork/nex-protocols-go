// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationDeletePokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDeletePokemonParam struct {
	nex.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
	DeleteFlag       uint8
}

// ExtractFromStream extracts a GlobalTradeStationDeletePokemonParam structure from a stream
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationDeletePokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	globalTradeStationDeletePokemonParam.DeleteFlag, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDeletePokemonParam.DeleteFlag from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationDeletePokemonParam and returns a byte array
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationDeletePokemonParam.PrepareUploadKey)
	stream.WriteUInt8(globalTradeStationDeletePokemonParam.DeleteFlag)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationDeletePokemonParam
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationDeletePokemonParam()

	copied.SetStructureVersion(globalTradeStationDeletePokemonParam.StructureVersion())

	copied.PrepareUploadKey = globalTradeStationDeletePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.DeleteFlag = globalTradeStationDeletePokemonParam.DeleteFlag

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationDeletePokemonParam)

	if globalTradeStationDeletePokemonParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !globalTradeStationDeletePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if globalTradeStationDeletePokemonParam.DeleteFlag != other.DeleteFlag {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) String() string {
	return globalTradeStationDeletePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDeletePokemonParam *GlobalTradeStationDeletePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDeletePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationDeletePokemonParam.StructureVersion()))

	if globalTradeStationDeletePokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDeletePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sDeleteFlag: %d,\n", indentationValues, globalTradeStationDeletePokemonParam.DeleteFlag))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDeletePokemonParam returns a new GlobalTradeStationDeletePokemonParam
func NewGlobalTradeStationDeletePokemonParam() *GlobalTradeStationDeletePokemonParam {
	return &GlobalTradeStationDeletePokemonParam{}
}
