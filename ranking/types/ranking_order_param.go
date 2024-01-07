// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingOrderParam holds parameters for ordering rankings
type RankingOrderParam struct {
	types.Structure
	OrderCalculation *types.PrimitiveU8
	GroupIndex       *types.PrimitiveU8
	GroupNum         *types.PrimitiveU8
	TimeScope        *types.PrimitiveU8
	Offset           *types.PrimitiveU32
	Length           *types.PrimitiveU8
}

// ExtractFrom extracts the RankingOrderParam from the given readable
func (rankingOrderParam *RankingOrderParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingOrderParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingOrderParam header. %s", err.Error())
	}

	err = rankingOrderParam.OrderCalculation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.OrderCalculation from stream. %s", err.Error())
	}

	err = rankingOrderParam.GroupIndex.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupIndex from stream. %s", err.Error())
	}

	err = rankingOrderParam.GroupNum.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupNum from stream. %s", err.Error())
	}

	err = rankingOrderParam.TimeScope.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.TimeScope from stream. %s", err.Error())
	}

	err = rankingOrderParam.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Offset from stream. %s", err.Error())
	}

	err = rankingOrderParam.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Length from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingOrderParam to the given writable
func (rankingOrderParam *RankingOrderParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingOrderParam.OrderCalculation.WriteTo(contentWritable)
	rankingOrderParam.GroupIndex.WriteTo(contentWritable)
	rankingOrderParam.GroupNum.WriteTo(contentWritable)
	rankingOrderParam.TimeScope.WriteTo(contentWritable)
	rankingOrderParam.Offset.WriteTo(contentWritable)
	rankingOrderParam.Length.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingOrderParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingOrderParam
func (rankingOrderParam *RankingOrderParam) Copy() types.RVType {
	copied := NewRankingOrderParam()

	copied.StructureVersion = rankingOrderParam.StructureVersion

	copied.OrderCalculation = rankingOrderParam.OrderCalculation
	copied.GroupIndex = rankingOrderParam.GroupIndex
	copied.GroupNum = rankingOrderParam.GroupNum
	copied.TimeScope = rankingOrderParam.TimeScope
	copied.Offset = rankingOrderParam.Offset
	copied.Length = rankingOrderParam.Length

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingOrderParam *RankingOrderParam) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingOrderParam); !ok {
		return false
	}

	other := o.(*RankingOrderParam)

	if rankingOrderParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !rankingOrderParam.OrderCalculation.Equals(other.OrderCalculation) {
		return false
	}

	if !rankingOrderParam.GroupIndex.Equals(other.GroupIndex) {
		return false
	}

	if !rankingOrderParam.GroupNum.Equals(other.GroupNum) {
		return false
	}

	if !rankingOrderParam.TimeScope.Equals(other.TimeScope) {
		return false
	}

	if !rankingOrderParam.Offset.Equals(other.Offset) {
		return false
	}

	return rankingOrderParam.Length == other.Length
}

// String returns a string representation of the struct
func (rankingOrderParam *RankingOrderParam) String() string {
	return rankingOrderParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingOrderParam *RankingOrderParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingOrderParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingOrderParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOrderCalculation: %d,\n", indentationValues, rankingOrderParam.OrderCalculation))
	b.WriteString(fmt.Sprintf("%sGroupIndex: %d,\n", indentationValues, rankingOrderParam.GroupIndex))
	b.WriteString(fmt.Sprintf("%sGroupNum: %d,\n", indentationValues, rankingOrderParam.GroupNum))
	b.WriteString(fmt.Sprintf("%sTimeScope: %d,\n", indentationValues, rankingOrderParam.TimeScope))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, rankingOrderParam.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %d\n", indentationValues, rankingOrderParam.Length))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingOrderParam returns a new RankingOrderParam
func NewRankingOrderParam() *RankingOrderParam {
	return &RankingOrderParam{}
}
