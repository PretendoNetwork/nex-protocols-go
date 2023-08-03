// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationTradeKey holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationTradeKey struct {
	nex.Structure
	DataID  uint64
	Version uint32
}

// ExtractFromStream extracts a GlobalTradeStationTradeKey structure from a stream
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationTradeKey.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.DataID from stream. %s", err.Error())
	}

	globalTradeStationTradeKey.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.Version from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationTradeKey and returns a byte array
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(globalTradeStationTradeKey.DataID)
	stream.WriteUInt32LE(globalTradeStationTradeKey.Version)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationTradeKey
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationTradeKey()

	copied.DataID = globalTradeStationTradeKey.DataID
	copied.Version = globalTradeStationTradeKey.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationTradeKey)

	if globalTradeStationTradeKey.DataID != other.DataID {
		return false
	}

	if globalTradeStationTradeKey.Version != other.Version {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) String() string {
	return globalTradeStationTradeKey.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradeKey{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationTradeKey.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, globalTradeStationTradeKey.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, globalTradeStationTradeKey.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradeKey returns a new GlobalTradeStationTradeKey
func NewGlobalTradeStationTradeKey() *GlobalTradeStationTradeKey {
	return &GlobalTradeStationTradeKey{}
}
