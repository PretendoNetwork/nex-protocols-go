// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePreparePostReplayParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePreparePostReplayParam struct {
	types.Structure
	Size          types.UInt32
	Mode          types.UInt8
	Style         types.UInt8
	Rule          types.UInt8
	Stage         types.UInt8
	ReplayType    types.UInt8
	CompetitionID types.UInt64
	Score         types.Int32
	Players       types.List[DataStoreReplayPlayer]
	Winners       types.List[types.UInt32]
	KeyVersion    types.UInt16
	ExtraData     types.List[types.String]
}

// WriteTo writes the DataStorePreparePostReplayParam to the given writable
func (dspprp DataStorePreparePostReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspprp.Size.WriteTo(contentWritable)
	dspprp.Mode.WriteTo(contentWritable)
	dspprp.Style.WriteTo(contentWritable)
	dspprp.Rule.WriteTo(contentWritable)
	dspprp.Stage.WriteTo(contentWritable)
	dspprp.ReplayType.WriteTo(contentWritable)
	dspprp.CompetitionID.WriteTo(contentWritable)
	dspprp.Score.WriteTo(contentWritable)
	dspprp.Players.WriteTo(contentWritable)
	dspprp.Winners.WriteTo(contentWritable)
	dspprp.KeyVersion.WriteTo(contentWritable)
	dspprp.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspprp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostReplayParam from the given readable
func (dspprp *DataStorePreparePostReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspprp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam header. %s", err.Error())
	}

	err = dspprp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Size. %s", err.Error())
	}

	err = dspprp.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Mode. %s", err.Error())
	}

	err = dspprp.Style.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Style. %s", err.Error())
	}

	err = dspprp.Rule.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Rule. %s", err.Error())
	}

	err = dspprp.Stage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Stage. %s", err.Error())
	}

	err = dspprp.ReplayType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ReplayType. %s", err.Error())
	}

	err = dspprp.CompetitionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.CompetitionID. %s", err.Error())
	}

	err = dspprp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Score. %s", err.Error())
	}

	err = dspprp.Players.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Players. %s", err.Error())
	}

	err = dspprp.Winners.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.Winners. %s", err.Error())
	}

	err = dspprp.KeyVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.KeyVersion. %s", err.Error())
	}

	err = dspprp.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostReplayParam
func (dspprp DataStorePreparePostReplayParam) Copy() types.RVType {
	copied := NewDataStorePreparePostReplayParam()

	copied.StructureVersion = dspprp.StructureVersion
	copied.Size = dspprp.Size.Copy().(types.UInt32)
	copied.Mode = dspprp.Mode.Copy().(types.UInt8)
	copied.Style = dspprp.Style.Copy().(types.UInt8)
	copied.Rule = dspprp.Rule.Copy().(types.UInt8)
	copied.Stage = dspprp.Stage.Copy().(types.UInt8)
	copied.ReplayType = dspprp.ReplayType.Copy().(types.UInt8)
	copied.CompetitionID = dspprp.CompetitionID.Copy().(types.UInt64)
	copied.Score = dspprp.Score.Copy().(types.Int32)
	copied.Players = dspprp.Players.Copy().(types.List[DataStoreReplayPlayer])
	copied.Winners = dspprp.Winners.Copy().(types.List[types.UInt32])
	copied.KeyVersion = dspprp.KeyVersion.Copy().(types.UInt16)
	copied.ExtraData = dspprp.ExtraData.Copy().(types.List[types.String])

	return copied
}

// Equals checks if the given DataStorePreparePostReplayParam contains the same data as the current DataStorePreparePostReplayParam
func (dspprp DataStorePreparePostReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostReplayParam); !ok {
		return false
	}

	other := o.(*DataStorePreparePostReplayParam)

	if dspprp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspprp.Size.Equals(other.Size) {
		return false
	}

	if !dspprp.Mode.Equals(other.Mode) {
		return false
	}

	if !dspprp.Style.Equals(other.Style) {
		return false
	}

	if !dspprp.Rule.Equals(other.Rule) {
		return false
	}

	if !dspprp.Stage.Equals(other.Stage) {
		return false
	}

	if !dspprp.ReplayType.Equals(other.ReplayType) {
		return false
	}

	if !dspprp.CompetitionID.Equals(other.CompetitionID) {
		return false
	}

	if !dspprp.Score.Equals(other.Score) {
		return false
	}

	if !dspprp.Players.Equals(other.Players) {
		return false
	}

	if !dspprp.Winners.Equals(other.Winners) {
		return false
	}

	if !dspprp.KeyVersion.Equals(other.KeyVersion) {
		return false
	}

	return dspprp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePreparePostReplayParam
func (dspprp DataStorePreparePostReplayParam) String() string {
	return dspprp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePreparePostReplayParam using the provided indentation level
func (dspprp DataStorePreparePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dspprp.Size))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dspprp.Mode))
	b.WriteString(fmt.Sprintf("%sStyle: %s,\n", indentationValues, dspprp.Style))
	b.WriteString(fmt.Sprintf("%sRule: %s,\n", indentationValues, dspprp.Rule))
	b.WriteString(fmt.Sprintf("%sStage: %s,\n", indentationValues, dspprp.Stage))
	b.WriteString(fmt.Sprintf("%sReplayType: %s,\n", indentationValues, dspprp.ReplayType))
	b.WriteString(fmt.Sprintf("%sCompetitionID: %s,\n", indentationValues, dspprp.CompetitionID))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dspprp.Score))
	b.WriteString(fmt.Sprintf("%sPlayers: %s,\n", indentationValues, dspprp.Players))
	b.WriteString(fmt.Sprintf("%sWinners: %s,\n", indentationValues, dspprp.Winners))
	b.WriteString(fmt.Sprintf("%sKeyVersion: %s,\n", indentationValues, dspprp.KeyVersion))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dspprp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostReplayParam returns a new DataStorePreparePostReplayParam
func NewDataStorePreparePostReplayParam() DataStorePreparePostReplayParam {
	return DataStorePreparePostReplayParam{
		Size:          types.NewUInt32(0),
		Mode:          types.NewUInt8(0),
		Style:         types.NewUInt8(0),
		Rule:          types.NewUInt8(0),
		Stage:         types.NewUInt8(0),
		ReplayType:    types.NewUInt8(0),
		CompetitionID: types.NewUInt64(0),
		Score:         types.NewInt32(0),
		Players:       types.NewList[DataStoreReplayPlayer](),
		Winners:       types.NewList[types.UInt32](),
		KeyVersion:    types.NewUInt16(0),
		ExtraData:     types.NewList[types.String](),
	}

}
