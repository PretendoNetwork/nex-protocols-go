// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReplayMetaInfo is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreReplayMetaInfo struct {
	types.Structure
	ReplayID   types.UInt64
	Size       types.UInt32
	Mode       types.UInt8
	Style      types.UInt8
	Rule       types.UInt8
	Stage      types.UInt8
	ReplayType types.UInt8
	Players    types.List[DataStoreReplayPlayer]
	Winners    types.List[types.UInt32]
}

// WriteTo writes the DataStoreReplayMetaInfo to the given writable
func (dsrmi DataStoreReplayMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrmi.ReplayID.WriteTo(contentWritable)
	dsrmi.Size.WriteTo(contentWritable)
	dsrmi.Mode.WriteTo(contentWritable)
	dsrmi.Style.WriteTo(contentWritable)
	dsrmi.Rule.WriteTo(contentWritable)
	dsrmi.Stage.WriteTo(contentWritable)
	dsrmi.ReplayType.WriteTo(contentWritable)
	dsrmi.Players.WriteTo(contentWritable)
	dsrmi.Winners.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrmi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReplayMetaInfo from the given readable
func (dsrmi *DataStoreReplayMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrmi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo header. %s", err.Error())
	}

	err = dsrmi.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayID. %s", err.Error())
	}

	err = dsrmi.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Size. %s", err.Error())
	}

	err = dsrmi.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Mode. %s", err.Error())
	}

	err = dsrmi.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Style. %s", err.Error())
	}

	err = dsrmi.Rule.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Rule. %s", err.Error())
	}

	err = dsrmi.Stage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Stage. %s", err.Error())
	}

	err = dsrmi.ReplayType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayType. %s", err.Error())
	}

	err = dsrmi.Players.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Players. %s", err.Error())
	}

	err = dsrmi.Winners.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Winners. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReplayMetaInfo
func (dsrmi DataStoreReplayMetaInfo) Copy() types.RVType {
	copied := NewDataStoreReplayMetaInfo()

	copied.StructureVersion = dsrmi.StructureVersion
	copied.ReplayID = dsrmi.ReplayID.Copy().(types.UInt64)
	copied.Size = dsrmi.Size.Copy().(types.UInt32)
	copied.Mode = dsrmi.Mode.Copy().(types.UInt8)
	copied.Style = dsrmi.Style.Copy().(types.UInt8)
	copied.Rule = dsrmi.Rule.Copy().(types.UInt8)
	copied.Stage = dsrmi.Stage.Copy().(types.UInt8)
	copied.ReplayType = dsrmi.ReplayType.Copy().(types.UInt8)
	copied.Players = dsrmi.Players.Copy().(types.List[DataStoreReplayPlayer])
	copied.Winners = dsrmi.Winners.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given DataStoreReplayMetaInfo contains the same data as the current DataStoreReplayMetaInfo
func (dsrmi DataStoreReplayMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreReplayMetaInfo); !ok {
		return false
	}

	other := o.(DataStoreReplayMetaInfo)

	if dsrmi.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrmi.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dsrmi.Size.Equals(other.Size) {
		return false
	}

	if !dsrmi.Mode.Equals(other.Mode) {
		return false
	}

	if !dsrmi.Style.Equals(other.Style) {
		return false
	}

	if !dsrmi.Rule.Equals(other.Rule) {
		return false
	}

	if !dsrmi.Stage.Equals(other.Stage) {
		return false
	}

	if !dsrmi.ReplayType.Equals(other.ReplayType) {
		return false
	}

	if !dsrmi.Players.Equals(other.Players) {
		return false
	}

	return dsrmi.Winners.Equals(other.Winners)
}

// CopyRef copies the current value of the DataStoreReplayMetaInfo
// and returns a pointer to the new copy
func (dsrmi DataStoreReplayMetaInfo) CopyRef() types.RVTypePtr {
	copied := dsrmi.Copy().(DataStoreReplayMetaInfo)
	return &copied
}

// Deref takes a pointer to the DataStoreReplayMetaInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsrmi *DataStoreReplayMetaInfo) Deref() types.RVType {
	return *dsrmi
}

// String returns the string representation of the DataStoreReplayMetaInfo
func (dsrmi DataStoreReplayMetaInfo) String() string {
	return dsrmi.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReplayMetaInfo using the provided indentation level
func (dsrmi DataStoreReplayMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReplayMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dsrmi.ReplayID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dsrmi.Size))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dsrmi.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dsrmi.Style))
	b.WriteString(fmt.Sprintf("%sRule: %s,\n", indentationValues, dsrmi.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %s,\n", indentationValues, dsrmi.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %s,\n", indentationValues, dsrmi.ReplayType))
	b.WriteString(fmt.Sprintf("%sPlayers: %s,\n", indentationValues, dsrmi.Players))
	b.WriteString(fmt.Sprintf("%sWinners: %s,\n", indentationValues, dsrmi.Winners))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReplayMetaInfo returns a new DataStoreReplayMetaInfo
func NewDataStoreReplayMetaInfo() DataStoreReplayMetaInfo {
	return DataStoreReplayMetaInfo{
		ReplayID:   types.NewUInt64(0),
		Size:       types.NewUInt32(0),
		Mode:       types.NewUInt8(0),
		Style:      types.NewUInt8(0),
		Rule:       types.NewUInt8(0),
		Stage:      types.NewUInt8(0),
		ReplayType: types.NewUInt8(0),
		Players:    types.NewList[DataStoreReplayPlayer](),
		Winners:    types.NewList[types.UInt32](),
	}

}
