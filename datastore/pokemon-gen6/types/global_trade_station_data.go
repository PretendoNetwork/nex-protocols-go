// Package types implements all the types used by the DataStore (Pokemon Gen6) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// GlobalTradeStationData holds data for the DataStore (Pokemon Gen6) protocol
type GlobalTradeStationData struct {
	types.Structure
	DataID      *types.PrimitiveU64
	OwnerID     *types.PID
	UpdatedTime *types.DateTime
	IndexData   *types.QBuffer
	Version     *types.PrimitiveU32
}

// ExtractFrom extracts the GlobalTradeStationData from the given readable
func (globalTradeStationData *GlobalTradeStationData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = globalTradeStationData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GlobalTradeStationData header. %s", err.Error())
	}

	err = globalTradeStationData.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.DataID from stream. %s", err.Error())
	}

	err = globalTradeStationData.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.OwnerID from stream. %s", err.Error())
	}

	err = globalTradeStationData.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.UpdatedTime from stream. %s", err.Error())
	}

	err = globalTradeStationData.IndexData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.IndexData from stream. %s", err.Error())
	}

	err = globalTradeStationData.Version.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GlobalTradeStationData.Version from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the GlobalTradeStationData to the given writable
func (globalTradeStationData *GlobalTradeStationData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	globalTradeStationData.DataID.WriteTo(contentWritable)
	globalTradeStationData.OwnerID.WriteTo(contentWritable)
	globalTradeStationData.UpdatedTime.WriteTo(contentWritable)
	globalTradeStationData.IndexData.WriteTo(contentWritable)
	globalTradeStationData.Version.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	globalTradeStationData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of GlobalTradeStationData
func (globalTradeStationData *GlobalTradeStationData) Copy() types.RVType {
	copied := NewGlobalTradeStationData()

	copied.StructureVersion = globalTradeStationData.StructureVersion

	copied.DataID = globalTradeStationData.DataID.Copy().(*types.PrimitiveU64)
	copied.OwnerID = globalTradeStationData.OwnerID.Copy().(*types.PID)
	copied.UpdatedTime = globalTradeStationData.UpdatedTime.Copy().(*types.DateTime)
	copied.IndexData = globalTradeStationData.IndexData.Copy().(*types.QBuffer)
	copied.Version = globalTradeStationData.Version.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (globalTradeStationData *GlobalTradeStationData) Equals(o types.RVType) bool {
	if _, ok := o.(*GlobalTradeStationData); !ok {
		return false
	}

	other := o.(*GlobalTradeStationData)

	if globalTradeStationData.StructureVersion != other.StructureVersion {
		return false
	}

	if !globalTradeStationData.DataID.Equals(other.DataID) {
		return false
	}

	if !globalTradeStationData.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !globalTradeStationData.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if !globalTradeStationData.IndexData.Equals(other.IndexData) {
		return false
	}

	if !globalTradeStationData.Version.Equals(other.Version) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, globalTradeStationData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, globalTradeStationData.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, globalTradeStationData.OwnerID))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, globalTradeStationData.UpdatedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIndexData: %s,\n", indentationValues, globalTradeStationData.IndexData))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, globalTradeStationData.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGlobalTradeStationData returns a new GlobalTradeStationData
func NewGlobalTradeStationData() *GlobalTradeStationData {
	return &GlobalTradeStationData{
		DataID: types.NewPrimitiveU64(0),
		OwnerID: types.NewPID(0),
		UpdatedTime: types.NewDateTime(0),
		IndexData: types.NewQBuffer(nil),
		Version: types.NewPrimitiveU32(0),
	}
}
