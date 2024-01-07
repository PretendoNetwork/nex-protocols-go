// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationRecordKey holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationRecordKey struct {
	types.Structure
	DataID   *types.PrimitiveU64
	Password *types.PrimitiveU64
}

// ExtractFrom extracts the GlobalTradeStationRecordKey from the given readable
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationRecordKey.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationRecordKey header. %s", err.Error())
	}

	err = globalTradeStationRecordKey.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.DataID from stream. %s", err.Error())
	}

	err = globalTradeStationRecordKey.Password.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationRecordKey.Password from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationRecordKey to the given writable
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationRecordKey.DataID.WriteTo(contentWritable)
	globalTradeStationRecordKey.Password.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationRecordKey.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationRecordKey
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) Copy() types.RVType {
	copied := NewGlobalTradeStationRecordKey()

	copied.StructureVersion = globalTradeStationRecordKey.StructureVersion

	copied.DataID = globalTradeStationRecordKey.DataID.Copy().(*types.PrimitiveU64)
	copied.Password = globalTradeStationRecordKey.Password.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationRecordKey *GlobalTradeStationRecordKey) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationRecordKey); !ok {
		return false
	}

	other := o.(*GlobalTradeStationRecordKey)

	if globalTradeStationRecordKey.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationRecordKey.DataID.Equals(other.DataID) {
		return false
	}

	if !globalTradeStationRecordKey.Password.Equals(other.Password) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationRecordKey.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, globalTradeStationRecordKey.DataID))
	b.WriteString(fmt.Sprintf("%sPassword: %s,\n", indentationValues, globalTradeStationRecordKey.Password))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationRecordKey returns a new GlobalTradeStationRecordKey
func NewGlobalTradeStationRecordKey() *GlobalTradeStationRecordKey {
	return &GlobalTradeStationRecordKey{
		DataID: types.NewPrimitiveU64(0),
		Password: types.NewPrimitiveU64(0),
	}
}
