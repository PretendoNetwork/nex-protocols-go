// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePreparePostReplayParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePreparePostReplayParam struct {
	types.Structure
	Size          *types.PrimitiveU32
	Mode          *types.PrimitiveU8
	Style         *types.PrimitiveU8
	Rule          *types.PrimitiveU8
	Stage         *types.PrimitiveU8
	ReplayType    *types.PrimitiveU8
	CompetitionID *types.PrimitiveU64
	Score         *types.PrimitiveS32
	Players       *types.List[*DataStoreReplayPlayer]
	Winners       *types.List[*types.PrimitiveU32]
	KeyVersion    *types.PrimitiveU16
	ExtraData     *types.List[*types.String]
}

// ExtractFrom extracts the DataStorePreparePostReplayParam from the given readable
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePreparePostReplayParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePreparePostReplayParam header. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Size. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Mode. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Style. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Rule.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Rule. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Stage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Stage. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.ReplayType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ReplayType. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.CompetitionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.CompetitionID. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Score. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Players.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Players. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.Winners.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Winners. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.KeyVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.KeyVersion. %s", err.Error())
	}

	err = dataStorePreparePostReplayParam.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePreparePostReplayParam to the given writable
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePreparePostReplayParam.Size.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Mode.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Style.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Rule.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Stage.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.ReplayType.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.CompetitionID.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Score.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Players.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.Winners.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.KeyVersion.WriteTo(contentWritable)
	dataStorePreparePostReplayParam.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePreparePostReplayParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePreparePostReplayParam
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Copy() types.RVType {
	copied := NewDataStorePreparePostReplayParam()

	copied.StructureVersion = dataStorePreparePostReplayParam.StructureVersion

	copied.Size = dataStorePreparePostReplayParam.Size.Copy().(*types.PrimitiveU32)
	copied.Mode = dataStorePreparePostReplayParam.Mode.Copy().(*types.PrimitiveU8)
	copied.Style = dataStorePreparePostReplayParam.Style.Copy().(*types.PrimitiveU8)
	copied.Rule = dataStorePreparePostReplayParam.Rule.Copy().(*types.PrimitiveU8)
	copied.Stage = dataStorePreparePostReplayParam.Stage.Copy().(*types.PrimitiveU8)
	copied.ReplayType = dataStorePreparePostReplayParam.ReplayType.Copy().(*types.PrimitiveU8)
	copied.CompetitionID = dataStorePreparePostReplayParam.CompetitionID.Copy().(*types.PrimitiveU64)
	copied.Score = dataStorePreparePostReplayParam.Score.Copy().(*types.PrimitiveS32)
	copied.Players = dataStorePreparePostReplayParam.Players.Copy().(*types.List[*DataStoreReplayPlayer])
	copied.Winners = dataStorePreparePostReplayParam.Winners.Copy().(*types.List[*types.PrimitiveU32])
	copied.KeyVersion = dataStorePreparePostReplayParam.KeyVersion.Copy().(*types.PrimitiveU16)
	copied.ExtraData = dataStorePreparePostReplayParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostReplayParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostReplayParam)

	if dataStorePreparePostReplayParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePreparePostReplayParam.Size.Equals(other.Size) {
		return false
	}

	if !dataStorePreparePostReplayParam.Mode.Equals(other.Mode) {
		return false
	}

	if !dataStorePreparePostReplayParam.Style.Equals(other.Style) {
		return false
	}

	if !dataStorePreparePostReplayParam.Rule.Equals(other.Rule) {
		return false
	}

	if !dataStorePreparePostReplayParam.Stage.Equals(other.Stage) {
		return false
	}

	if !dataStorePreparePostReplayParam.ReplayType.Equals(other.ReplayType) {
		return false
	}

	if !dataStorePreparePostReplayParam.CompetitionID.Equals(other.CompetitionID) {
		return false
	}

	if !dataStorePreparePostReplayParam.Score.Equals(other.Score) {
		return false
	}

	if !dataStorePreparePostReplayParam.Players.Equals(other.Players) {
		return false
	}

	if !dataStorePreparePostReplayParam.Winners.Equals(other.Winners) {
		return false
	}

	if !dataStorePreparePostReplayParam.KeyVersion.Equals(other.KeyVersion) {
		return false
	}

	if !dataStorePreparePostReplayParam.ExtraData.Equals(other.ExtraData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) String() string {
	return dataStorePreparePostReplayParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePreparePostReplayParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStorePreparePostReplayParam.Size))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dataStorePreparePostReplayParam.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dataStorePreparePostReplayParam.Style))
	b.WriteString(fmt.Sprintf("%sRule: %s,\n", indentationValues, dataStorePreparePostReplayParam.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %s,\n", indentationValues, dataStorePreparePostReplayParam.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %s,\n", indentationValues, dataStorePreparePostReplayParam.ReplayType))
	b.WriteString(fmt.Sprintf("%sCompetitionID: %s,\n", indentationValues, dataStorePreparePostReplayParam.CompetitionID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStorePreparePostReplayParam.Score))
	b.WriteString(fmt.Sprintf("%sPlayers: %s,\n", indentationValues, dataStorePreparePostReplayParam.Players))
	b.WriteString(fmt.Sprintf("%sWinners: %s,\n", indentationValues, dataStorePreparePostReplayParam.Winners))
	b.WriteString(fmt.Sprintf("%sKeyVersion: %s,\n", indentationValues, dataStorePreparePostReplayParam.KeyVersion))
	b.WriteString(fmt.Sprintf("%sExtraData: %s\n", indentationValues, dataStorePreparePostReplayParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostReplayParam returns a new DataStorePreparePostReplayParam
func NewDataStorePreparePostReplayParam() *DataStorePreparePostReplayParam {
	dataStorePreparePostReplayParam := &DataStorePreparePostReplayParam{
		Size: types.NewPrimitiveU32(0),
		Mode: types.NewPrimitiveU8(0),
		Style: types.NewPrimitiveU8(0),
		Rule: types.NewPrimitiveU8(0),
		Stage: types.NewPrimitiveU8(0),
		ReplayType: types.NewPrimitiveU8(0),
		CompetitionID: types.NewPrimitiveU64(0),
		Score: types.NewPrimitiveS32(0),
		Players: types.NewList[*DataStoreReplayPlayer](),
		Winners: types.NewList[*types.PrimitiveU32](),
		KeyVersion: types.NewPrimitiveU16(0),
		ExtraData: types.NewList[*types.String](),
	}

	dataStorePreparePostReplayParam.Players.Type = NewDataStoreReplayPlayer()
	dataStorePreparePostReplayParam.Winners.Type = types.NewPrimitiveU32(0)
	dataStorePreparePostReplayParam.ExtraData.Type = types.NewString("")

	return dataStorePreparePostReplayParam
}
