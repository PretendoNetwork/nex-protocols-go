// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationDownloadMyPokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationDownloadMyPokemonParam struct {
	nex.Structure
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFromStream extracts a GlobalTradeStationDownloadMyPokemonParam structure from a stream
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationDownloadMyPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationDownloadMyPokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	return nil
}

// Bytes encodes the GlobalTradeStationDownloadMyPokemonParam and returns a byte array
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationDownloadMyPokemonParam.PrepareUploadKey)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationDownloadMyPokemonParam
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationDownloadMyPokemonParam()

	copied.PrepareUploadKey = globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationDownloadMyPokemonParam)

	return globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey)
}

// String returns a string representation of the struct
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) String() string {
	return globalTradeStationDownloadMyPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationDownloadMyPokemonParam *GlobalTradeStationDownloadMyPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationDownloadMyPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationDownloadMyPokemonParam.StructureVersion()))

	if globalTradeStationDownloadMyPokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationDownloadMyPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationDownloadMyPokemonParam returns a new GlobalTradeStationDownloadMyPokemonParam
func NewGlobalTradeStationDownloadMyPokemonParam() *GlobalTradeStationDownloadMyPokemonParam {
	return &GlobalTradeStationDownloadMyPokemonParam{}
}
