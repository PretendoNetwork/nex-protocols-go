// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationTradePokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationTradePokemonParam struct {
	nex.Structure
	TradeKey         *GlobalTradeStationTradeKey
	PrepareTradeKey  *GlobalTradeStationRecordKey
	PrepareUploadKey *GlobalTradeStationRecordKey
	Period           uint16
	IndexData        []byte
	PokemonData      []byte
	Signature        []byte
	NeedData         bool
}

// ExtractFromStream extracts a GlobalTradeStationTradePokemonParam structure from a stream
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationTradePokemonParam.TradeKey, err = nex.StreamReadStructure(stream, NewGlobalTradeStationTradeKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.TradeKey from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.PrepareTradeKey, err = nex.StreamReadStructure(stream, NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareTradeKey from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.PrepareUploadKey, err = nex.StreamReadStructure(stream, NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Period from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.IndexData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.IndexData from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.PokemonData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.PokemonData from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.Signature, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.Signature from stream. %s", err.Error())
	}

	globalTradeStationTradePokemonParam.NeedData, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradePokemonParam.NeedData from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationTradePokemonParam and returns a byte array
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationTradePokemonParam.TradeKey)
	stream.WriteStructure(globalTradeStationTradePokemonParam.PrepareTradeKey)
	stream.WriteStructure(globalTradeStationTradePokemonParam.PrepareUploadKey)
	stream.WriteUInt16LE(globalTradeStationTradePokemonParam.Period)
	stream.WriteQBuffer(globalTradeStationTradePokemonParam.IndexData)
	stream.WriteQBuffer(globalTradeStationTradePokemonParam.PokemonData)
	stream.WriteQBuffer(globalTradeStationTradePokemonParam.Signature)
	stream.WriteBool(globalTradeStationTradePokemonParam.NeedData)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationTradePokemonParam
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationTradePokemonParam()

	copied.SetStructureVersion(globalTradeStationTradePokemonParam.StructureVersion())

	copied.TradeKey = globalTradeStationTradePokemonParam.TradeKey.Copy().(*GlobalTradeStationTradeKey)
	copied.PrepareTradeKey = globalTradeStationTradePokemonParam.PrepareTradeKey.Copy().(*GlobalTradeStationRecordKey)
	copied.PrepareUploadKey = globalTradeStationTradePokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Period = globalTradeStationTradePokemonParam.Period
	copied.IndexData = globalTradeStationTradePokemonParam.IndexData
	copied.PokemonData = globalTradeStationTradePokemonParam.PokemonData
	copied.Signature = globalTradeStationTradePokemonParam.Signature
	copied.NeedData = globalTradeStationTradePokemonParam.NeedData

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationTradePokemonParam)

	if globalTradeStationTradePokemonParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !globalTradeStationTradePokemonParam.TradeKey.Equals(other.TradeKey) {
		return false
	}

	if !globalTradeStationTradePokemonParam.PrepareTradeKey.Equals(other.PrepareTradeKey) {
		return false
	}

	if !globalTradeStationTradePokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if globalTradeStationTradePokemonParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(globalTradeStationTradePokemonParam.IndexData, other.IndexData) {
		return false
	}

	if !bytes.Equal(globalTradeStationTradePokemonParam.PokemonData, other.PokemonData) {
		return false
	}

	if !bytes.Equal(globalTradeStationTradePokemonParam.Signature, other.Signature) {
		return false
	}

	if globalTradeStationTradePokemonParam.NeedData != other.NeedData {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) String() string {
	return globalTradeStationTradePokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationTradePokemonParam *GlobalTradeStationTradePokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradePokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationTradePokemonParam.StructureVersion()))

	if globalTradeStationTradePokemonParam.TradeKey != nil {
		b.WriteString(fmt.Sprintf("%sTradeKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.TradeKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTradeKey: nil\n", indentationValues))
	}

	if globalTradeStationTradePokemonParam.PrepareTradeKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareTradeKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.PrepareTradeKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareTradeKey: nil\n", indentationValues))
	}

	if globalTradeStationTradePokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationTradePokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, globalTradeStationTradePokemonParam.Period))
	b.WriteString(fmt.Sprintf("%sIndexData: %x,\n", indentationValues, globalTradeStationTradePokemonParam.IndexData))
	b.WriteString(fmt.Sprintf("%sPokemonData: %x,\n", indentationValues, globalTradeStationTradePokemonParam.PokemonData))
	b.WriteString(fmt.Sprintf("%sSignature: %x,\n", indentationValues, globalTradeStationTradePokemonParam.Signature))
	b.WriteString(fmt.Sprintf("%sNeedData: %t,\n", indentationValues, globalTradeStationTradePokemonParam.NeedData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradePokemonParam returns a new GlobalTradeStationTradePokemonParam
func NewGlobalTradeStationTradePokemonParam() *GlobalTradeStationTradePokemonParam {
	return &GlobalTradeStationTradePokemonParam{}
}
