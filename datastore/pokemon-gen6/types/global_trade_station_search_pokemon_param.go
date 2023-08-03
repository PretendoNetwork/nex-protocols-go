// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationSearchPokemonParam holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationSearchPokemonParam struct {
	nex.Structure
	PrepareUploadKey  *GlobalTradeStationRecordKey
	Conditions        []uint32
	ResultOrderColumn uint8
	ResultOrder       uint8
	UploadedAfter     *nex.DateTime
	UploadedBefore    *nex.DateTime
	ResultRange       *nex.ResultRange
}

// ExtractFromStream extracts a GlobalTradeStationSearchPokemonParam structure from a stream
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	prepareUploadKey, err := stream.ReadStructure(NewGlobalTradeStationRecordKey())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.PrepareUploadKey from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.PrepareUploadKey = prepareUploadKey.(*GlobalTradeStationRecordKey)

	globalTradeStationSearchPokemonParam.Conditions, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.Conditions from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.ResultOrderColumn, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrderColumn from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.ResultOrder, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultOrder from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.UploadedAfter, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedAfter from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.UploadedBefore, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.UploadedBefore from stream. %s", err.Error())
	}

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationSearchPokemonParam.ResultRange from stream. %s", err.Error())
	}

	globalTradeStationSearchPokemonParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the GlobalTradeStationSearchPokemonParam and returns a byte array
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(globalTradeStationSearchPokemonParam.PrepareUploadKey)
	stream.WriteListUInt32LE(globalTradeStationSearchPokemonParam.Conditions)
	stream.WriteUInt8(globalTradeStationSearchPokemonParam.ResultOrderColumn)
	stream.WriteUInt8(globalTradeStationSearchPokemonParam.ResultOrder)
	stream.WriteDateTime(globalTradeStationSearchPokemonParam.UploadedAfter)
	stream.WriteDateTime(globalTradeStationSearchPokemonParam.UploadedBefore)
	stream.WriteStructure(globalTradeStationSearchPokemonParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationSearchPokemonParam
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationSearchPokemonParam()

	copied.PrepareUploadKey = globalTradeStationSearchPokemonParam.PrepareUploadKey.Copy().(*GlobalTradeStationRecordKey)
	copied.Conditions = make([]uint32, len(globalTradeStationSearchPokemonParam.Conditions))

	copy(copied.Conditions, globalTradeStationSearchPokemonParam.Conditions)

	copied.ResultOrderColumn = globalTradeStationSearchPokemonParam.ResultOrderColumn
	copied.ResultOrder = globalTradeStationSearchPokemonParam.ResultOrder
	copied.UploadedAfter = globalTradeStationSearchPokemonParam.UploadedAfter.Copy()
	copied.UploadedBefore = globalTradeStationSearchPokemonParam.UploadedBefore.Copy()
	copied.ResultRange = globalTradeStationSearchPokemonParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationSearchPokemonParam)

	if !globalTradeStationSearchPokemonParam.PrepareUploadKey.Equals(other.PrepareUploadKey) {
		return false
	}

	if len(globalTradeStationSearchPokemonParam.Conditions) != len(other.Conditions) {
		return false
	}

	for i := 0; i < len(globalTradeStationSearchPokemonParam.Conditions); i++ {
		if globalTradeStationSearchPokemonParam.Conditions[i] != other.Conditions[i] {
			return false
		}
	}

	if globalTradeStationSearchPokemonParam.ResultOrderColumn != other.ResultOrderColumn {
		return false
	}

	if globalTradeStationSearchPokemonParam.ResultOrder != other.ResultOrder {
		return false
	}

	if !globalTradeStationSearchPokemonParam.UploadedAfter.Equals(other.UploadedAfter) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.UploadedBefore.Equals(other.UploadedBefore) {
		return false
	}

	if !globalTradeStationSearchPokemonParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) String() string {
	return globalTradeStationSearchPokemonParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationSearchPokemonParam *GlobalTradeStationSearchPokemonParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationSearchPokemonParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationSearchPokemonParam.StructureVersion()))

	if globalTradeStationSearchPokemonParam.PrepareUploadKey != nil {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: %s\n", indentationValues, globalTradeStationSearchPokemonParam.PrepareUploadKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrepareUploadKey: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sConditions: %v,\n", indentationValues, globalTradeStationSearchPokemonParam.Conditions))
	b.WriteString(fmt.Sprintf("%sResultOrderColumn: %d,\n", indentationValues, globalTradeStationSearchPokemonParam.ResultOrderColumn))
	b.WriteString(fmt.Sprintf("%sResultOrder: %d,\n", indentationValues, globalTradeStationSearchPokemonParam.ResultOrder))

	if globalTradeStationSearchPokemonParam.UploadedAfter != nil {
		b.WriteString(fmt.Sprintf("%sUploadedAfter: %s\n", indentationValues, globalTradeStationSearchPokemonParam.UploadedAfter.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUploadedAfter: nil\n", indentationValues))
	}

	if globalTradeStationSearchPokemonParam.UploadedBefore != nil {
		b.WriteString(fmt.Sprintf("%sUploadedBefore: %s\n", indentationValues, globalTradeStationSearchPokemonParam.UploadedBefore.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUploadedBefore: nil\n", indentationValues))
	}

	if globalTradeStationSearchPokemonParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, globalTradeStationSearchPokemonParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationSearchPokemonParam returns a new GlobalTradeStationSearchPokemonParam
func NewGlobalTradeStationSearchPokemonParam() *GlobalTradeStationSearchPokemonParam {
	return &GlobalTradeStationSearchPokemonParam{}
}
