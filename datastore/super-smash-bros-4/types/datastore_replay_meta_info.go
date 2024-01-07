// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreReplayMetaInfo is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreReplayMetaInfo struct {
	types.Structure
	ReplayID   *types.PrimitiveU64
	Size       *types.PrimitiveU32
	Mode       *types.PrimitiveU8
	Style      *types.PrimitiveU8
	Rule       *types.PrimitiveU8
	Stage      *types.PrimitiveU8
	ReplayType *types.PrimitiveU8
	Players    *types.List[*DataStoreReplayPlayer]
	Winners    *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the DataStoreReplayMetaInfo from the given readable
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreReplayMetaInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreReplayMetaInfo header. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayID. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Size. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Mode. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Style. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Rule.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Rule. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Stage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Stage. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.ReplayType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.ReplayType. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Players.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Players. %s", err.Error())
	}

	err = dataStoreReplayMetaInfo.Winners.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReplayMetaInfo.Winners. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreReplayMetaInfo to the given writable
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreReplayMetaInfo.ReplayID.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Size.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Mode.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Style.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Rule.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Stage.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.ReplayType.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Players.WriteTo(contentWritable)
	dataStoreReplayMetaInfo.Winners.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreReplayMetaInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreReplayMetaInfo
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Copy() types.RVType {
	copied := NewDataStoreReplayMetaInfo()

	copied.StructureVersion = dataStoreReplayMetaInfo.StructureVersion

	copied.ReplayID = dataStoreReplayMetaInfo.ReplayID.Copy().(*types.PrimitiveU64)
	copied.Size = dataStoreReplayMetaInfo.Size.Copy().(*types.PrimitiveU32)
	copied.Mode = dataStoreReplayMetaInfo.Mode.Copy().(*types.PrimitiveU8)
	copied.Style = dataStoreReplayMetaInfo.Style.Copy().(*types.PrimitiveU8)
	copied.Rule = dataStoreReplayMetaInfo.Rule.Copy().(*types.PrimitiveU8)
	copied.Stage = dataStoreReplayMetaInfo.Stage.Copy().(*types.PrimitiveU8)
	copied.ReplayType = dataStoreReplayMetaInfo.ReplayType.Copy().(*types.PrimitiveU8)
	copied.Players = dataStoreReplayMetaInfo.Players.Copy().(*types.List[*DataStoreReplayPlayer])
	copied.Winners = dataStoreReplayMetaInfo.Winners.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReplayMetaInfo); !ok {
		return false
	}

	other := o.(*DataStoreReplayMetaInfo)

	if dataStoreReplayMetaInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreReplayMetaInfo.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dataStoreReplayMetaInfo.Size.Equals(other.Size) {
		return false
	}

	if !dataStoreReplayMetaInfo.Mode.Equals(other.Mode) {
		return false
	}

	if !dataStoreReplayMetaInfo.Style.Equals(other.Style) {
		return false
	}

	if !dataStoreReplayMetaInfo.Rule.Equals(other.Rule) {
		return false
	}

	if !dataStoreReplayMetaInfo.Stage.Equals(other.Stage) {
		return false
	}

	if !dataStoreReplayMetaInfo.ReplayType.Equals(other.ReplayType) {
		return false
	}

	if !dataStoreReplayMetaInfo.Players.Equals(other.Players) {
		return false
	}

	if !dataStoreReplayMetaInfo.Winners.Equals(other.Winners) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) String() string {
	return dataStoreReplayMetaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReplayMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreReplayMetaInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dataStoreReplayMetaInfo.ReplayID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStoreReplayMetaInfo.Size))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dataStoreReplayMetaInfo.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dataStoreReplayMetaInfo.Style))
	b.WriteString(fmt.Sprintf("%sRule: %s,\n", indentationValues, dataStoreReplayMetaInfo.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %s,\n", indentationValues, dataStoreReplayMetaInfo.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %s,\n", indentationValues, dataStoreReplayMetaInfo.ReplayType))
	b.WriteString(fmt.Sprintf("%sPlayers: %s,\n", indentationValues, dataStoreReplayMetaInfo.Players))
	b.WriteString(fmt.Sprintf("%sWinners: %s\n", indentationValues, dataStoreReplayMetaInfo.Winners))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReplayMetaInfo returns a new DataStoreReplayMetaInfo
func NewDataStoreReplayMetaInfo() *DataStoreReplayMetaInfo {
	dataStoreReplayMetaInfo := &DataStoreReplayMetaInfo{
		ReplayID: types.NewPrimitiveU64(0),
		Size: types.NewPrimitiveU32(0),
		Mode: types.NewPrimitiveU8(0),
		Style: types.NewPrimitiveU8(0),
		Rule: types.NewPrimitiveU8(0),
		Stage: types.NewPrimitiveU8(0),
		ReplayType: types.NewPrimitiveU8(0),
		Players: types.NewList[*DataStoreReplayPlayer](),
		Winners: types.NewList[*types.PrimitiveU32](),
	}

	dataStoreReplayMetaInfo.Players.Type = NewDataStoreReplayPlayer()
	dataStoreReplayMetaInfo.Winners.Type = types.NewPrimitiveU32(0)

	return dataStoreReplayMetaInfo
}
