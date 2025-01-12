// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePostFightingPowerScoreParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePostFightingPowerScoreParam struct {
	types.Structure
	Mode             types.UInt8
	Score            types.UInt32
	IsWorldHighScore types.Bool
}

// WriteTo writes the DataStorePostFightingPowerScoreParam to the given writable
func (dspfpsp DataStorePostFightingPowerScoreParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspfpsp.Mode.WriteTo(contentWritable)
	dspfpsp.Score.WriteTo(contentWritable)
	dspfpsp.IsWorldHighScore.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspfpsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePostFightingPowerScoreParam from the given readable
func (dspfpsp *DataStorePostFightingPowerScoreParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspfpsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam header. %s", err.Error())
	}

	err = dspfpsp.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Mode. %s", err.Error())
	}

	err = dspfpsp.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Score. %s", err.Error())
	}

	err = dspfpsp.IsWorldHighScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.IsWorldHighScore. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePostFightingPowerScoreParam
func (dspfpsp DataStorePostFightingPowerScoreParam) Copy() types.RVType {
	copied := NewDataStorePostFightingPowerScoreParam()

	copied.StructureVersion = dspfpsp.StructureVersion
	copied.Mode = dspfpsp.Mode.Copy().(types.UInt8)
	copied.Score = dspfpsp.Score.Copy().(types.UInt32)
	copied.IsWorldHighScore = dspfpsp.IsWorldHighScore.Copy().(types.Bool)

	return copied
}

// Equals checks if the given DataStorePostFightingPowerScoreParam contains the same data as the current DataStorePostFightingPowerScoreParam
func (dspfpsp DataStorePostFightingPowerScoreParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStorePostFightingPowerScoreParam); !ok {
		return false
	}

	other := o.(DataStorePostFightingPowerScoreParam)

	if dspfpsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspfpsp.Mode.Equals(other.Mode) {
		return false
	}

	if !dspfpsp.Score.Equals(other.Score) {
		return false
	}

	return dspfpsp.IsWorldHighScore.Equals(other.IsWorldHighScore)
}

// CopyRef copies the current value of the DataStorePostFightingPowerScoreParam
// and returns a pointer to the new copy
func (dspfpsp DataStorePostFightingPowerScoreParam) CopyRef() types.RVTypePtr {
	copied := dspfpsp.Copy().(DataStorePostFightingPowerScoreParam)
	return &copied
}

// Deref takes a pointer to the DataStorePostFightingPowerScoreParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dspfpsp *DataStorePostFightingPowerScoreParam) Deref() types.RVType {
	return *dspfpsp
}

// String returns the string representation of the DataStorePostFightingPowerScoreParam
func (dspfpsp DataStorePostFightingPowerScoreParam) String() string {
	return dspfpsp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePostFightingPowerScoreParam using the provided indentation level
func (dspfpsp DataStorePostFightingPowerScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostFightingPowerScoreParam{\n")
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dspfpsp.Mode))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dspfpsp.Score))
	b.WriteString(fmt.Sprintf("%sIsWorldHighScore: %s,\n", indentationValues, dspfpsp.IsWorldHighScore))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostFightingPowerScoreParam returns a new DataStorePostFightingPowerScoreParam
func NewDataStorePostFightingPowerScoreParam() DataStorePostFightingPowerScoreParam {
	return DataStorePostFightingPowerScoreParam{
		Mode:             types.NewUInt8(0),
		Score:            types.NewUInt32(0),
		IsWorldHighScore: types.NewBool(false),
	}

}
