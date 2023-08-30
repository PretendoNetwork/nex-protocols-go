// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GlobalTradeStationRecordKey holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationRecordKey struct {
	nex.Structure
	DataID   uint64
	Password uint64
}

// ExtractFromStream extracts a GlobalTradeStationRecordKey structure from a stream
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	globalTradeStationRecordKey.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.DataID from stream. %s", err.Error())
	}

	globalTradeStationRecordKey.Password, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.Password from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the GlobalTradeStationRecordKey and returns a byte array
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(globalTradeStationRecordKey.DataID)
	stream.WriteUInt64LE(globalTradeStationRecordKey.Password)

	return stream.Bytes()
}

// Copy returns a new copied instance of GlobalTradeStationRecordKey
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) Copy() nex.StructureInterface {
	copied := NewGlobalTradeStationRecordKey()

	copied.SetStructureVersion(globalTradeStationRecordKey.StructureVersion())

	copied.DataID = globalTradeStationRecordKey.DataID
	copied.Password = globalTradeStationRecordKey.Password

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GlobalTradeStationRecordKey)

	if globalTradeStationRecordKey.StructureVersion() != other.StructureVersion() {
		return false
	}

	if globalTradeStationRecordKey.DataID != other.DataID {
		return false
	}

	if globalTradeStationRecordKey.Password != other.Password {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) String() string {
	return globalTradeStationRecordKey.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationRecordKey{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, globalTradeStationRecordKey.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, globalTradeStationRecordKey.DataID))
	b.WriteString(fmt.Sprintf("%sPassword: %d,\n", indentationValues, globalTradeStationRecordKey.Password))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationRecordKey returns a new GlobalTradeStationRecordKey
func NewGlobalTradeStationRecordKey() *GlobalTradeStationRecordKey {
	return &GlobalTradeStationRecordKey{}
}
