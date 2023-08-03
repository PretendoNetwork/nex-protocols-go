// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationDownloadOtherPokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDownloadOtherPokemonParam struct {
	nex.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFromStream extracts a GlobalTradeStationDownloadOtherPokemonParam structure from a stream
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadOtherPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	return nil
}

// Bytes encodes the GlobalTradeStationDownloadOtherPokemonParam and returns a byte array
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationDownloadOtherPokemonParam
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationDownloadOtherPokemonParam()

	copied.PrepareUploadKey = globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationDownloadOtherPokemonParam)

	return globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns a string representation of the struct
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) String() string {
	return globalTradeStationDownloadOtherPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadOtherPokemonParam *GlobalTradeStationDownloadOtherPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadOtherPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationDownloadOtherPokemonParam.StructureVersion()))

	if globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDownloadOtherPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadOtherPokemonParam returns a new GlobalTradeStationDownloadOtherPokemonParam
func NewGlobalTradeStationDownloadOtherPokemonParam() *GlobalTradeStationDownloadOtherPokemonParam {
	return &GlobalTradeStationDownloadOtherPokemonParam{}
}
