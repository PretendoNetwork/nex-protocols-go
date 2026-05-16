// Package types implements all the types used by the legacy Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/ranking/legacy/constants"
)

// RankingOrderParam is a type within the Ranking protocol
type RankingOrderParam struct {
	types.Structure
	ScoreIndex      constants.ScoreIndex
	ScoreOrder      constants.OrderBy
	RankCalculation constants.OrderCalculation
	Unknown1        types.UInt8
	Unknown2        types.UInt8
	Unknown3        types.UInt8
	Unknown4        types.UInt32
}

// WriteTo writes the RankingOrderParam to the given writable
func (rop RankingOrderParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rop.ScoreIndex.WriteTo(contentWritable)
	rop.ScoreOrder.WriteTo(contentWritable)
	rop.RankCalculation.WriteTo(contentWritable)
	rop.Unknown1.WriteTo(contentWritable)
	rop.Unknown2.WriteTo(contentWritable)
	rop.Unknown3.WriteTo(contentWritable)
	rop.Unknown4.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rop.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingOrderParam from the given readable
func (rop *RankingOrderParam) ExtractFrom(readable types.Readable) error {
	if err := rop.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam header. %s", err.Error())
	}

	if err := rop.ScoreIndex.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.ScoreIndex. %s", err.Error())
	}

	if err := rop.ScoreOrder.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.ScoreOrder. %s", err.Error())
	}

	if err := rop.RankCalculation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.RankCalculation. %s", err.Error())
	}

	if err := rop.Unknown1.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Unknown1. %s", err.Error())
	}

	if err := rop.Unknown2.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Unknown2. %s", err.Error())
	}

	if err := rop.Unknown3.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Unknown3. %s", err.Error())
	}

	if err := rop.Unknown4.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Unknown4. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingOrderParam
func (rop RankingOrderParam) Copy() types.RVType {
	copied := NewRankingOrderParam()

	copied.StructureVersion = rop.StructureVersion
	copied.ScoreIndex = rop.ScoreIndex
	copied.ScoreOrder = rop.ScoreOrder
	copied.RankCalculation = rop.RankCalculation
	copied.Unknown1 = rop.Unknown1.Copy().(types.UInt8)
	copied.Unknown2 = rop.Unknown2.Copy().(types.UInt8)
	copied.Unknown3 = rop.Unknown3.Copy().(types.UInt8)
	copied.Unknown4 = rop.Unknown4.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given RankingOrderParam contains the same data as the current RankingOrderParam
func (rop RankingOrderParam) Equals(o types.RVType) bool {
	if _, ok := o.(RankingOrderParam); !ok {
		return false
	}

	other := o.(RankingOrderParam)

	if rop.StructureVersion != other.StructureVersion {
		return false
	}

	if rop.ScoreIndex != other.ScoreIndex {
		return false
	}

	if rop.ScoreOrder != other.ScoreOrder {
		return false
	}

	if rop.RankCalculation != other.RankCalculation {
		return false
	}

	if !rop.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !rop.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !rop.Unknown3.Equals(other.Unknown3) {
		return false
	}

	return rop.Unknown4.Equals(other.Unknown4)
}

// CopyRef copies the current value of the RankingOrderParam
// and returns a pointer to the new copy
func (rop RankingOrderParam) CopyRef() types.RVTypePtr {
	copied := rop.Copy().(RankingOrderParam)
	return &copied
}

// Deref takes a pointer to the RankingOrderParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rop *RankingOrderParam) Deref() types.RVType {
	return *rop
}

// String returns the string representation of the RankingOrderParam
func (rop RankingOrderParam) String() string {
	return rop.FormatToString(0)
}

// FormatToString pretty-prints the RankingOrderParam using the provided indentation level
func (rop RankingOrderParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingOrderParam{\n")
	b.WriteString(fmt.Sprintf("%sScoreIndex: %s,\n", indentationValues, rop.ScoreIndex))
	b.WriteString(fmt.Sprintf("%sScoreOrder: %s,\n", indentationValues, rop.ScoreOrder))
	b.WriteString(fmt.Sprintf("%sRankCalculation: %s,\n", indentationValues, rop.RankCalculation))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rop.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rop.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, rop.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, rop.Unknown4))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingOrderParam returns a new RankingOrderParam
func NewRankingOrderParam() RankingOrderParam {
	return RankingOrderParam{
		ScoreIndex:      constants.ScoreIndex0,
		ScoreOrder:      constants.OrderByAscending,
		RankCalculation: constants.OrderCalculation113,
		Unknown1:        types.NewUInt8(0),
		Unknown2:        types.NewUInt8(0),
		Unknown3:        types.NewUInt8(0),
		Unknown4:        types.NewUInt32(0),
	}

}
