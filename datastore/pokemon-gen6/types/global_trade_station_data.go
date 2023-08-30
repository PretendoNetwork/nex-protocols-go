// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationData holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationData struct {
	nex.Structure
	DataID      uint64
	OwnerID     uint32
	UpdatedTime *nex.DateTime
	IndexData   []byte
	Version     uint32
}

// ExtractFromStream extracts a GlobalTradeStationData structure from a stream
func (globalTradeStationData *GlobalTradeStationData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationData.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.DataID from stream. %s", err.Error())
	}

	globalTradeStationData.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.OwnerID from stream. %s", err.Error())
	}

	globalTradeStationData.UpdatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.UpdatedTime from stream. %s", err.Error())
	}

	globalTradeStationData.IndexData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.IndexData from stream. %s", err.Error())
	}

	globalTradeStationData.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.Version from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationData and returns a byte array
func (globalTradeStationData *GlobalTradeStationData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(globalTradeStationData.DataID)
	stream.WriteUInt32LE(globalTradeStationData.OwnerID)
	stream.WriteDateTime(globalTradeStationData.UpdatedTime)
	stream.WriteQBuffer(globalTradeStationData.IndexData)
	stream.WriteUInt32LE(globalTradeStationData.Version)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationData
func (globalTradeStationData *GlobalTradeStationData) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationData()

	copied.SetStructureVersion(globalTradeStationData.StructureVersion())

	copied.DataID = globalTradeStationData.DataID
	copied.OwnerID = globalTradeStationData.OwnerID
	copied.UpdatedTime = globalTradeStationData.UpdatedTime.Copy()
	copied.IndexData = globalTradeStationData.IndexData
	copied.Version = globalTradeStationData.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationData *GlobalTradeStationData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationData)

	if globalTradeStationData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if globalTradeStationData.DataID != other.DataID {
		return false
	}

	if globalTradeStationData.OwnerID != other.OwnerID {
		return false
	}

	if !globalTradeStationData.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !bytes.Equal(globalTradeStationData.IndexData, other.IndexData) {
		return false
	}

	if globalTradeStationData.Version != other.Version {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationData *GlobalTradeStationData) String() string {
	return globalTradeStationData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationData *GlobalTradeStationData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, globalTradeStationData.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, globalTradeStationData.OwnerID))

	if globalTradeStationData.UpdatedTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, globalTradeStationData.UpdatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sIndexData: %x,\n", indentationValues, globalTradeStationData.IndexData))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, globalTradeStationData.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationData returns a new GlobalTradeStationData
func NewGlobalTradeStationData() *GlobalTradeStationData {
	return &GlobalTradeStationData{}
}
