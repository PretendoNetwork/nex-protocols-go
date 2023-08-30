// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationPrepareTradePokemonParam holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationPrepareTradePokemonParam struct {
	nex.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareUploadKey *GlobalTradeStationRecordKey
}

// ExtractFromStream extracts a GlobalTradeStationPrepareTradePokemonParam structure from a stream
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	tradeKey, err := stream.ReadStructure(NewGlobalTradeStationTradeKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.TradeKey from stream. %s", err.Error())
	}

	globalTradeStationPrepareTradePokemonParam.TradeKey = tradeKey.(*GlobalTradeStationTradeKey)

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationPrepareTradePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationPrepareTradePokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	return nil
}

// Bytes encodes the GlobalTradeStationPrepareTradePokemonParam and returns a byte array
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationPrepareTradePokemonParam.TradeKey)
	stream.WriteStructure(globalTradeStationPrepareTradePokemonParam.PrepareUploadKey)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationPrepareTradePokemonParam
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationPrepareTradePokemonParam()

	copied.SetStructureVersion(globalTradeStationPrepareTradePokemonParam.StructureVersion())

	copied.TradeKey = globalTradeStationPrepareTradePokemonParam.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareUploadKey = globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationPrepareTradePokemonParam)

	if globalTradeStationPrepareTradePokemonParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !globalTradeStationPrepareTradePokemonParam.TradeKey.Equals(other.TradeKey) {
		return false
	}

	if !globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) String() string {
	return globalTradeStationPrepareTradePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationPrepareTradePokemonParam *GlobalTradeStationPrepareTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationPrepareTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationPrepareTradePokemonParam.StructureVersion()))

	if globalTradeStationPrepareTradePokemonParam.TradeKey != nil {
		b.WriteString(fmt.Sprintf("%sTradeKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonParam.TradeKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTradeKey: nil\n", indentationValues))
	}

	if globalTradeStationPrepareTradePokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationPrepareTradePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationPrepareTradePokemonParam returns a new GlobalTradeStationPrepareTradePokemonParam
func NewGlobalTradeStationPrepareTradePokemonParam() *GlobalTradeStationPrepareTradePokemonParam {
	return &GlobalTradeStationPrepareTradePokemonParam{}
}
