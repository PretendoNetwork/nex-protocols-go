// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GlobalTradeStationData is a type within the DataStore protocol
type GlobalTradeStationData struct {
	types.Structure
	DataID      types.UInt64
	OwnerID     types.PID
	UpdatedTime types.DateTime
	IndexData   types.QBuffer
	Version     types.UInt32
}

// WriteTo writes the GlobalTradeStationData to the given writable
func (gtsd GlobalTradeStationData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gtsd.DataID.WriteTo(contentWritable)
	gtsd.OwnerID.WriteTo(contentWritable)
	gtsd.UpdatedTime.WriteTo(contentWritable)
	gtsd.IndexData.WriteTo(contentWritable)
	gtsd.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gtsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GlobalTradeStationData from the given readable
func (gtsd *GlobalTradeStationData) ExtractFrom(readable types.Readable) error {
	if err := gtsd.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData header. %s", err.Error())
	}

	if err := gtsd.DataID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData.DataID. %s", err.Error())
	}

	if err := gtsd.OwnerID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData.OwnerID. %s", err.Error())
	}

	if err := gtsd.UpdatedTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData.UpdatedTime. %s", err.Error())
	}

	if err := gtsd.IndexData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData.IndexData. %s", err.Error())
	}

	if err := gtsd.Version.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GlobalTradeStationData.Version. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GlobalTradeStationData
func (gtsd GlobalTradeStationData) Copy() types.RVType {
	copied := NewGlobalTradeStationData()

	copied.StructureVersion = gtsd.StructureVersion
	copied.DataID = gtsd.DataID.Copy().(types.UInt64)
	copied.OwnerID = gtsd.OwnerID.Copy().(types.PID)
	copied.UpdatedTime = gtsd.UpdatedTime.Copy().(types.DateTime)
	copied.IndexData = gtsd.IndexData.Copy().(types.QBuffer)
	copied.Version = gtsd.Version.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given GlobalTradeStationData contains the same data as the current GlobalTradeStationData
func (gtsd GlobalTradeStationData) Equals(o types.RVType) bool {
	if _, ok := o.(GlobalTradeStationData); !ok {
		return false
	}

	other := o.(GlobalTradeStationData)

	if gtsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !gtsd.DataID.Equals(other.DataID) {
		return false
	}

	if !gtsd.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !gtsd.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !gtsd.IndexData.Equals(other.IndexData) {
		return false
	}

	return gtsd.Version.Equals(other.Version)
}

// CopyRef copies the current value of the GlobalTradeStationData
// and returns a pointer to the new copy
func (gtsd GlobalTradeStationData) CopyRef() types.RVTypePtr {
	copied := gtsd.Copy().(GlobalTradeStationData)
	return &copied
}

// Deref takes a pointer to the GlobalTradeStationData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gtsd *GlobalTradeStationData) Deref() types.RVType {
	return *gtsd
}

// String returns the string representation of the GlobalTradeStationData
func (gtsd GlobalTradeStationData) String() string {
	return gtsd.FormatToString(0)
}

// FormatToString pretty-prints the GlobalTradeStationData using the provided indentation level
func (gtsd GlobalTradeStationData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GlobalTradeStationData{\n")
	fmt.Fprintf(&b, "%sDataID: %s,\n", indentationValues, gtsd.DataID)
	fmt.Fprintf(&b, "%sOwnerID: %s,\n", indentationValues, gtsd.OwnerID.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sUpdatedTime: %s,\n", indentationValues, gtsd.UpdatedTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sIndexData: %s,\n", indentationValues, gtsd.IndexData)
	fmt.Fprintf(&b, "%sVersion: %s,\n", indentationValues, gtsd.Version)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGlobalTradeStationData returns a new GlobalTradeStationData
func NewGlobalTradeStationData() GlobalTradeStationData {
	return GlobalTradeStationData{
		DataID:      types.NewUInt64(0),
		OwnerID:     types.NewPID(0),
		UpdatedTime: types.NewDateTime(0),
		IndexData:   types.NewQBuffer(nil),
		Version:     types.NewUInt32(0),
	}

}
