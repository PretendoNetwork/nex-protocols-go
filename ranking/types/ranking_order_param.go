// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/ranking/constants"
)

// RankingOrderParam is a type within the Ranking protocol
type RankingOrderParam struct {
	types.Structure
	OrderCalculation constants.OrderCalculation
	GroupIndex       constants.FilterGroupIndex
	GroupNum         types.UInt8
	TimeScope        constants.TimeScope
	Offset           types.UInt32
	Length           types.UInt8
}

// WriteTo writes the RankingOrderParam to the given writable
func (rop RankingOrderParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rop.OrderCalculation.WriteTo(contentWritable)
	rop.GroupIndex.WriteTo(contentWritable)
	rop.GroupNum.WriteTo(contentWritable)
	rop.TimeScope.WriteTo(contentWritable)
	rop.Offset.WriteTo(contentWritable)
	rop.Length.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rop.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingOrderParam from the given readable
func (rop *RankingOrderParam) ExtractFrom(readable types.Readable) error {
	if err := rop.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam header. %s", err.Error())
	}

	if err := rop.OrderCalculation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.OrderCalculation. %s", err.Error())
	}

	if err := rop.GroupIndex.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupIndex. %s", err.Error())
	}

	if err := rop.GroupNum.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupNum. %s", err.Error())
	}

	if err := rop.TimeScope.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.TimeScope. %s", err.Error())
	}

	if err := rop.Offset.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Offset. %s", err.Error())
	}

	if err := rop.Length.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Length. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingOrderParam
func (rop RankingOrderParam) Copy() types.RVType {
	copied := NewRankingOrderParam()

	copied.StructureVersion = rop.StructureVersion
	copied.OrderCalculation = rop.OrderCalculation
	copied.GroupIndex = rop.GroupIndex
	copied.GroupNum = rop.GroupNum.Copy().(types.UInt8)
	copied.TimeScope = rop.TimeScope
	copied.Offset = rop.Offset.Copy().(types.UInt32)
	copied.Length = rop.Length.Copy().(types.UInt8)

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

	if rop.OrderCalculation != other.OrderCalculation {
		return false
	}

	if rop.GroupIndex != other.GroupIndex {
		return false
	}

	if !rop.GroupNum.Equals(other.GroupNum) {
		return false
	}

	if rop.TimeScope != other.TimeScope {
		return false
	}

	if !rop.Offset.Equals(other.Offset) {
		return false
	}

	return rop.Length.Equals(other.Length)
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
	b.WriteString(fmt.Sprintf("%sOrderCalculation: %s,\n", indentationValues, rop.OrderCalculation))
	b.WriteString(fmt.Sprintf("%sGroupIndex: %s,\n", indentationValues, rop.GroupIndex))
	b.WriteString(fmt.Sprintf("%sGroupNum: %s,\n", indentationValues, rop.GroupNum))
	b.WriteString(fmt.Sprintf("%sTimeScope: %s,\n", indentationValues, rop.TimeScope))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, rop.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %s,\n", indentationValues, rop.Length))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingOrderParam returns a new RankingOrderParam
func NewRankingOrderParam() RankingOrderParam {
	return RankingOrderParam{
		OrderCalculation: constants.OrderCalculation113,
		GroupIndex:       constants.FilterGroupIndex0,
		GroupNum:         types.NewUInt8(0),
		TimeScope:        constants.TimeScopeCustom0,
		Offset:           types.NewUInt32(0),
		Length:           types.NewUInt8(0),
	}

}
