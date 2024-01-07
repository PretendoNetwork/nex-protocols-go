// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePostFightingPowerScoreParam is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStorePostFightingPowerScoreParam struct {
	types.Structure
	Mode             *types.PrimitiveU8
	Score            *types.PrimitiveU32
	IsWorldHighScore *types.PrimitiveBool
}

// ExtractFrom extracts the DataStorePostFightingPowerScoreParam from the given readable
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePostFightingPowerScoreParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePostFightingPowerScoreParam header. %s", err.Error())
	}

	err = dataStorePostFightingPowerScoreParam.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Mode. %s", err.Error())
	}

	err = dataStorePostFightingPowerScoreParam.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.Score. %s", err.Error())
	}

	err = dataStorePostFightingPowerScoreParam.IsWorldHighScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePostFightingPowerScoreParam.IsWorldHighScore. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePostFightingPowerScoreParam to the given writable
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePostFightingPowerScoreParam.Mode.WriteTo(contentWritable)
	dataStorePostFightingPowerScoreParam.Score.WriteTo(contentWritable)
	dataStorePostFightingPowerScoreParam.IsWorldHighScore.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePostFightingPowerScoreParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePostFightingPowerScoreParam
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Copy() types.RVType {
	copied := NewDataStorePostFightingPowerScoreParam()

	copied.StructureVersion = dataStorePostFightingPowerScoreParam.StructureVersion

	copied.Mode = dataStorePostFightingPowerScoreParam.Mode.Copy().(*types.PrimitiveU8)
	copied.Score = dataStorePostFightingPowerScoreParam.Score.Copy().(*types.PrimitiveU32)
	copied.IsWorldHighScore = dataStorePostFightingPowerScoreParam.IsWorldHighScore.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePostFightingPowerScoreParam); !ok {
		return false
	}

	other := o.(*DataStorePostFightingPowerScoreParam)

	if dataStorePostFightingPowerScoreParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePostFightingPowerScoreParam.Mode.Equals(other.Mode) {
		return false
	}

	if !dataStorePostFightingPowerScoreParam.Score.Equals(other.Score) {
		return false
	}

	if !dataStorePostFightingPowerScoreParam.IsWorldHighScore.Equals(other.IsWorldHighScore) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) String() string {
	return dataStorePostFightingPowerScoreParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePostFightingPowerScoreParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePostFightingPowerScoreParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMode: %s,\n", indentationValues, dataStorePostFightingPowerScoreParam.Mode))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, dataStorePostFightingPowerScoreParam.Score))
	b.WriteString(fmt.Sprintf("%sIsWorldHighScore: %s\n", indentationValues, dataStorePostFightingPowerScoreParam.IsWorldHighScore))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePostFightingPowerScoreParam returns a new DataStorePostFightingPowerScoreParam
func NewDataStorePostFightingPowerScoreParam() *DataStorePostFightingPowerScoreParam {
	return &DataStorePostFightingPowerScoreParam{
		Mode: types.NewPrimitiveU8(0),
		Score: types.NewPrimitiveU32(0),
		IsWorldHighScore: types.NewPrimitiveBool(false),
	}
}
