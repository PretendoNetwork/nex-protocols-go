// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationTradeKey is a type within the DataStore protocol
type GlobalTradeStationTradeKey struct {
	types.Structure
	DataID  types.UInt64
	Version types.UInt32
}

// WriteTo writes the GlobalTradeStationTradeKey to the given writable
func (gtstk GlobalTradeStationTradeKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtstk.DataID.WriteTo(contentWritable)
	gtstk.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtstk.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationTradeKey from the given readable
func (gtstk *GlobalTradeStationTradeKey) ExtractFrom(readable types.Readable) error {
	var err error

	err = gtstk.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey header. %s", err.Error())
	}

	err = gtstk.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.DataID. %s", err.Error())
	}

	err = gtstk.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationTradeKey.Version. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationTradeKey
func (gtstk GlobalTradeStationTradeKey) Copy() types.RVType {
	copied := NewGlobalTradeStationTradeKey()

	copied.StructureVersion = gtstk.StructureVersion
	copied.DataID = gtstk.DataID.Copy().(types.UInt64)
	copied.Version = gtstk.Version.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given GlobalTradeStationTradeKey contains the same data as the current GlobalTradeStationTradeKey
func (gtstk GlobalTradeStationTradeKey) Equals(o types.RVType) bool {
	if _, ok := o.(GlobalTradeStationTradeKey); !ok {
		return false
	}

	other := o.(GlobalTradeStationTradeKey)

	if gtstk.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtstk.DataID.Equals(other.DataID) {
		return false
	}

	return gtstk.Version.Equals(other.Version)
}

// CopyRef copies the current value of the GlobalTradeStationTradeKey
// and returns a pointer to the new copy
func (gtstk GlobalTradeStationTradeKey) CopyRef() types.RVTypePtr {
	copied := gtstk.Copy().(GlobalTradeStationTradeKey)
	return &copied
}

// Deref takes a pointer to the GlobalTradeStationTradeKey
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gtstk *GlobalTradeStationTradeKey) Deref() types.RVType {
	return *gtstk
}

// String returns the string representation of the GlobalTradeStationTradeKey
func (gtstk GlobalTradeStationTradeKey) String() string {
	return gtstk.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationTradeKey using the provided indentation level
func (gtstk GlobalTradeStationTradeKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationTradeKey{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, gtstk.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, gtstk.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationTradeKey returns a new GlobalTradeStationTradeKey
func NewGlobalTradeStationTradeKey() GlobalTradeStationTradeKey {
	return GlobalTradeStationTradeKey{
		DataID:  types.NewUInt64(0),
		Version: types.NewUInt32(0),
	}

}
