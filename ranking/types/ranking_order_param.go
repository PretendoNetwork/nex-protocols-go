// Package ranking_types implements all the types used by the Ranking protocol
package ranking_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingOrderParam holds parameters for ordering rankings
type RankingOrderParam struct {
	nex.Structure
	OrderCalculation uint8
	GroupIndex       uint8
	GroupNum         uint8
	TimeScope        uint8
	Offset           uint32
	Length           uint8
}

// ExtractFromStream extracts a RankingOrderParam structure from a stream
func (rankingOrderParam *RankingOrderParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingOrderParam.OrderCalculation, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.OrderCalculation from stream. %s", err.Error())
	}

	rankingOrderParam.GroupIndex, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupIndex from stream. %s", err.Error())
	}

	rankingOrderParam.GroupNum, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.GroupNum from stream. %s", err.Error())
	}

	rankingOrderParam.TimeScope, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.TimeScope from stream. %s", err.Error())
	}

	rankingOrderParam.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Offset from stream. %s", err.Error())
	}

	rankingOrderParam.Length, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingOrderParam.Length from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingOrderParam and returns a byte array
func (rankingOrderParam *RankingOrderParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(rankingOrderParam.OrderCalculation)
	stream.WriteUInt8(rankingOrderParam.GroupIndex)
	stream.WriteUInt8(rankingOrderParam.GroupNum)
	stream.WriteUInt8(rankingOrderParam.TimeScope)
	stream.WriteUInt32LE(rankingOrderParam.Offset)
	stream.WriteUInt8(rankingOrderParam.Length)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingOrderParam
func (rankingOrderParam *RankingOrderParam) Copy() nex.StructureInterface {
	copied := NewRankingOrderParam()

	copied.OrderCalculation = rankingOrderParam.OrderCalculation
	copied.GroupIndex = rankingOrderParam.GroupIndex
	copied.GroupNum = rankingOrderParam.GroupNum
	copied.TimeScope = rankingOrderParam.TimeScope
	copied.Offset = rankingOrderParam.Offset
	copied.Length = rankingOrderParam.Length

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingOrderParam *RankingOrderParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingOrderParam)

	if rankingOrderParam.OrderCalculation != other.OrderCalculation {
		return false
	}

	if rankingOrderParam.GroupIndex != other.GroupIndex {
		return false
	}

	if rankingOrderParam.GroupNum != other.GroupNum {
		return false
	}

	if rankingOrderParam.TimeScope != other.TimeScope {
		return false
	}

	if rankingOrderParam.Offset != other.Offset {
		return false
	}

	return rankingOrderParam.Length != other.Length
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingOrderParam.StructureVersion()))
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
