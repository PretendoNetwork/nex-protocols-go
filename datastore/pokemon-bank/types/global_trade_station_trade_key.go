// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationTradeKey holds data for the DataStore (Pokemon Bank) protocol
type GlobalTradeStationTradeKey struct {
	types.Structure
	DataID  *types.PrimitiveU64
	Version *types.PrimitiveU32
}

// ExtractFrom extracts the GlobalTradeStationTradeKey from the given readable
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationTradeKey.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationTradeKey header. %s", err.Error())
	}

	err = globalTradeStationTradeKey.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.DataID from stream. %s", err.Error())
	}

	err = globalTradeStationTradeKey.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.Version from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationTradeKey to the given writable
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationTradeKey.DataID.WriteTo(contentWritable)
	globalTradeStationTradeKey.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationTradeKey.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationTradeKey
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) Copy() types.RVType {
	copied := NewGlobalTradeStationTradeKey()

	copied.StructureVersion = globalTradeStationTradeKey.StructureVersion

	copied.DataID = globalTradeStationTradeKey.DataID.Copy().(*types.PrimitiveU64)
	copied.Version = globalTradeStationTradeKey.Version.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationTradeKey *GlobalTradeStationTradeKey) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationTradeKey); !ok {
		return false
	}

	other := o.(*GlobalTradeStationTradeKey)

	if globalTradeStationTradeKey.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationTradeKey.DataID.Equals(other.DataID) {
		return false
	}

	if !globalTradeStationTradeKey.Version.Equals(other.Version) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationTradeKey.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, globalTradeStationTradeKey.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, globalTradeStationTradeKey.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradeKey returns a new GlobalTradeStationTradeKey
func NewGlobalTradeStationTradeKey() *GlobalTradeStationTradeKey {
	return &GlobalTradeStationTradeKey{
		DataID: types.NewPrimitiveU64(0),
		Version: types.NewPrimitiveU32(0),
	}
}
